package service

import (
	"context"
	"errors"

	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/mysql"
	payment "github.com/MelodyDeep/TikTok-E-commerce_payment/rpc_gen/kitex_gen/rpc/payment"
)

type GetOrderByIDService struct {
	ctx context.Context
}

// NewGetOrderByIDService 创建 GetOrderByIDService
func NewGetOrderByIDService(ctx context.Context) *GetOrderByIDService {
	return &GetOrderByIDService{ctx: ctx}
}

// Run 查询订单详情
func (s *GetOrderByIDService) Run(req *payment.OrderRequest) (resp *payment.OrderResponse, err error) {
	// 1. 通过订单ID查询订单
	order, err := mysql.OrderRepo.GetOrderByID(s.ctx, req.OrderId)
	if err != nil {
		return nil, err
	}

	// 2. 如果订单不存在，返回错误
	if order == nil {
		return nil, errors.New("订单不存在")
	}

	// 3. 组装返回结果
	resp = &payment.OrderResponse{
		OrderId:       order.OrderID,
		UserId:        order.UserID,
		Amount:        order.Amount,
		PaymentMethod: order.PaymentMethod,
		Status:        order.Status,
		CreatedAt:     order.CreatedAt.Unix(), // 转换为时间戳格式
		UpdatedAt:     order.UpdatedAt.Unix(),
	}

	return resp, nil
}
