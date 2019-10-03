package message

import (
	"eroauz/models"
	"eroauz/serializer"
)

type UpdateService struct {
	ID     uint `json:"id" param:"id" form:"id"`
	Read   bool `json:"is_read" form:"id_read"`
	result models.Message
}

func (service *UpdateService) Update(create uint) *serializer.Response {
	var msg models.Message
	if err := models.DB.Where("ID = ?", service.ID).First(&msg).Error; err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	if err := models.DB.Model(&msg).Update(models.Message{
		Read: service.Read,
	}); err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	service.result = msg
	return nil
}
func (service *UpdateService) Response() interface{} {
	return serializer.BuildMessageResponse(service.result)
}
