package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var db *gorm.DB

func DBConnection() *sql.DB{
	dsn := GetDBConfig()
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("DB Error: %w",err))
	}
	CreateTable(db)
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	return sqlDB
}
// DBのdsnを取得する
func GetDBConfig() string {
	dsn := fmt.Sprintf("user:password@tcp(localhost:3306)/db?charset=utf8mb4&parseTime=True&loc=Local")
	return dsn
}

// Task型のテーブルを作成する
func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&Task{})
}
