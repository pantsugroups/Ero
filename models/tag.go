package models

import (
	//"github.com/agtorre/go-cookbook/chapter13/vendoring/models"
	"github.com/jinzhu/gorm"
)

const (
	N = 1
	// Novel
	A = 2
	// Archive
	U = 0

// Unknown
)

type Tag struct {
	gorm.Model
	Title string
	RID   int //Relation ID
	RType int //Relation Type
}

// Create The Tags
func Create(title string, ID int, Type int) (Tag, error) {
	tag := Tag{
		Title: title,
		RID:   ID,
		RType: Type,
	}
	if err := DB.Create(&tag).Error; err != nil {
		return tag, err
	} else {
		return tag, nil
	}
}

//func Int2StringTag(Type int) string {
//	if Type == N {
//		return "Novel"
//	} else if Type == A {
//		return "Archive"
//	} else if Type == U {
//		return "Unknown"
//	} else {
//		return "Unknown"
//	}
//}
//
//func String2intTag(Type string) int {
//	if Type == "Novel" {
//		return N
//	} else if Type == "Archive" {
//		return A
//	} else {
//		return U
//	}
//}
