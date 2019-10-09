package message

import (
	"eroauz/models"
	"eroauz/serializer"
)

type GetService struct {
	ID     uint `json:"id" form:"id" param:"id" null:"false"`
	result models.Message
}

// EroAPI godoc
// @Summary 获取消息详细
// @Description 接收者必须为自己或者权限必须为管理员
// @Tags message
// @Accept html
// @Produce json
// @Success 200 {object} serializer.MessageResponse
// @Failure 500 {object} serializer.Response
// @Param id path integer true "消息ID"
// @Router /api/v1/message/:id [get]
// @Security ApiKeyAuth
func (service *GetService) Get(create uint) *serializer.Response {
	var message models.Message
	if err := models.DB.Where("ID = ?", service.ID).First(&message).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "获取失败",
			Error:  err.Error(),
		}
	}
	var send models.User
	u, err := models.GetUser(message.SendID)
	send = u
	if err != nil {
		send.Nickname = "已删除用户"
	}
	message.Send = send
	var recv models.User
	r, err := models.GetUser(message.SendID)
	recv = r
	if err != nil {
		send.Nickname = "已删除用户"
	}
	message.Recv = recv

	service.result = message
	return nil
}
func (service *GetService) Response() interface{} {
	return serializer.BuildMessageResponse(service.result)

}
