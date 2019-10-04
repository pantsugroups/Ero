package models

import "github.com/jinzhu/gorm"

type Archive struct {
	gorm.Model
	Cover          string
	Title          string
	JapTitle       string
	Author         string
	Content        string
	PrimaryContent string
	Create         User `gorm:"ForeignKey:CreateID;"`
	CreateID       uint
	Tag            string
}

func GetArchive(ID interface{}) (Archive, error) {
	var archive Archive
	result := DB.First(&archive, ID)
	return archive, result.Error
}
