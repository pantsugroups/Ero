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
	Volume_ int = 1
	Image_  int = 2
	Other_  int = 3
	Remote_Volume_ int = 4
	Remote_Image_  int = 5
	Remote_Other_  int = 6
)

func GetFile(ID interface{}) (File, error) {
	var file File
	result := DB.First(&file, ID)
	return file, result.Error
}
