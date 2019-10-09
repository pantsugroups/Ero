package novel

import (
	"eroauz/models"
	"eroauz/serializer"
)

type DeleteService struct {
	ID uint `json:"ID" form:"ID" param:"id" null:"false"`
}

// EroAPI godoc
// @Summary 删除小说
// @Description 接收者必须为自己或者用户权限为管理员
// @Tags novel
// @Accept html
// @Produce json
// @Success 200 {object} serializer.Response
// @Failure 500 {object} serializer.Response
// @Param id path integer true "小说ID"
// @Router /api/v1/novel/:id [delete]
// @Security ApiKeyAuth
func (service *DeleteService) Delete(create uint) *serializer.Response {
	var novel models.Novel
	if err := models.DB.Where("ID = ?", service.ID).First(&novel).Error; err != nil {
		return &serializer.Response{
			Status: 404,
			Msg:    "寻找匹配ID失败",
			Error:  err.Error(),
		}
	}
	if err := models.DB.Delete(&novel).Error; err != nil {
		return &serializer.Response{
			Status: 500,
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
