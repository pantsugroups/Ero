package relationship

import (
	"eroauz/models"
	"eroauz/serializer"
)

type DeleteN2CService struct {
	Novel    uint `json:"novel" form:"novel" null:"false"`
	Category uint `json:"category" form:"category" null:"category"`
}

func (service *DeleteN2CService) Delete() *serializer.Response {

	n, err := models.GetNovel(service.Novel)
	if err != nil {
		return &serializer.Response{
			Status: 40007,
			Msg:    "找不到Novel ID",
		}
	}
	c, err := models.GetCategory(service.Category)
	if err != nil {
		return &serializer.Response{
			Status: 40007,
			Msg:    "找不到Novel ID",
		}
	}
	target := models.NovelCategory{
		Novel:    n,
		Category: c,
	}
	if err := models.DB.Delete(&target).Error; err != nil {
		return &serializer.Response{
			Status: 40005,
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
