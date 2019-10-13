package relationship

import (
	"eroauz/models"
	"eroauz/serializer"
)

type DeleteA2CService struct {
	Archive  uint `json:"archive" form:"archive" null:"false"`
	Category uint `json:"category" form:"category" null:"false"`
}

// EroAPI godoc
// @Summary 删除文章分类关联
// @Description 接收者必须为自己管理员
// @Tags archive,category,admin
// @Accept html
// @Produce json
// @Success 200 {object} serializer.Response
// @Failure 500 {object} serializer.Response
// @Param novel formData integer true "小说ID"
// @Param category formData integer true "分类ID"
// @Router /api/v1/category/archive/ [delete]
// @Security ApiKeyAuth
func (service *DeleteA2CService) Delete(create uint) *serializer.Response {

	a, err := models.GetArchive(service.Archive)
	if err != nil {
		return &serializer.Response{
			Status: 404,
			Msg:    "找不到Archive ID",
		}
	}
	c, err := models.GetCategory(service.Category)
	if err != nil {
		return &serializer.Response{
			Status: 404,
			Msg:    "找不到Category ID",
		}
	}
	target := models.ArchiveCategory{
		Archive:  a,
		Category: c,
	}
	if err := models.DB.Delete(&target).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "删除失败",
		}
	}
	return nil
}
func (service *DeleteA2CService) Response() interface{} {
	return serializer.Response{
		Status: 0,
		Msg:    "成功",
	}
}
