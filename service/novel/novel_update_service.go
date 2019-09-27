package novel

import (
	"eroauz/models"
	"eroauz/serializer"
)

type UpdateService struct{
	ID uint `json:"id" form:"id" null:"false"`
	Title string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
	Cover string `json:"cover" form:"cover"`
	Ended bool `json:"ended" form:"ended"`
	Level int  `json:"level" form:"level"`
}
func (service *UpdateService)Update()(models.Novel,*serializer.Response){
	var novel models.Novel
	if err := models.DB.Where("ID = ?",service.ID).First(&novel).Error;err != nil{
		return novel, &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	if err := models.DB.Model(&novel).Update(models.Novel{
		Title:service.Title,
		Author:service.Author,
		Cover:service.Cover,
		Ended:service.Ended,
		Level:service.Level,
	});err != nil{
		return novel, &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	return novel,nil
}