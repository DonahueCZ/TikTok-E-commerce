package service

import (
	"context"

	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/redis"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/model"
	product "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	var products []*model.Product
	products, err = model.NewCachedQuery(s.ctx, mysql.DB, redis.RedisClient).GetProductsByQuery(req.Query)
	if err != nil {
		return nil, err
	}

	resp = &product.SearchProductsResp{}
	resp.Results = make([]*product.Product, 0, len(products))
	for _, prd := range products {
		resp.Results = append(resp.Results,
			&product.Product{
				Id:          prd.ID,
				Name:        prd.Name,
				Description: prd.Description,
				Picture:     prd.Picture,
				Price:       prd.Price,
				Stock:       prd.Stock,
			},
		)
	}
	return
}
