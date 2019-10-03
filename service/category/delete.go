package category

import (
	"eroauz/models"
	"eroauz/serializer"
)

type DeleteService struct {
	ID uint `json:"ID" form:"ID" param:"id" null:"false"`
}

func (service *DeleteService) Delete(create uint) *serializer.Response {
	var category models.Category
	if err := models.DB.Where("ID = ?", service.ID).First(&category).Error; err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "寻找匹配ID失败",
		}
	}
	if err := models.DB.Delete(&category).Error; err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "删除失败",
		}
	}
	return nil
}
func (service *DeleteService) Response() interface{} {
	return serializer.Response{
		Status: 0,
		Msg:    "成功",
	}
}
