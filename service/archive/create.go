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

func (service *CreateService) Create(create uint) *serializer.Response {
	user, _ := models.GetUser(create)
	archive := models.Archive{
		Title:          service.Title,
		JapTitle:       service.JapTitle,
		Content:        service.Content,
		Author:         service.Author,
		PrimaryContent: service.PrimaryContent,
		Cover:          service.Cover,
		Create:         user,
	}
	if err := models.DB.Create(&archive).Error; err != nil {
		return &serializer.Response{
			Status: 40004,
			Msg:    "创建失败",
		}
	}
	service.result = archive
	return nil
}
func (service *CreateService) Response() interface{} {
	return serializer.BuildArchiveResponse(service.result)

}
