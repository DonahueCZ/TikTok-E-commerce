package dal

import (
	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
