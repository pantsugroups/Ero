package comment

import (
	"eroauz/models"
	"eroauz/serializer"
)

type DeleteService struct {
	ID uint `json:"id" form:"id" param:"id" null:"false"`
}

func (service *DeleteService) Delete() *serializer.Response {
	var comment models.Comment
	if err := models.DB.Where("ID = ?", service.ID).First(&comment).Error; err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "寻找匹配ID失败",
			Error:  err.Error(),
		}
	}
	if err := models.DB.Delete(&comment).Error; err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "删除失败",
			Error:  err.Error(),
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
