package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Title string
	Count int
}

func GetCategory(ID interface{}) (Category, error) {
	var c Category
	result := DB.First(&c, ID)
	return c, result.Error
}
