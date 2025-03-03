package main

import (
	"fmt"
	"net"
	"time"

	"github.com/MelodyDeep/TikTok-E-commerce/app/cart/biz/dal"
	"github.com/MelodyDeep/TikTok-E-commerce/app/cart/conf"
	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/infra/rpc"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	etcd "github.com/kitex-contrib/registry-etcd"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	_ = godotenv.Load()
	dal.Init()
	rpc.InitClient()

	opts := kitexInit()

	svr := cartservice.NewServer(new(CartServiceImpl), opts...)

	err := svr.Run()
	fmt.Println("svr run ...")
	if err != nil {
		klog.Error(err.Error())
	}
	fmt.Println("server stoped")
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	//service registery
	r, err := etcd.NewEtcdRegistry(conf.GetConf().Registry.RegistryAddress)
	if err != nil {
		klog.Fatal(err)
	}

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}), server.WithRegistry(r))

	// klog
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
