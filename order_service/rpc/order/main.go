package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"order/rpc/order/config"
	"order/rpc/order/kitex_gen/demo/order_service/orderservice"
	"order/rpc/order/logger"
	"os"
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils/kitexutil"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	etcdRegistry "github.com/kitex-contrib/registry-etcd"
)

func main() {
	// 设置log
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelDebug)

	conf := config.GetConfig()

	startTime := time.Now().Format("2006-01-02_15-04-05")
	_, err := os.Open(conf.LogDir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(conf.LogDir, os.ModePerm)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	f, err := os.OpenFile(fmt.Sprintf("%s/%s.log", conf.LogDir, startTime), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	klog.SetOutput(logger.NewLoggerWriter(f))

	// 启动服务

	r, err := etcdRegistry.NewEtcdRegistry([]string{conf.EtcdHost})
	//etcdClient, _ := etcd.NewClient(etcd.Options{})
	addr, _ := net.ResolveTCPAddr("tcp", conf.ListenOn)
	if err != nil {
		panic(err)
	}
	svr := orderservice.NewServer(
		new(OrderServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.ServiceName}), server.WithRegistry(r),
		//server.WithSuite(etcdServer.NewSuite(conf.ServiceName, etcdClient)),
		server.WithMiddleware(func(next endpoint.Endpoint) endpoint.Endpoint {
			return func(ctx context.Context, req, resp any) error {
				method, ok := kitexutil.GetMethod(ctx)
				if !ok {
					return next(ctx, req, resp)
				}
				cluster, ok := kitexutil.GetCallerAddr(ctx)
				if !ok {
					return next(ctx, req, resp)
				}
				klog.Infof("rev: method: %s, cluster: %s", method, cluster)
				return next(ctx, req, resp)
			}
		}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
