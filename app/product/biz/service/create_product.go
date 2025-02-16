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

type CreateProductService struct {
	ctx context.Context
} // NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create note info
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	// Finish your business logic.
	var (
		dao         = model.NewCacheDao(s.ctx, mysql.DB.Session(&gorm.Session{}), redis.RedisClient)
		categories  = make([]model.Category, len(req.Product.Categories))
		tmpCategory *model.Category
		prd         = &model.Product{
			StoreId:     req.Product.StoreId,
			Name:        req.Product.Name,
			Description: req.Product.Description,
			Picture:     req.Product.Picture,
			Price:       req.Product.Price,
			Stock:       req.Product.Stock,
		}
	)
	// get or init categories
	for i, category := range req.Product.Categories {
		tmpCategory, err = dao.GetOrCreateCategoryByName(category)
		if err != nil {
			return
		}
		categories[i] = *tmpCategory
	}

	// bind categories and product
	prd.Categories = categories

	// create product
	id := dao.CreateProduct(prd)
	if id == 0 {
		return nil, errors.New("fail to create product")
	}
	return &product.CreateProductResp{Id: id}, nil
}
