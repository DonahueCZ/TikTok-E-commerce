package mysql

import (
	"log"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 全局数据库实例
var (
	DB   *gorm.DB
	once sync.Once
)

// InitDB 初始化 MySQL 连接
func InitDB(dsn string) {
	once.Do(func() {
		var err error
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("数据库连接失败: %v", err)
		}

		// 获取底层 sql.DB
		sqlDB, err := DB.DB()
		if err != nil {
			log.Fatalf("获取 sql.DB 失败: %v", err)
		}

		// 配置连接池
		sqlDB.SetMaxOpenConns(100)          // 最大连接数
		sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
		sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大生命周期

		log.Println("MySQL 连接初始化成功！")
	})
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("数据库未初始化！请先调用 InitDB()")
	}
	return DB
}
