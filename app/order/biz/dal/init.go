package dal

import (
	"github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
