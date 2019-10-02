package models

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	Title string
	Send  User // 为空则是系统消息
	Recv  User
	Read  bool
}

func GetMessage(ID interface{}) (Message, error) {
	var msg Message
	result := DB.First(&msg, ID)
	return msg, result.Error
}
