package service

import (
	"context"
	"testing"

	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/redis"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/model"
	product "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
)

var prdList = []model.Product{
	{Name: "apple"},
	{Name: "banana"},
	{Name: "juice"},
}

func TestDeleteProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteProductService(ctx)
	// todo: edit your unit test
	dao := model.NewCacheDao(s.ctx, mysql.DB, redis.RedisClient)
	for _, prd := range prdList {
		id := dao.CreateProduct(&prd)
		if id == 0 {
			t.Errorf("fail to create test cases")
			continue
		}
		resp, err := s.Run(&product.DeleteProductReq{
			Id: id,
		})
		if err != nil || !resp.Success {
			t.Errorf("DeleteProduct(): %v", err)
			continue
		}
		_, err = dao.GetProductById(id)
		if err == nil  {
			t.Errorf("fail to delete product: %v", err)
		}
	}
}
