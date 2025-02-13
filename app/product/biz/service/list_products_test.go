package service

import (
	"context"
	"testing"

	product "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
)

var (
	categoryProductsMap = map[string][]string{
		"Electronics":          {"Smartphones", "Laptops"},
		"Mobile Devices":       {"Smartphones"},
		"Computers":            {"Laptops"},
		"Home Furniture":       {"Sofas", "Dining Tables"},
		"Living Room Items":    {"Sofas"},
		"Dining Room Items":    {"Dining Tables"},
		"Apparel":              {"T-Shirts", "Suits"},
		"Casual Wear":          {"T-Shirts"},
		"Formal Wear":          {"Suits"},
		"Sports Equipment":     {"Basketballs", "Tennis Rackets"},
		"Ball Games":           {"Basketballs"},
		"Racket Sports":        {"Tennis Rackets"},
		"Food":                 {"Chocolate Bars", "Coffee Beans"},
		"Snacks":               {"Chocolate Bars"},
		"Beverage Ingredients": {"Coffee Beans"},
	}
	maxProductsLen = 2
)

func TestListProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListProductsService(ctx)
	// init req and assert value

	// req := &product.ListProductsReq{}
	// resp, err := s.Run(req)
	// t.Logf("err: %v", err)
	// t.Logf("resp: %v", resp)

	// // todo: edit your unit test
	for c, list := range categoryProductsMap {
		// test query by category name
		resp, err := s.Run(&product.ListProductsReq{
			Page:         1,
			PageSize:     uint32(maxProductsLen),
			CategoryName: c,
		})
		if err != nil {
			t.Errorf("ListProducts() err: %v", err)
			continue
		}
		testMap := map[string]struct{}{}
		for _, prd := range resp.Products {
			testMap[prd.Name] = struct{}{}
		}
		for _, prd := range list {
			if _, ok := testMap[prd]; !ok {
				t.Errorf("ListProducts() err: %v expected but not found", prd)
				continue
			}
		}

		// test pagination
		for i := 1; i <= maxProductsLen; i++ {
			resp, err = s.Run(&product.ListProductsReq{
				Page:         uint32(i),
				PageSize:     1,
				CategoryName: c,
			})
			if err != nil {
				t.Errorf("ListProducts() err: %v", err)
				continue
			}
			if i <= len(categoryProductsMap[c]) && len(resp.Products) != 1 {
				t.Error("expected get one, but not")
				continue
			}
			if i > len(categoryProductsMap[c]) && len(resp.Products) != 0 {
				t.Error("expected get none, but not")
			}
		}
	}
	// with no category test
	for i := range 10 {
		resp, err := s.Run(&product.ListProductsReq{
			Page:         1,
			PageSize:     uint32(i+1),
			CategoryName: "",
		})
		if err != nil {
			t.Errorf("ListProducts() err: %v", err)
			return
		}
		if len(resp.Products) != (i+1) {
			t.Errorf("expect %d, but get %d", i+1, len(resp.Products))
		}
	}
}
