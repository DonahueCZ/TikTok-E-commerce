package service

import (
	"context"
	"errors"

	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/redis"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/model"
	product "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	// Finish your business logic.
	dao := model.NewCacheDao(s.ctx, mysql.DB.Session(&gorm.Session{}), redis.RedisClient)
	// check if store id is matched
	prd, err := dao.GetProductById(req.Id)
	if err != nil {
		return nil, err
	}
	if prd.StoreId != req.StoreId {
		return nil, errors.New("matched store id")
	}
	err = dao.DeleteProduct(&model.Product{Base: model.Base{ID: req.Id}, StoreId: req.StoreId})
	if err != nil {
		return nil, err
	}
	resp = &product.DeleteProductResp{
		Success: true,
	}
	return
}
