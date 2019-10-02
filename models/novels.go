package models

import "github.com/jinzhu/gorm"

type Novel struct {
	gorm.Model
	Title       string
	Author      string
	Cover       string
	Description string
	Subscribed  int
	Ended       bool //是否完结
	Level       int
	Creater     User
}

const (
	// 普通用户都可以查看的程度
	Level1 = 0
	// 只有正式会员才能看的东西
	Level2 = 1
	// 只有老司机才能看的
	Level3 = 2
)

func GetNovel(ID interface{}) (Novel, error) {
	var novel Novel
	result := DB.First(&novel, ID)
	return novel, result.Error
}

//func Int2String_Novel(Type int) string{
//	if Type == Level1{
//		return "Level1"
//	}else if Type == Level2{
//		return "Level2"
//	}else if Type == Level3{
//		return "Level3"
//	}else{
//		return "Level1"
//	}
//}
