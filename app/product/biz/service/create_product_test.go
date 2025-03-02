package service

import (
	"context"
	"testing"

	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/model"
	product "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
)

var (
	electronics        = model.Category{Name: "EThings"}
	homeAppliances     = model.Category{Name: "Home Appliances"}
	clothing           = model.Category{Name: "Clothing"}
	createProductCases = []model.Product{
		{
			Name:        "Smart Watch",
			Description: "A high - end smart watch with multiple health monitoring functions.",
			Picture:     "smart_watch.jpg",
			Price:       39900, // 表示 399.00
			Stock:       50,
			Categories:  []model.Category{electronics},
		},
		{
			Name:        "Refrigerator",
			Description: "A large - capacity refrigerator with energy - saving features.",
			Picture:     "refrigerator.jpg",
			Price:       599900, // 表示 5999.00
			Stock:       20,
			Categories:  []model.Category{homeAppliances},
		},
		{
			Name:        "Cotton T - Shirt",
			Description: "A comfortable cotton t - shirt in a classic design.",
			Picture:     "cotton_tshirt.jpg",
			Price:       3900, // 表示 39.00
			Stock:       100,
			Categories:  []model.Category{clothing},
		},
		{
			Name:        "Wireless Earbuds",
			Description: "Wireless earbuds with excellent sound quality and long battery life.",
			Picture:     "wireless_earbuds.jpg",
			Price:       19900, // 表示 199.00
			Stock:       80,
			Categories:  []model.Category{electronics},
		},
		{
			Name:        "Vacuum Cleaner",
			Description: "A powerful vacuum cleaner for efficient home cleaning.",
			Picture:     "vacuum_cleaner.jpg",
			Price:       299900, // 表示 2999.00
			Stock:       15,
			Categories:  []model.Category{homeAppliances},
		},
	}
	newCategoryCases = []model.Category{electronics, homeAppliances, clothing}
)

func TestCreateProduct_Run(t *testing.T) {
	// fmt.Printf("%+v", createProductCases)
	ctx := context.Background()
	s := NewCreateProductService(ctx)

	// check if product create successfully
	for _, c := range createProductCases {
		category := make([]string, len(c.Categories))
		for i, v := range c.Categories {
			category[i] = v.Name
		}
		resp, err := s.Run(&product.CreateProductReq{
			Product: &product.Product{
				StoreId:     1,
				Name:        c.Name,
				Description: c.Description,
				Picture:     c.Picture,
				Price:       c.Price,
				Stock:       c.Stock,
				Categories:  category,
			},
		})
		if err != nil {
			t.Errorf("CreateProduct(%v) err: %v", category, err)
			continue
		}
		prd, err := model.NewDao(ctx, mysql.DB).GetProductById(resp.Id)
		if err != nil {
			t.Errorf("CreateProduct() err: %v", err)
			continue
		}
		if prd.Name != c.Name || prd.Description != c.Description ||
			prd.Picture != c.Picture || prd.Price != c.Price ||
			prd.Stock != c.Stock {
			t.Errorf("expect %v, but got %v", c, prd)
		}
	}
	// check if related category create
	for _, c := range newCategoryCases {
		category, err := model.NewDao(ctx, mysql.DB).GetCategoryByName(c.Name)
		if err != nil {
			t.Errorf("CreateProduct(%v) err: %v", c.Name, err)
			continue
		}
		if category.Name != c.Name {
			t.Errorf(ExpectButGetErrTemplate, c.Name, category.Name)
		}
	}
}
