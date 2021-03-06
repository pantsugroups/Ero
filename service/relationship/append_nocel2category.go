package relationship

import (
	"eroauz/models"
	"eroauz/serializer"
)

type AppendN2CService struct {
	Novel    uint `json:"novel" form:"novel"  null:"false"`
	Category uint `json:"category" form:"category"  null:"false"`
	result   models.NovelCategory
}

// EroAPI godoc
// @Summary 添加小说分类关联
// @Description 必须为管理员
// @Tags novel,category,admin
// @Accept html
// @Produce json
// @Success 200 {object} serializer.Response
// @Failure 500 {object} serializer.Response
// @Param novel formData integer true "小说ID"
// @Param category formData integer true "分类ID"
// @Router /api/v1/category/novel/ [post]
// @Security ApiKeyAuth
func (service *AppendN2CService) Create(create uint) *serializer.Response {
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
			Msg:    "找不到Novel ID",
		}
	}
	if c.Type != models.Novel_ {
		return &serializer.Response{
			Status: 500,
			Msg:    "不能将小说添加到其他分区",
		}
	}
	relationship := models.NovelCategory{
		Novel:    n,
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
func (service *AppendN2CService) Response() interface{} {
	return serializer.Response{
		Status: 0,
		Msg:    "成功",
	}
}
