package service

import (
	"context"

	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/redis"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/model"
	product "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

type UpdateProductService struct {
	ctx context.Context
} // NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	// Finish your business logic.
	var (
		prd            *model.Product
		prevCategories []model.Category
	)
	dao := model.NewCacheDao(s.ctx, mysql.DB.Session(&gorm.Session{}), redis.RedisClient)

	prd, err = dao.GetProductById(req.Product.Id)
	if err != nil {
		return nil, err
	}
	prevCategories = prd.Categories
	prd = &model.Product{
		Base:        prd.Base,
		Name:        req.Product.Name,
		Description: req.Product.Description,
		Picture:     req.Product.Picture,
		Price:       req.Product.Price,
		Stock:       req.Product.Stock,
		Categories:  make([]model.Category, len(req.Product.Categories)),
	}

	// fmt.Println("budong", req.Product.Categories)
	for i, categoryName := range req.Product.Categories {
		// get or create category
		category, err := dao.GetOrCreateCategoryByName(categoryName)
		if err != nil {
			return nil, err
		}
		prd.Categories[i] = *category
	}
	model.NewCacheDao(s.ctx, mysql.DB, redis.RedisClient).UpdateProduct(prd)

	for _, category := range prevCategories {
		dao.DelUnusedCategory(&category)
	}
	return &product.UpdateProductResp{Success: true}, nil
}
