package models

import "github.com/jinzhu/gorm"

const (
	Unknown  int = 0
	Archive_ int = 1 // 文章评论
	Novel_   int = 2 // 小说评论
)

type Comment struct {
	gorm.Model
	Title    string
	AuthorID uint
	Author   User `gorm:"ForeignKey:AuthorID;"`
	Type     int  //评论类型
	RId      uint // 关联的ID
	RCid     uint // 是否评论中评论
}

func String2IntComment(Type string) int {
	if Type == "archive" {
		return Archive_
	} else if Type == "novel" {
		return Novel_
	} else {
		return Unknown
	}
}
func GetComment(ID interface{}) (Comment, error) {
	var comment Comment
	result := DB.First(&comment, ID)
	return comment, result.Error
}
