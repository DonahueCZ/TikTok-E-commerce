package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/model"
	product "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
)

var (
	oldCategoryCases = []model.Category{
		{Name: "category1"},
		{Name: "category2"},
	}
	newCategoryCases2 = []model.Category{
		{Name: "category3"},
		{Name: "category4"},
	}
	updateCases = model.Product{
		Name:       "UpdateTest",
		Categories: oldCategoryCases,
	}
)

func TestUpdateProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateProductService(ctx)
	// init req and assert value

	// req := &product.UpdateProductReq{}
	// resp, err := s.Run(req)
	// t.Logf("err: %v", err)
	// t.Logf("resp: %v", resp)

	// todo: edit your unit test

	// create update case
	dao := model.NewDao(s.ctx, mysql.DB)
	dao.CreateProduct(&updateCases)

	category := make([]string, len(updateCases.Categories))
	for i, v := range newCategoryCases2 {
		category[i] = v.Name
	}
	// fmt.Println(updateCases.ID)
	resp, err := s.Run(&product.UpdateProductReq{
		Product: &product.Product{
			Id:         updateCases.ID,
			Name:       updateCases.Name,
			Categories: category,
		},
	})
	if err != nil || !resp.Success {
		t.Errorf("UpdateProduct_Run(): %v", err)
	}

	prd, err := dao.GetProductById(updateCases.ID)
	if err != nil {
		t.Errorf("failed get product")
	}
	cMap := map[string]*model.Category{}
	for _, category := range newCategoryCases2 {
		c, err := dao.GetCategoryByName(category.Name)
		if err != nil {
			t.Errorf("failed to get category")
		}
		cMap[category.Name] = c
	}
	for _, category := range prd.Categories {
		if !reflect.DeepEqual(category, *cMap[category.Name]) {
			// fmt.Println(category)
			// fmt.Println(cMap[category.Name])
			t.Errorf("failed to update category")
		}
	}
	// check if old category is deleted
	for _, category := range oldCategoryCases {
		_, err := dao.GetCategoryByName(category.Name)
		if err == nil {
			t.Errorf("failed to delete unused category")
		}
	}
	// check if evil update operation can be detected
	_, err = s.Run(&product.UpdateProductReq{
		Product: &product.Product{
			Id:         updateCases.ID,
			StoreId:    1,
			Name:       updateCases.Name,
			Categories: category,
		},
	})
	if err == nil {
		t.Errorf("failed to stop evil update operation")
	}
}
