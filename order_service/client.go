package main

import (
	"context"
	"fmt"
	"order/rpc/order/config"
	"order/rpc/order/kitex_gen/demo/order_service"
	"order/rpc/order/kitex_gen/demo/order_service/orderservice"
	"time"

	"github.com/cloudwego/kitex/client"
	etcdRegistry "github.com/kitex-contrib/registry-etcd"
)

func main() {
	conf := config.GetConfig()
	r, err := etcdRegistry.NewEtcdResolver([]string{conf.EtcdHost})
	if err != nil {
		panic(err)
	}
	//etcdClient, err := etcd.NewClient(etcd.Options{})
	//if err != nil {
	//	panic(err)
	//}
	cli, err := orderservice.NewClient(
		conf.ServiceName,
		client.WithResolver(r),
		//client.WithSuite(etcdclient.NewSuite(conf.ServiceName, conf.ServiceName+".client", etcdClient)),
	)
	if err != nil {
		panic(err)
	}
	req := &order_service.CreateOrderRequest{
		UserId:     time.Now().Unix(),
		GoodsId:    1,
		GoodsCount: 1,
		Cost:       100,
		AddresseeInfo: &order_service.AddresseeInfo{
			Name:    "test",
			Phone:   "1234567890",
			Address: "test",
		},
	}
	resp, err := cli.CreateOrder(context.Background(), req)

	// req := &order_service.QueryOrderByIdRequest{
	// 	OrderId: 1,
	// }
	// resp, err := cli.QueryOrderById(context.Background(), req)
	// req := &order_service.QueryOrdersByUserIdRequest{
	// 	UserId:   1738497618,
	// 	PageNum:  1,
	// 	PageSize: 10,
	// }
	//resp, err := cli.QueryOrdersByUserId(context.Background(), req)
	//req := &order_service.UpdateOrderRequest{
	//	OrderId:    1,
	//	Status:     ordermd.IsPaid,
	//	GoodsCount: 1,
	//	Cost:       1,
	//	AddresseeInfo: &order_service.AddresseeInfo{
	//		Name:    "test",
	//		Phone:   "1234567890",
	//		Address: "hello",
	//	},
	//}
	//resp, err := cli.UpdateOrder(context.Background(), req)
	// req := &order_service.UpdateOrderStatusRequest{
	// 	OrderId: 1,
	// 	Status:  ordermd.IsPaid,
	// }
	// resp, err := cli.UpdateOrderStatus(context.Background(), req)
	// req := &order_service.UpdateOrderAddresseeInfoRequest{
	// 	OrderId: 1,
	// 	AddresseeInfo: &order_service.AddresseeInfo{
	// 		Name:    "test",
	// 		Phone:   "1234567890",
	// 		Address: "hello",
	// 	},
	// }
	// resp, err := cli.UpdateOrderAddresseeInfo(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	for {
		time.Sleep(time.Second * 10)
	}
}
