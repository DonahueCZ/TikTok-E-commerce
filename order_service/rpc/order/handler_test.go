package main

import (
	"context"
	order_service "order/rpc/order/kitex_gen/demo/order_service"
	"reflect"
	"testing"
)

func TestOrderServiceImpl_CreateOrder(t *testing.T) {
	type args struct {
		ctx context.Context
		req *order_service.CreateOrderRequest
	}
	tests := []struct {
		name     string
		s        *OrderServiceImpl
		args     args
		wantResp *order_service.BaseResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.s.CreateOrder(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderServiceImpl.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("OrderServiceImpl.CreateOrder() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestOrderServiceImpl_QueryOrderById(t *testing.T) {
	type args struct {
		ctx context.Context
		req *order_service.QueryOrderByIdRequest
	}
	tests := []struct {
		name     string
		s        *OrderServiceImpl
		args     args
		wantResp *order_service.QueryOrderResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.s.QueryOrderById(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderServiceImpl.QueryOrderById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("OrderServiceImpl.QueryOrderById() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestOrderServiceImpl_QueryOrdersByUserId(t *testing.T) {
	type args struct {
		ctx context.Context
		req *order_service.QueryOrdersByUserIdRequest
	}
	tests := []struct {
		name     string
		s        *OrderServiceImpl
		args     args
		wantResp *order_service.QueryOrdersResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.s.QueryOrdersByUserId(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderServiceImpl.QueryOrdersByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("OrderServiceImpl.QueryOrdersByUserId() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestOrderServiceImpl_UpdateOrder(t *testing.T) {
	type args struct {
		ctx context.Context
		req *order_service.UpdateOrderRequest
	}
	tests := []struct {
		name     string
		s        *OrderServiceImpl
		args     args
		wantResp *order_service.BaseResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.s.UpdateOrder(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderServiceImpl.UpdateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("OrderServiceImpl.UpdateOrder() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestOrderServiceImpl_UpdateOrderStatus(t *testing.T) {
	type args struct {
		ctx context.Context
		req *order_service.UpdateOrderStatusRequest
	}
	tests := []struct {
		name     string
		s        *OrderServiceImpl
		args     args
		wantResp *order_service.BaseResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.s.UpdateOrderStatus(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderServiceImpl.UpdateOrderStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("OrderServiceImpl.UpdateOrderStatus() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestOrderServiceImpl_UpdateOrderAddresseeInfo(t *testing.T) {
	type args struct {
		ctx context.Context
		req *order_service.UpdateOrderAddresseeInfoRequest
	}
	tests := []struct {
		name     string
		s        *OrderServiceImpl
		args     args
		wantResp *order_service.BaseResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.s.UpdateOrderAddresseeInfo(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderServiceImpl.UpdateOrderAddresseeInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("OrderServiceImpl.UpdateOrderAddresseeInfo() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
