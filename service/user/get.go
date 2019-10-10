package user

import (
	"eroauz/models"
	"eroauz/serializer"
)

type GetService struct {
	ID     uint `json:"id" form:"id" param:"id"`
	result models.User
}

// EroAPI godoc
// @Summary 获取用户详细
// @Description 如果ID为0就是查看自己
// @Tags user
// @Accept html
// @Produce json
// @Success 200 {object} serializer.UserResponse
// @Failure 500 {object} serializer.Response
// @Param id path int true "用户ID"
// @Router /api/v1/user/:id [get]
func (service *GetService) Get(create uint) *serializer.Response {
	var user models.User
	if service.ID == 0 {
		service.ID = create
	}
	if err := models.DB.Where("ID = ?", service.ID).First(&user).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "获取失败",
			Error:  err.Error(),
		}
	}
	service.result = user
	return nil
}
func (service *GetService) Response() interface{} {
	return serializer.BuildUserResponse(service.result)

}
