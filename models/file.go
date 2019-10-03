package models

import "github.com/jinzhu/gorm"

type File struct {
	gorm.Model
	Path string
	Type int
}

const (
	Volume_ int = 1
	Image_  int = 2
	Other_  int = 3
)

func GetFile(ID interface{}) (File, error) {
	var file File
	result := DB.First(&file, ID)
	return file, result.Error
}
