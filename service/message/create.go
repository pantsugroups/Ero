package message

import (
	"eroauz/models"
	"eroauz/serializer"
)

type CreateService struct {
	Title string `json:"title" form:"title"`
	//Send   uint   `json:"send" form:"send"`
	Recv   uint `json:"recv" form:"recv"`
	result models.Message
}

func (service *CreateService) Create(creater uint) *serializer.Response {
	u, err := models.GetUser(creater)
	if err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "寻找匹配ID失败",
		}
	}
	r, err := models.GetUser(service.Recv)
	if err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "寻找匹配ID失败",
		}
	}
	msg := models.Message{
		Send:  u,
		Recv:  r,
		Title: service.Title,
		Read:  false,
	}
	if err := models.DB.Create(&msg).Error; err != nil {
		return &serializer.Response{
			Status: 40004,
			Msg:    "创建失败",
		}
	}
	service.result = msg
	return nil
}
func (service *CreateService) Response() interface{} {
	return serializer.BuildMessageResponse(service.result)

}
