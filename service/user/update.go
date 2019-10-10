package user

import (
	"eroauz/models"
	"eroauz/serializer"
)

type UpdateService struct {
	Nickname string `json:"nickname" form:"nickname"`
	Avatar   string `json:"avatar" form:"avatar"`
	Hito     string `json:"hito" form:"hito"`
	Bio      string `json:"bio" form:"bio"`
	Website  string `json:"website" form:"website"`
	result   models.User
}

// EroAPI godoc
// @Summary 更新用户信息
// @Description
// @Tags user
// @Accept html
// @Produce json
// @Success 200 {object} serializer.UserResponse
// @Failure 500 {object} serializer.Response
// @Param nickname formData string false "昵称"
// @Param avatar formData string false "用户头像"
// @Router /api/v1/user/ [put]
// @Security ApiKeyAuth
func (service *UpdateService) Update(create uint) *serializer.Response {
	var user models.User
	if err := models.DB.Where("ID = ?", create).First(&user).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "获取失败",
		}
	}
	if err := models.DB.Model(&user).Update(models.User{
		Nickname: service.Nickname,
		Avatar:   service.Avatar,
		Hito:     service.Hito,
		Website:  service.Website,
		Bio:      service.Bio,
	}).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "获取失败",
		}
	}
	service.result = user
	return nil
}
func (service *UpdateService) Response() interface{} {
	return serializer.BuildUserResponse(service.result)
}
