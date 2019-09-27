package models

type NovelSubscribe struct{
	User   User   `gorm:"ForeignKey:NovelSubscribe"`
	Novel Novel  `gorm:"ForeignKey:NovelSubscribe"`
}