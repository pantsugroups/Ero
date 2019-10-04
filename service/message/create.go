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

// EroAPI godoc
// @Summary 创建消息
// @Description 有这个接口的实现，但是暂未启用.jpg
// @Tags message
// @Accept html
// @Produce json
// @Success 200 {object} serializer.MessageResponse
// @Failure 500 {object} serializer.Response
// @Param title formData string true "消息内容"
// @Param recv formData integer true "回复用户 ID"
// @Router /api/v1/message/ [post]
// @Security ApiKeyAuth
func (service *CreateService) Create(create uint) *serializer.Response {
	u, err := models.GetUser(create)
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
