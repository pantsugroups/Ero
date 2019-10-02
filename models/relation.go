package models

type NovelSubscribe struct {
	User  User  `gorm:"ForeignKey:NovelSubscribe"`
	Novel Novel `gorm:"ForeignKey:NovelSubscribe"`
}
type NovelCategory struct {
	Novel    Novel    `gorm:"ForeignKey:NovelCategory"`
	Category Category `gorm:"ForeignKey:NovelCategory"`
}
