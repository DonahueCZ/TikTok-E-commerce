package dal

import (
	"github.com/MelodyDeep/TikTok-E-commerce/app/cart/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
