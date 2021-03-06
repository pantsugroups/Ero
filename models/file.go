package models

import "github.com/jinzhu/gorm"

type File struct {
	gorm.Model
	User     User `gorm:"ForeignKey:UserID;"`
	UserID   uint
	Path     string
	FileName string
	Type     int
}

const (
	Remote_ int = 0
	Volume_ int = 1
	Image_  int = 2
	Other_  int = 3

)

func GetFile(ID interface{}) (File, error) {
	var file File
	result := DB.First(&file, ID)
	return file, result.Error
}
