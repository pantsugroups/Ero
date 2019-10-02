package relationship

import (
	"eroauz/models"
	"eroauz/serializer"
)

type AppendN2CService struct {
	Novel    uint `json:"novel" form:"novel"`
	Category uint `json:"category" form:"category"`
	result   models.NovelCategory
}

func (service *AppendN2CService) Create(creater uint) *serializer.Response {
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
	relationship := models.NovelCategory{
		Novel:    n,
		Category: c,
	}
	if err := models.DB.Create(&relationship).Error; err != nil {
		return &serializer.Response{
			Status: 40007,
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
