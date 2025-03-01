package service

import (
	"context"
	"testing"

	product "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
)

var getPrdResp *product.GetProductResp

func TestGetProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetProductService(ctx)
	// todo: edit your unit test
	for _, wantedPrd := range productMap {
		getPrdResp, err = s.Run(&product.GetProductReq{Id: uint32(wantedPrd.ID)})
		if err != nil {
			t.Errorf("GetProduct() err: %v", err)
			continue
		}
		if getPrdResp.Product.Id != uint32(wantedPrd.ID) || getPrdResp.Product.Name != wantedPrd.Name {
			t.Errorf("GetProduct() err: %v", err)
		}
		// fmt.Println(getPrdResp.Product.Categories)
	}
	_, err = s.Run(&product.GetProductReq{Id: uint32(0)})
	if err == nil {
		t.Error("GetProduct() can't accept an empty id")
	}
}
