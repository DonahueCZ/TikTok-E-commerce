package main

import (
	paymentservice "TikTok-E-commerce-payment/kitex_gen/payment_proto_idl/idl/paymentservice/paymentservice"
	"log"
)

func main() {
	svr := paymentservice.NewServer(new(PaymentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
