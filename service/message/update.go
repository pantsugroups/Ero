package message

import (
	"eroauz/models"
	"eroauz/serializer"
)

type UpdateService struct {
	ID     uint `json:"id" param:"id" form:"id" null:"false"`
	Read   bool `json:"is_read" form:"id_read"`
	result models.Message
}

// EroAPI godoc
// @Summary 更新消息
// @Description 必须为管理员或者接收者为自己
// @Tags message
// @Accept html
// @Produce json
// @Success 200 {object} serializer.MessageResponse
// @Failure 500 {object} serializer.Response
// @Param id path integer false "消息ID"
// @Param read formData boolean true "是否已读"
// @Router /api/v1/message/:id [put]
// @Security ApiKeyAuth
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
