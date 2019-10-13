package models

import (
	"eroauz/conf"
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
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	DB.Debug().AutoMigrate(&User{}, &Archive{}, &Novel{}, &Volume{}, &Comment{}, &Category{},
		&Message{}, &File{}, &NovelCategory{}, NovelSubscribe{}, ArchiveCategory{})
	// 初始化
	var count int
	if err := DB.Model(&Category{}).Count(&count).Error; err != nil {
		panic(err) // 安装错误
	}
	if count == 0 {
		category := Category{
			Title: "Default",
			Type:  1,
			Count: 0,
		}
		if err := DB.Create(&category).Error; err != nil {
			panic(err)
		}
		category = Category{
			Title: "Default",
			Type:  2,
			Count: 0,
		}
		if err := DB.Create(&category).Error; err != nil {
			panic(err)
		}
	}
	count = 0
	DB.Model(&User{}).Where("user_name = ?", "admin").Count(&count)
	if count == 0 { // 管理员不存在
		user := User{
			Nickname: "admin",
			UserName: "admin",
			Status:   Admin,
			Point:    250,
			Avatar:   conf.DefaultAvatar,
		}
		if err := user.SetPassword("qwq123"); err != nil {
			panic(err)
		}
		if err := DB.Create(&user).Error; err != nil {
			panic(err)
		}
	}

}
