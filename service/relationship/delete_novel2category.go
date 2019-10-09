package relationship

import (
	"eroauz/models"
	"eroauz/serializer"
)

type DeleteN2CService struct {
	Novel    uint `json:"novel" form:"novel" null:"false"`
	Category uint `json:"category" form:"category" null:"false"`
}

// EroAPI godoc
// @Summary 删除小说分类关联
// @Description 接收者必须为自己管理员
// @Tags novel,category,admin
// @Accept html
// @Produce json
// @Success 200 {object} serializer.Response
// @Failure 500 {object} serializer.Response
// @Param novel formData integer true "小说ID"
// @Param category formData integer true "分类ID"
// @Router /api/v1/novel/:id [delete]
// @Security ApiKeyAuth
func (service *DeleteN2CService) Delete(create uint) *serializer.Response {

	n, err := models.GetNovel(service.Novel)
	if err != nil {
		return &serializer.Response{
			Status: 404,
			Msg:    "找不到Novel ID",
		}
	}
	c, err := models.GetCategory(service.Category)
	if err != nil {
		return &serializer.Response{
			Status: 404,
			Msg:    "找不到Category ID",
		}
	}
	target := models.NovelCategory{
		Novel:    n,
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
func (service *DeleteN2CService) Response() interface{} {
	return serializer.Response{
		Status: 0,
		Msg:    "成功",
	}
}
