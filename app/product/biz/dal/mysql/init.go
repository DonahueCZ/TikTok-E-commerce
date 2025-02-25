package mysql

import (
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/model"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	// migrate Product and Category table
	err = DB.AutoMigrate(&model.Product{}, &model.Category{})
	if err != nil {
		panic(err)
	}
}
