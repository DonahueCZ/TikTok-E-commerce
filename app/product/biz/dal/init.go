package dal

import (
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
