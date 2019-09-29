package novel

import (
	"eroauz/models"
	"eroauz/serializer"
)

type GetService struct {
	ID     uint `json:"id" form:"id" null:"false"`
	result models.Novel
}

func (service *GetService) Get() *serializer.Response {
	var novel models.Novel
	if err := models.DB.Where("ID = ?", service.ID).First(&novel).Error; err != nil {
		return &serializer.Response{
			Status: 40003,
			Msg:    "获取失败",
		}
	}
	service.result = novel
	return nil
}
func (service *GetService) Response() interface{} {
	return serializer.BuildNovelResponse(service.result)
}
