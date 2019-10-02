package category

import (
	"eroauz/models"
	"eroauz/serializer"
)

type CreateService struct {
	Title  string `json:"title" form:"title"`
	result models.Category
}

func (service *CreateService) Create(creater uint) *serializer.Response {
	category := models.Category{
		Title: service.Title,
		Count: 0,
	}
	if err := models.DB.Create(&category).Error; err != nil {
		return &serializer.Response{
			Status: 40007,
			Msg:    "创建失败",
		}
	}
	service.result = category
	return nil
}
func (service *CreateService) Response() interface{} {
	return serializer.BuildCategoryResponse(service.result)
}
