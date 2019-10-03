package user

import (
	"eroauz/models"
	"eroauz/serializer"
)

type SuperUpdateService struct {
	ID       uint   `json:"id" form:"id" param:"id" null:"false"`
	Nickname string `json:"nickname" form:"nickname"`
	Avatar   string `json:"avatar" form:"avatar"`
	Point    int    `json:"point" form:"point"`
	Mail     string `json:"mail" form:"mail"`
	Status   string `json:"status" form:"status"`
	result   models.User
}

func (service *SuperUpdateService) Update(create uint) *serializer.Response {
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
		Point:    service.Point,
		Mail:     service.Mail,
		Status:   service.Status,
	}); err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	service.result = user
	return nil
}
func (service *SuperUpdateService) Response() interface{} {
	return serializer.BuildUserResponse(service.result)
}
