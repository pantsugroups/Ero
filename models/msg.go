package models

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	Title string

	Send   User `gorm:"ForeignKey:SendID;"` // 为空则是系统消息
	Recv   User `gorm:"ForeignKey:RecvID;"`
	Read   bool
	SendID uint
	RecvID uint
}

func GetMessage(ID interface{}) (Message, error) {
	var msg Message
	result := DB.First(&msg, ID)
	return msg, result.Error
}
