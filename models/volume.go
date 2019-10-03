package models

import (
	"eroauz/conf"
	"github.com/jinzhu/gorm"
	"path"
)

type Volume struct {
	gorm.Model
	Title string
	Cover string

	Novel   Novel `gorm:"ForeignKey:NovelID"`
	NovelID uint

	File   File `gorm:"ForeignKey:FileID"`
	FileID uint
}

func (volume *Volume) CheckCover() {
	if volume.Cover == "" {
		volume.Cover = path.Join(conf.StaticPath, "cover.jpg")
	}
}
func GetVolume(ID interface{}) (Volume, error) {
	var volume Volume
	result := DB.First(&volume, ID)
	return volume, result.Error
}
