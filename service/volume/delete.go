package volume

import (
	"eroauz/models"
	"eroauz/serializer"
)

type DeleteService struct {
	ID uint `json:"ID" form:"ID" param:"id" null:"false"`
}

// EroAPI godoc
// @Summary 删除分卷
// @Description 必须为管理员
// @Tags volume,admin
// @Accept html
// @Produce json
// @Success 200 {object} serializer.Response
// @Failure 500 {object} serializer.Response
// @Param id path integer true "小说ID"
// @Router /api/v1/volume/:id [delete]
// @Security ApiKeyAuth
func (service *DeleteService) Delete(create uint) *serializer.Response {
	var volume models.Volume
	if err := models.DB.Where("ID = ?", service.ID).First(&volume).Error; err != nil {
		return &serializer.Response{
			Status: 404,
			Msg:    "寻找匹配ID失败",
		}
	}
	if err := models.DB.Delete(&volume).Error; err != nil {
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
