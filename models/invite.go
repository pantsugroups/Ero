package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Invite struct {
	gorm.Model
	Code      string
	TimeLimit time.Time
	Create    uint
}

func GetInvite(ID interface{}) (Invite, error) {
	var i Invite
	result := DB.First(&i, ID)
	return i, result.Error
}
