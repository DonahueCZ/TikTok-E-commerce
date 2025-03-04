package dao

import (
	"context"
	"errors"

	"github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/dal"
	"github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/dal/ordermd"
	cacheredis "github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/dal/redis"

	"sync"
	"time"

	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

type OrderDAO struct {
	db    *gorm.DB
	cache *redis.Client
}

var (
	instanceDAO *OrderDAO
	onceDAO     sync.Once
)

func GetOrderDAO() *OrderDAO {
	return NewOrderDAO()
}

func NewOrderDAO() *OrderDAO {
	onceDAO.Do(func() {
		dal.Init()
		db := mysql.DB
		cache := cacheredis.RedisClient
		err := db.AutoMigrate(&ordermd.Order{})
		if err != nil {
			panic(err)
		}
		instanceDAO = &OrderDAO{
			db:    db,
			cache: cache,
		}

	})
	return instanceDAO
}

func (dao *OrderDAO) DB() *gorm.DB {
	return dao.db
}

func (dao *OrderDAO) Cache() *redis.Client {
	return dao.cache
}

func (dao *OrderDAO) Insert(order *ordermd.Order) error {
	return dao.db.Create(order).Error
}

func (dao *OrderDAO) FindOne(id int64) (*ordermd.Order, error) {
	q := &ordermd.Order{
		Id: id,
	}
	var redisResult ordermd.Order
	err := dao.cache.Get(context.Background(), q.GetCacheKey()).Scan(&redisResult)
	if err == nil {
		return &redisResult, nil
	}
	err = dao.db.First(q).Error
	if err != nil {
		return nil, err
	}
	// 检查订单是否为未支付状态且已过期
	if q.Status == ordermd.IsNotPaid {
		if time.Now().Unix()-q.CreateTime > ordermd.ExpireTime {
			q.Status = ordermd.IsExpire
			err := dao.Update(q)
			if err != nil {
				return nil, err
			}
		}
	}
	// 将从数据库中获取的数据存储到Redis中
	_ = dao.cache.HSet(context.Background(), q.GetCacheKey(), q, time.Hour).Err()

	return q, err
}

func (dao *OrderDAO) Update(order *ordermd.Order) error {
	err := dao.db.Save(order).Error
	if err != nil {
		return err
	}
	cacheKey := order.GetCacheKey()
	// 更新缓存
	_ = dao.cache.HSet(context.Background(), cacheKey, order, time.Hour).Err()
	return nil
}

func (dao *OrderDAO) Delete(id int64) error {
	order := &ordermd.Order{
		Id: id,
	}
	err := dao.db.Delete(order).Error
	if err != nil {
		return err
	}

	err = dao.cache.Del(context.Background(), order.GetCacheKey()).Err()
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}
	return nil
}

func (dao *OrderDAO) FindByUserId(userId int64, page, pageSize int32) ([]*ordermd.Order, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 10 {
		pageSize = 10
	}
	var orders []*ordermd.Order
	q := &ordermd.Order{
		UserId: userId,
	}
	err := dao.db.Where(q).Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	for _, order := range orders {
		if q.Status == ordermd.IsNotPaid {
			if time.Now().Unix()-order.CreateTime > ordermd.ExpireTime {
				err := dao.Update(order)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return orders, err
}
