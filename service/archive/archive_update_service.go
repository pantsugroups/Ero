package archive

import (
	"eroauz/models"
	"eroauz/serializer"
)

type UpdateService struct{
	ID  uint `json:"id" form:"id" null:"false"`
	Title  string `json:"title" form:"title"`
	JapTitle string `json:"japanese_title" form:"japanese_title"`
	Cover  string `json:"cover" form:"cover"`
	Content string `json:"content" form:"content"`
	Author string `json:"author" form:"author"`
	Mail   string `json:"mail" form:"mail"`
	PrimaryContent string `json:"primary_content" form:"primary_content"`
}
func (service *UpdateService)Update()(models.Archive,*serializer.Response){
	var archive models.Archive
	if err := models.DB.Where("ID = ?",service.ID).First(&archive).Error;err != nil{
		return archive, &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	if err := models.DB.Model(&archive).Update(models.Archive{
		Title:service.Title,
		JapTitle:service.JapTitle,
		Cover:service.Cover,
		Content:service.Content,
		Author:service.Author,
		PrimaryContent:service.PrimaryContent,
	});err != nil{
		return archive, &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	return archive,nil
}