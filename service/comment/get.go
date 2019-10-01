package comment

import (
	"eroauz/models"
	"eroauz/serializer"
)

type GetService struct {
	ID     uint `json:"id" form:"id" param:"id" null:"false"`
	result models.Comment
}

func (service *GetService) Get() *serializer.Response {
	var archive models.Comment
	if err := models.DB.Where("ID = ?", service.ID).First(&archive).Error; err != nil {
		return &serializer.Response{
			Status: 40003,
			Msg:    "获取失败",
			Error:  err.Error(),
		}
	}
	service.result = archive
	return nil
}
func (service *GetService) Response() interface{} {
	return serializer.BuildCommentResponse(service.result)

}
