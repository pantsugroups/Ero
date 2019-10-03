package category

import (
	"eroauz/models"
	"eroauz/serializer"
)

type UpdateService struct {
	ID     uint   `json:"id" form:"id" param:"id" null:"false"`
	Title  string `json:"title" form:"title"`
	result models.Category
}

func (service *UpdateService) Update(create uint) *serializer.Response {
	var category models.Category
	if err := models.DB.Where("ID = ?", service.ID).First(&category).Error; err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	if err := models.DB.Model(&category).Update(models.Novel{
		Title: service.Title,
	}); err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	service.result = category
	return nil
}
func (service *UpdateService) Response() interface{} {
	return serializer.BuildCategoryResponse(service.result)
}
