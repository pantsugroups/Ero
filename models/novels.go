package models

import "github.com/jinzhu/gorm"


type Novel struct{
	gorm.Model
	Title string
	Author string
	Cover string
	Description string
	Subscribed int
	Ended bool
	Level int
}
const(
	// 普通用户都可以查看的程度
	Level1 = 0
	// 只有正式会员才能看的东西
	Level2 = 1
	// 只有老司机才能看的
	Level3 = 2
)