package archive

import (
	"eroauz/models"
	"eroauz/serializer"
)

type CreateService struct{
	Title  string `json:"title" form:"title"`
	JapTitle string `json:"japanese_title" form:"japanese_title"`
	Author   string `json:"author" form:"author"`
	Content  string `json:"content" form:"content"`
	PrimaryContent string `json:"primary_content" form:"primary_content"`
	Cover string `json:"cover" form:"cover"`
}
func (service CreateService)Create()(models.Archive,*serializer.Response){
	archive:= models.Archive{
		Title:service.Title,
		JapTitle:service.JapTitle,
		Content:service.Content,
		Author:service.Author,
		PrimaryContent:service.PrimaryContent,
		Cover:service.Cover,
	}
	if err := models.DB.Create(&archive).Error; err != nil {
		return archive, &serializer.Response{
			Status: 40004,
			Msg:    "创建失败",
		}
	}
	return archive,nil
}