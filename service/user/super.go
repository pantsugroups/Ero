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

// EroAPI godoc
// @Summary 超级：更新用户信息
// @Description 必须为管理
// @Tags user,admin
// @Accept html
// @Produce json
// @Success 200 {object} serializer.UserResponse
// @Failure 500 {object} serializer.Response
// @Param id path int false "用户ID"
// @Param nickname formData string false "昵称"
// @Param avatar formData string false "用户头像"
// @Param point formData integer false "下载点数"
// @Param status formData string false "用户状态：driver，active，inactive，suspend，admin"
// @Router /api/v1/admin/user/:id [put]
// @Security ApiKeyAuth
func (service *SuperUpdateService) Update(create uint) *serializer.Response {
	var user models.User
	if err := models.DB.Where("ID = ?", service.ID).First(&user).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "获取失败",
		}
	}
	u := models.User{
		Nickname: service.Nickname,
		Avatar:   service.Avatar,
		Point:    service.Point,
		Mail:     service.Mail,
		Status:   service.Status,
	}
	u.CheckAvatar()
	if err := models.DB.Model(&user).Update(u); err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "获取失败",
		}
	}
	service.result = user
	return nil
}
func (service *SuperUpdateService) Response() interface{} {
	return serializer.BuildUserResponse(service.result)
}
