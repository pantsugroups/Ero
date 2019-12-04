package models

import "github.com/jinzhu/gorm"

type Archive struct {
	gorm.Model
	Cover          string
	Title          string
	JapTitle       string
	Author         string
	Content        string `gorm:"type:text"`
	PrimaryContent string `gorm:"type:text"`
	Create         User   `gorm:"ForeignKey:CreateID;"`
	CreateID       uint
	Tag            string
	Pass           bool
}

func GetArchive(ID interface{}) (Archive, error) {
	var archive Archive
	result := DB.First(&archive, ID)
	return archive, result.Error
}
