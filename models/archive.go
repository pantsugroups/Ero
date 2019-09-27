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
}
// GetUser 用ID获取用户
func GetArchive(ID interface{}) (Archive, error) {
	var archive Archive
	result := DB.First(&Archive{}, ID)
	return archive, result.Error
}
