package relationship

import (
	"eroauz/models"
	"eroauz/serializer"
)

type AppendA2CService struct {
	Archive  uint `json:"archive" form:"archive"  null:"false"`
	Category uint `json:"category" form:"category"  null:"false"`
	result   models.ArchiveCategory
}

// EroAPI godoc
// @Summary 添加文章分类关联
// @Description 必须为管理员
// @Tags archive,category,admin
// @Accept html
// @Produce json
// @Success 200 {object} serializer.Response
// @Failure 500 {object} serializer.Response
// @Param archive formData integer true "文章ID"
// @Param category formData integer true "分类ID"
// @Router /api/v1/category/archive/ [post]
// @Security ApiKeyAuth
func (service *AppendA2CService) Create(create uint) *serializer.Response {
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
			Msg:    "找不到Novel ID",
		}
	}
	if c.Type != models.Archive_ {
		return &serializer.Response{
			Status: 500,
			Msg:    "不能将文章添加到其他分区",
		}
	}
	relationship := models.ArchiveCategory{
		Archive:  a,
		Category: c,
	}
	if err := models.DB.Create(&relationship).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "创建失败",
		}
	}

	models.DB.Model(&c).Update("count", c.Count+1)
	service.result = relationship
	return nil
}
func (service *AppendA2CService) Response() interface{} {
	return serializer.Response{
		Status: 0,
		Msg:    "成功",
	}
}
