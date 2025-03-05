package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/MelodyDeep/TikTok-E-commerce/app/user/biz/dal"
	"github.com/MelodyDeep/TikTok-E-commerce/app/user/biz/dal/mysql"
	cacheredis "github.com/MelodyDeep/TikTok-E-commerce/app/user/biz/dal/redis"
	"github.com/MelodyDeep/TikTok-E-commerce/app/user/biz/dal/usermd"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"sync"
	"time"
)

type UserDAO struct {
	db    *gorm.DB
	cache *redis.Client
}

var (
	instanceDAO *UserDAO
	onceDAO     sync.Once
)

func GetUserDAO() *UserDAO { return NewUserDao() }

func NewUserDao() *UserDAO {
	onceDAO.Do(func() {
		dal.Init()
		db := mysql.DB
		cache := cacheredis.RedisClient
		err := db.AutoMigrate(&usermd.User{})
		if err != nil {
			klog.Error("Failed to auto migrate:", err)
			return
		}
		instanceDAO = &UserDAO{
			db:    db,
			cache: cache,
		}
	})
	return instanceDAO
}

func (dao *UserDAO) DB() *gorm.DB { return dao.db }

func (dao *UserDAO) Cache() *redis.Client { return dao.cache }

// 插入数据
func (dao *UserDAO) Insert(user *usermd.User) error { return dao.db.Create(user).Error }

//func (dao *UserDAO) FindOne(userid int64) (*usermd.User, error) {
//	q := &usermd.User{
//		UserId: userid,
//	}
//	var redisResult usermd.User
//	err := dao.cache.Get(context.Background(), q.GetCacheKey()).Scan(&redisResult)
//	if err == nil {
//		return &redisResult, nil
//	}
//	if err != redis.Nil {
//		return nil, err
//	}
//	err = dao.db.First(q).Error
//	if err != nil {
//		return nil, err
//	}
//	err = dao.cache.Set(context.Background(), q.GetCacheKey(), q, time.Hour).Err()
//	if err != nil {
//		return nil, err
//	}
//	return q, nil
//}

func (dao *UserDAO) FindOne(ctx context.Context, userid int64) (*usermd.User, error) {
	q := &usermd.User{
		UserId: userid,
	}
	cacheKey := q.GetCacheKey()

	// 尝试从 Redis 缓存中获取用户信息
	val, err := dao.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var redisResult usermd.User
		err := json.Unmarshal([]byte(val), &redisResult)
		if err != nil {
			klog.Error("用户信息反序列化失败，UserId:", userid, " 错误信息:", err)
			return nil, fmt.Errorf("用户信息反序列化失败: %v", err)
		}
		klog.Info("从 Redis 缓存中成功获取用户信息，UserId:", userid)
		return &redisResult, nil
	}
	if err != redis.Nil {
		klog.Error("Redis 查询失败，UserId:", userid, " 错误信息:", err)
		return nil, err
	}

	// 如果缓存中没有，从数据库中查询
	var user usermd.User
	err = dao.db.Where("user_id = ?", userid).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			klog.Info("用户不存在，UserId:", userid)
			return nil, nil // 用户不存在，返回 nil, nil
		}
		klog.Error("数据库查询失败，UserId:", userid, " 错误信息:", err)
		return nil, err // 其他错误，返回具体错误信息
	}

	// 将查询到的用户信息存入 Redis 缓存
	userJSON, err := json.Marshal(user)
	if err != nil {
		klog.Error("用户信息序列化失败，UserId:", userid, " 错误信息:", err)
		return nil, fmt.Errorf("用户信息序列化失败: %v", err)
	}

	err = dao.cache.Set(ctx, cacheKey, userJSON, time.Hour).Err()
	if err != nil {
		klog.Error("设置 Redis 缓存失败，UserId:", userid, " 错误信息:", err)
		return nil, fmt.Errorf("设置 Redis 缓存失败: %v", err)
	}

	return &user, nil
}

