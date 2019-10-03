package archive

import (
	"eroauz/models"
	"eroauz/serializer"
)

type UpdateService struct {
	ID             uint   `json:"id" form:"id" param:"id" null:"false"`
	Title          string `json:"title" form:"title"`
	JapTitle       string `json:"japanese_title" form:"japanese_title"`
	Cover          string `json:"cover" form:"cover"`
	Content        string `json:"content" form:"content"`
	Author         string `json:"author" form:"author"`
	Mail           string `json:"mail" form:"mail"`
	PrimaryContent string `json:"primary_content" form:"primary_content"`
	result         models.Archive
}

func (service *UpdateService) Update(create uint) *serializer.Response {
	var archive models.Archive
	if err := models.DB.Where("ID = ?", service.ID).First(&archive).Error; err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	if err := models.DB.Model(&archive).Update(models.Archive{
		Title:          service.Title,
		JapTitle:       service.JapTitle,
		Cover:          service.Cover,
		Content:        service.Content,
		Author:         service.Author,
		PrimaryContent: service.PrimaryContent,
	}); err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	service.result = archive
	return nil
}
func (service *UpdateService) Response() interface{} {
	return service.result

}
