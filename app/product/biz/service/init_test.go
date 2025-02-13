package service

import (
	"context"

	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/redis"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/model"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	ctx          context.Context
	db           *gorm.DB
	envPath      = "test.env"
	productCases = [][2]string{
		{"Smartphones", "Portable devices for communication and various applications."},
		{"Laptops", "Powerful mobile computers for work, entertainment, and more."},
		{"Sofas", "Comfortable seating furniture for living rooms."},
		{"Dining Tables", "Furniture used for dining and gatherings."},
		{"T-Shirts", "Casual tops typically made of cotton."},
		{"Suits", "Formal outfits usually worn for business or special occasions."},
		{"Basketballs", "Spherical balls used in the sport of basketball."},
		{"Tennis Rackets", "Equipment used to hit the ball in tennis."},
		{"Chocolate Bars", "Solid chocolate treats, often with different flavors."},
		{"Coffee Beans", "Seeds from the coffee plant used to make coffee."},
	}
	categoryCases = []string{
		"Electronics",
		"Mobile Devices",
		"Computers",
		"Home Furniture",
		"Living Room Items",
		"Dining Room Items",
		"Apparel",
		"Casual Wear",
		"Formal Wear",
		"Sports Equipment",
		"Ball Games",
		"Racket Sports",
		"Food",
		"Snacks",
		"Beverage Ingredients",
	}
	productCategoryMap = map[string][]string{
		"Smartphones":    {"Electronics", "Mobile Devices"},
		"Laptops":        {"Electronics", "Computers"},
		"Sofas":          {"Home Furniture", "Living Room Items"},
		"Dining Tables":  {"Home Furniture", "Dining Room Items"},
		"T-Shirts":       {"Apparel", "Casual Wear"},
		"Suits":          {"Apparel", "Formal Wear"},
		"Basketballs":    {"Sports Equipment", "Ball Games"},
		"Tennis Rackets": {"Sports Equipment", "Racket Sports"},
		"Chocolate Bars": {"Food", "Snacks"},
		"Coffee Beans":   {"Food", "Beverage Ingredients"},
	}
	productMap  map[string](*model.Product)
	categoryMap map[string](*model.Category)
	err         error
)

func init() {
	ctx = context.Background()

	// load env variables (GO_ENV)
	if err = godotenv.Load(envPath); err != nil {
		panic(err)
	}
	dal.Init()
	db = mysql.DB

	// reset table
	db.Migrator().DropTable(model.ProductTable, model.CategoryTable, model.ProductCategoryTable)
	db.AutoMigrate(&model.Product{}, &model.Category{})

	// generate data cases & init db
	productMap, categoryMap = map[string](*model.Product){}, map[string](*model.Category){}
	for _, mCase := range categoryCases {
		category := &model.Category{
			Name: mCase,
		}
		db.Create(&category)
		categoryMap[mCase] = category
	}
	for i, mCase := range productCases {
		price := uint32(i * 10)
		product := &model.Product{
			Name:        mCase[0],
			Description: mCase[1],
			Price:       price,
		}
		for _, c := range productCategoryMap[mCase[0]] {
			product.Categories = append(product.Categories, *categoryMap[c])
		}
		db.Create(&product)
		productMap[mCase[0]] = product
	}


	// reset redis
	if err = redis.RedisClient.FlushDB(context.Background()).Err(); err != nil {
		panic(err)
	}
}
