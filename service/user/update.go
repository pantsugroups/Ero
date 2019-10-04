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

// EroAPI godoc
// @Summary 更新用户信息
// @Description
// @Tags user
// @Accept html
// @Produce json
// @Success 200 {object} serializer.UserResponse
// @Failure 500 {object} serializer.Response
// @Param id path int false "用户ID"
// @Param nickname formData string false "昵称"
// @Param avatar formData string false "用户头像"
// @Router /api/v1/user/:id [put]
// @Security ApiKeyAuth
func (service *UpdateService) Update(create uint) *serializer.Response {
	var user models.User
	if err := models.DB.Where("ID = ?", service.ID).First(&user).Error; err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	if create != service.ID {
		return &serializer.Response{
			Status: 403,
			Msg:    "没有权限",
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
