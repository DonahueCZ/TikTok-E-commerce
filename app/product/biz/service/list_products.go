package service

import (
	"context"

	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/redis"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/model"
	product "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	var (
		products []model.Product
		category *model.Category
		cq       = model.NewCachedQuery(s.ctx, mysql.DB.Session(&gorm.Session{}), redis.RedisClient)
	)
	if req.CategoryName != "" {
		category, err = cq.GetCategoryByName(req.CategoryName)
		if err != nil {
			return nil, err
		}
	}
	products, err = cq.GetProductPageByCategory(req.Page, req.PageSize, category)
	if err != nil {
		return nil, err
	}

	// fmt.Println(req.CategoryName, req.Page, req.PageSize, products)
	resp = &product.ListProductsResp{}
	resp.Products = make([]*product.Product, 0, len(products))
	for _, prd := range products {
		resp.Products = append(resp.Products,
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
