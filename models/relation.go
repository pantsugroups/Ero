package models

import "github.com/jinzhu/gorm"

type NovelSubscribe struct {
	gorm.Model
	UserID  uint
	User    User `gorm:"ForeignKey:UserID"`
	NovelID uint
	Novel   Novel `gorm:"ForeignKey:NovelID"`
}
type NovelCategory struct {
	gorm.Model
	NovelID    uint
	Novel      Novel `gorm:"ForeignKey:NovelID"`
	CategoryID uint
	Category   Category `gorm:"ForeignKey:CategoryID"`
}
type ArchiveCategory struct {
	gorm.Model
	ArchiveID  uint
	Archive    Archive `gorm:"ForeignKey:ArchiveID"`
	CategoryID uint
	Category   Category `gorm:"ForeignKey:CategoryID"`
}
