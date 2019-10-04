package user

import (
	"eroauz/models"
	"eroauz/serializer"
)

type GetService struct {
	ID     uint `json:"id" form:"id" param:"id" null:"false"`
	result models.User
}

// EroAPI godoc
// @Summary 获取用户详细
// @Description
// @Tags user
// @Accept html
// @Produce json
// @Success 200 {object} serializer.UserResponse
// @Failure 500 {object} serializer.Response
// @Param id path int true "用户ID"
// @Router /api/v1/user/:id [get]
func (service *GetService) Get(create uint) *serializer.Response {
	var user models.User
	if err := models.DB.Where("ID = ?", service.ID).First(&user).Error; err != nil {
		return &serializer.Response{
			Status: 40003,
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
