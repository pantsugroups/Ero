package message

import (
	"eroauz/models"
	"eroauz/serializer"
)

type GetService struct {
	ID     uint `json:"id" form:"id" param:"id" null:"false"`
	result models.Message
}

func (service *GetService) Get(create uint) *serializer.Response {
	var message models.Message
	if err := models.DB.Where("ID = ?", service.ID).First(&message).Error; err != nil {
		return &serializer.Response{
			Status: 40003,
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
