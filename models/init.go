package models

import (
	"time"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 初始化mysql链接
func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	DB.AutoMigrate(&User{}, &Archive{}, Comment{}, Category{})
	// 初始化
	var count int
	if err := DB.Model(&Category{}).Count(&count).Error; err != nil {
		panic(err) // 安装错误
	}
	if count == 0 {
		category := Category{
			Title: "Default",
			Count: 0,
		}
		if err := DB.Create(&category).Error; err != nil {
			panic(err)
		}
	}

}
