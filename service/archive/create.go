package archive

import (
	"eroauz/models"
	"eroauz/serializer"
)

type CreateService struct {
	Title          string `json:"title" form:"title"`
	JapTitle       string `json:"japanese_title" form:"japanese_title"`
	Author         string `json:"author" form:"author"`
	Content        string `json:"content" form:"content"`
	PrimaryContent string `json:"primary_content" form:"primary_content"`
	Cover          string `json:"cover" form:"cover"`
	result         models.Archive
}

func (service *CreateService) Create() *serializer.Response {
	archive := models.Archive{
		Title:          service.Title,
		JapTitle:       service.JapTitle,
		Content:        service.Content,
		Author:         service.Author,
		PrimaryContent: service.PrimaryContent,
		Cover:          service.Cover,
	}
	service.result = archive
	if err := models.DB.Create(&archive).Error; err != nil {
		return &serializer.Response{
			Status: 40004,
			Msg:    "创建失败",
		}
	}
	return nil
}
func (service *CreateService) Response() *serializer.Response {
	return &serializer.Response{
		Status: 0,
		Data:   serializer.BuildArchiveResponse(service.result),
	}
}
