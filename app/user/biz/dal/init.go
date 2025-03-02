package dal

import (
	"github.com/MelodyDeep/TikTok-E-commerce/app/user/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
