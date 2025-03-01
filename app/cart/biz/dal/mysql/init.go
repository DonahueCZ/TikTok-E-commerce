///初始化与 MySQL 数据库的连接，使用了 GORM 作为 ORM（对象关系映射）库

package mysql

import (
	"fmt"

	"github.com/MelodyDeep/TikTok-E-commerce/app/cart/biz/model"
	"github.com/MelodyDeep/TikTok-E-commerce/app/cart/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	// dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, "gorm", "gorm", "localhost")

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	DB.AutoMigrate(&model.Cart{})
	if err != nil {
		panic(err)
	}
}
