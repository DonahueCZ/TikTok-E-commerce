package main

import (
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/payment/paymentservice"
	etcdRegistry "github.com/kitex-contrib/registry-etcd"
	"net"
	"time"

	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
)

func main() {
	opts := kitexInit()

	svr := paymentservice.NewServer(new(PaymentServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// 解析服务器地址
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// 设置服务基本信息
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	// 注册 etcd 服务
	r, err := etcdRegistry.NewEtcdRegistry(conf.GetConf().Registry.RegistryAddress)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithRegistry(r))

	// 配置日志
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

	// 注册关闭时的清理逻辑
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
