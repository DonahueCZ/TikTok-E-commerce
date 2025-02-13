package service

import (
	"context"

	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/redis"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/model"
	product "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	var prd *model.Product
	prd, err = model.NewCachedQuery(s.ctx, mysql.DB, redis.RedisClient).GetProductById(req.Id)
	if err != nil {
		return nil, err
	}

	resp = &product.GetProductResp{
		Product: &product.Product{
			Id:          prd.ID,
			Name:        prd.Name,
			Description: prd.Description,
			Picture:     prd.Picture,
			Price:       prd.Price,
			Stock:       prd.Stock,
		},
	}
	return
}
