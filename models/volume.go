package models

import "github.com/jinzhu/gorm"

type Volume struct {
	gorm.Model
	Title string
	Cover string
	Novel Novel `gorm:"ForeignKey:Volume"`
	Files string
}