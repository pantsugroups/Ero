package models

import (
	"eroauz/conf"
	"github.com/jinzhu/gorm"
	"path"
)

const (
	Locao_Volume  int = 0
	Remote_Volume int = 1
)

type Volume struct {
	gorm.Model
	Title string
	Cover string

	Novel   Novel `gorm:"ForeignKey:NovelID"`
	NovelID uint

	File   File `gorm:"ForeignKey:FileID"`
	FileID uint

	ZIndex int

	Type int
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
