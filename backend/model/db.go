package model

import (
	"database/sql"
	"fmt"
	"os"

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
	// dsn := fmt.Sprintf(
	// 	"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASS"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_NAME"),
	// )
	dsn := "user:password@tcp(192.168.0.0:3306)//db?charset=utf8&parseTime=true"
	return dsn
}

// Task型のテーブルを作成する
func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&Task{})
}
