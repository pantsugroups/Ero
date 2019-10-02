package user

import (
	"eroauz/models"
	"eroauz/serializer"
)

type UpdateService struct {
	ID       uint   `json:"id" form:"id" param:"id" null:"false"`
	Nickname string `json:"nickname" form:"nickname"`
	Avatar   string `json:"avatar" form:"avatar"`

	result models.User
}

func (service *UpdateService) Update() *serializer.Response {
	var user models.User
	if err := models.DB.Where("ID = ?", service.ID).First(&user).Error; err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	if err := models.DB.Model(&user).Update(models.User{
		Nickname: service.Nickname,
		Avatar:   service.Avatar,
	}); err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	service.result = user
	return nil
}
func (service *UpdateService) Response() interface{} {
	return serializer.BuildUserResponse(service.result)
}
