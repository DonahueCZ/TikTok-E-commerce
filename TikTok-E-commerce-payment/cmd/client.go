package main

import (
	"TikTok-E-commerce-payment/idl/payment_proto_idl/idl/paymentservice" // 导入你生成的服务代码
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"log"
)

func main() {
	// 创建 Kitex 客户端，连接到服务器（假设服务器地址为 localhost:9000）
	c, err := client.NewClient("PaymentService", client.WithHostPorts("localhost:9000"))
	if err != nil {
		log.Fatal(err)
	}

	// 创建请求数据
	req := &paymentservice.PaymentRequest{
		OrderId:       "123456",
		Amount:        100.50,
		PaymentMethod: "credit_card",
	}

	// 调用 ProcessPayment RPC 方法
	resp, err := c.ProcessPayment(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	// 打印响应
	fmt.Printf("Payment Status: %s\n", resp.Status)
	fmt.Printf("Message: %s\n", resp.Message)
}
