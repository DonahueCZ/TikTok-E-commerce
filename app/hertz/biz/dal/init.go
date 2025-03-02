package dal

import (
	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
