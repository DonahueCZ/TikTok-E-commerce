package service

import (
	"context"
	"testing"

	product "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
)

var (
	searchMap = map[string]string{
		"portable":      "Smartphones",
		"devices":       "Smartphones",
		"communication": "Smartphones",
		"various":       "Smartphones",
		"applications":  "Smartphones",
		"powerful":      "Laptops",
		"mobile":        "Laptops",
		"computers":     "Laptops",
		"work":          "Laptops",
		"entertainment": "Laptops",
		"comfortable":   "Sofas",
		"seating":       "Sofas",
		"furniture":     "Sofas",
		"living":        "Sofas",
		"rooms":         "Sofas",
		"dining":        "Dining Tables",
		"gatherings":    "Dining Tables",
		"casual":        "T-Shirts",
		"tops":          "T-Shirts",
		"typically":     "T-Shirts",
		"made":          "T-Shirts",
		"cotton":        "T-Shirts",
		"formal":        "Suits",
		"outfits":       "Suits",
		"usually":       "Suits",
		"worn":          "Suits",
		"business":      "Suits",
		"special":       "Suits",
		"occasions":     "Suits",
		"spherical":     "Basketballs",
		"balls":         "Basketballs",
		"sport":         "Basketballs",
		"basketball":    "Basketballs",
		"equipment":     "Tennis Rackets",
		"hit":           "Tennis Rackets",
		"ball":          "Tennis Rackets",
		"tennis":        "Tennis Rackets",
		"solid":         "Chocolate Bars",
		"chocolate":     "Chocolate Bars",
		"treats":        "Chocolate Bars",
		"often":         "Chocolate Bars",
		"different":     "Chocolate Bars",
		"flavors":       "Chocolate Bars",
		"seeds":         "Coffee Beans",
		"coffee":        "Coffee Beans",
		"plant":         "Coffee Beans",
		"make":          "Coffee Beans",
	}
	searchPrdResp *product.SearchProductsResp
)

func TestSearchProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSearchProductsService(ctx)
	// init req and assert value

	// req := &product.SearchProductsReq{}
	// resp, err := s.Run(req)
	// t.Logf("err: %v", err)
	// t.Logf("resp: %v", resp)

	// todo: edit your unit test
	var prd *product.Product
	for keyword, wantedPrd := range searchMap {
		searchPrdResp, err = s.Run(&product.SearchProductsReq{Query: keyword})
		if err != nil {
			t.Errorf("SearchProduct() err: %v", err)
			continue
		}
		prd = searchPrdResp.Results[0]
		if prd.Id != uint32(productMap[wantedPrd].ID) || prd.Name != productMap[wantedPrd].Name {
			t.Errorf("GetProduct() err: %v", err)
		}
	}
}
