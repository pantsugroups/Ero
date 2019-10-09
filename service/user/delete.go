package user

import (
	"eroauz/models"
	"eroauz/serializer"
)

type DeleteService struct {
	ID uint `json:"ID" form:"ID" param:"id" null:"false"`
}

// EroAPI godoc
// @Summary 删除用户
// @Description 必须为管理员
// @Tags user,admin
// @Accept html
// @Produce json
// @Success 200 {object} serializer.Response
// @Failure 500 {object} serializer.Response
// @Param id path integer true "用户ID"
// @Router /api/v1/user/:id [delete]
// @Security ApiKeyAuth
func (service *DeleteService) Delete(create uint) *serializer.Response {
	var user models.User
	if err := models.DB.Where("ID = ?", service.ID).First(&user).Error; err != nil {
		return &serializer.Response{
			Status: 404,
			Msg:    "寻找匹配ID失败",
		}
	}
	if err := models.DB.Delete(&user).Error; err != nil {
		return &serializer.Response{
			Status: 404,
			Msg:    "删除失败",
		}
	}
	return nil
}
func (service *DeleteService) Response() interface{} {
	return serializer.Response{
		Status: 0,
		Msg:    "成功",
	}
}