func (dao *UserDAO) Update(user *usermd.User) error {
	err := dao.db.Save(user).Error
	if err != nil {
		return err
	}
	cacheKey := user.GetCacheKey()

	userJSON, err := json.Marshal(user)
	if err != nil {
		klog.Error("用户信息序列化失败，UserId:", user.UserId, " 错误信息:", err)
		return fmt.Errorf("用户信息序列化失败: %v", err)
	}

	err = dao.cache.Set(context.Background(), cacheKey, userJSON, time.Hour).Err()
	if err != nil {
		klog.Error("设置 Redis 缓存失败，UserId:", user.UserId, " 错误信息:", err)
		return fmt.Errorf("设置 Redis 缓存失败: %v", err)
	}

	return nil
}

func (dao *UserDAO) Delete(userId int64) error {
	user := &usermd.User{
		UserId: userId,
	}
	// 使用 GORM 的软删除功能
	err := dao.db.Delete(user).Error
	if err != nil {
		return err
	}

	err = dao.cache.Del(context.Background(), user.GetCacheKey()).Err()
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}
	return nil
}

func (dao *UserDAO) FindByEmail(email string) (*usermd.User, error) {
	// 尝试从 Redis 缓存中获取用户信息
	cacheKey := fmt.Sprintf("users:email:%s", email)
	val, err := dao.cache.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var cachedUser usermd.User
		err := json.Unmarshal([]byte(val), &cachedUser)
		if err != nil {
			klog.Error("用户信息反序列化失败，Email:", email, " 错误信息:", err)
			return nil, fmt.Errorf("用户信息反序列化失败: %v", err)
		}
		return &cachedUser, nil
	}
	if err != redis.Nil {
		klog.Error("Redis 查询失败:", err)
		return nil, err
	}

	// 如果缓存中没有，从数据库中查询
	var user usermd.User
	err = dao.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在，返回 nil, nil
		}
		klog.Error("数据库查询失败:", err)
		return nil, err // 其他错误，返回具体错误信息
	}

	// 将查询到的用户信息存入 Redis 缓存
	userJSON, err := json.Marshal(user)
	if err != nil {
		klog.Error("用户信息序列化失败，Email:", email, " 错误信息:", err)
		return nil, fmt.Errorf("用户信息序列化失败: %v", err)
	}

	err = dao.cache.Set(context.Background(), cacheKey, userJSON, time.Hour).Err()
	if err != nil {
		klog.Error("设置 Redis 缓存失败:", err)
	}

	return &user, nil
}

func (dao *UserDAO) FindByUsername(username string) (*usermd.User, error) {
	cacheKey := fmt.Sprintf("users:username:%s", username)
	val, err := dao.cache.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var cachedUser usermd.User
		err := json.Unmarshal([]byte(val), &cachedUser)
		if err != nil {
			klog.Error("用户信息反序列化失败，Username:", username, " 错误信息:", err)
			return nil, fmt.Errorf("用户信息反序列化失败: %v", err)
		}
		return &cachedUser, nil
	}
	if err != redis.Nil {
		klog.Error("Redis 查询失败:", err)
		return nil, err
	}

	var user usermd.User
	err = dao.db.Where("user_name = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在，返回 nil, nil
		}
		klog.Error("数据库查询失败:", err)
		return nil, err // 其他错误，返回具体错误信息
	}

	err = dao.cache.Set(context.Background(), cacheKey, user, time.Hour).Err()
	if err != nil {
		klog.Error("设置 Redis 缓存失败:", err)
	}

	return &user, nil
}

// GetUserPermissions 根据用户ID查询用户权限
func (dao *UserDAO) GetUserPermissions(ctx context.Context, userid int64) (int32, error) {
	user, err := dao.FindOne(ctx, userid)
	if err != nil {
		return -1, err
	}
	if user == nil {
		return -1, fmt.Errorf("用户不存在")
	}
	return user.UserPermissions, nil
}

