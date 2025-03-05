package main

import (
	"database/sql"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"time"

	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/service"
	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/conf"
	payment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment/paymentservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	_ "github.com/go-sql-driver/mysql" // ✅ 记得导入 MySQL 驱动
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	etcd "github.com/kitex-contrib/registry-etcd"
	"go.uber.org/zap/zapcore"
)

func main() {
	// ✅ 1. 读取数据库配置
	db, err := sql.Open("mysql", conf.GetConf().Database.DSN)
	if err != nil {
		klog.Fatal("数据库连接失败: ", err)
	}
	defer db.Close()

	// ✅ 2. 创建 OrderRepository
	orderRepo := mysql.NewOrderRepository(db)

	// ✅ 3. 启动 Kitex 服务器
	opts := kitexInit()
	svr := payment.NewServer(service.NewPaymentServiceImpl(*orderRepo), opts...) // ✅ 这里要传 `*orderRepo`，因为 `service.NewPaymentServiceImpl` 需要接口类型

	// ✅ 4. 运行服务器
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}

	asyncWriter := &lumberjack.Logger{
		Filename:   "logs/payment.log", // ✅ 修改为你自己的日志路径
		MaxSize:    100,                // 100MB
		MaxBackups: 5,
		MaxAge:     7, // 7天
	}
	klog.SetOutput(asyncWriter)
}

// ✅ Kitex 初始化函数
func kitexInit() (opts []server.Option) {
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	r, err := etcd.NewEtcdRegistry(conf.GetConf().Registry.RegistryAddress)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithRegistry(r))

	// ✅ 修正 lumberjack 作为日志
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())

	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)

	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
