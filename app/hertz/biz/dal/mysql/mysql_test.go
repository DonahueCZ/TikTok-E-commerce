package mysql

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMySQLConnection(t *testing.T) {
	dsn := "root:612324@tcp(127.0.0.1:4407)/payment_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatalf("数据库连接失败: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("数据库 Ping 失败: %v", err)
	} else {
		fmt.Println("✅ MySQL 连接成功")
	}
}
