package archive

import (
	"eroauz/models"
	"eroauz/serializer"
)

type DeleteService struct {
	ID uint `json:"id" form:"id" param:"id" null:"false"`
}

// EroAPI godoc
// @Summary 删除文章
// @Description 必须要登陆且创建者为自己或者管理员
// @Tags archive
// @Accept html
// @Produce json
// @Success 200 {object} serializer.Response
// @Failure 500 {object} serializer.Response
// @Param ID path integer true "文章ID"
// @Router /api/v1/archive/:id [delete]
// @Security ApiKeyAuth
func (service *DeleteService) Delete(create uint) *serializer.Response {
	var archive models.Archive
	if err := models.DB.Where("ID = ?", service.ID).First(&archive).Error; err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "寻找匹配ID失败",
			Error:  err.Error(),
		}
	}
	if err := models.DB.Delete(&archive).Error; err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "删除失败",
			Error:  err.Error(),
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
