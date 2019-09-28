package novel

import (
	"eroauz/models"
	"eroauz/serializer"
)

type CreateService struct {
	Title       string `json:"title" form:"title"`
	Author      string `json:"author" form:"title"`
	Cover       string `json:"cover" form:"cover"`
	Description string `json:"description" form:"description"`
}

func (service *CreateService) Create() (models.Novel, *serializer.Response) {
	novel := models.Novel{
		Title:       service.Title,
		Author:      service.Author,
		Cover:       service.Cover,
		Description: service.Description,
		Ended:       false,
		Level:       models.Level1,
		Subscribed:  0,
	}
	if err := models.DB.Create(&novel).Error; err != nil {
		return novel, &serializer.Response{
			Status: 40007,
			Msg:    "创建失败",
		}
	}
	return novel, nil
}
