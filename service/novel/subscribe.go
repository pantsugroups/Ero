package novel

import (
	"eroauz/models"
	"eroauz/serializer"
)

type SubscribeService struct {
	ID     uint `json:"id" form:"id" param:"id" null:"false"`
	result models.Novel
}

// EroAPI godoc
// @Summary 订阅小说
// @Description 必须登陆
// @Tags novel
// @Accept html
// @Produce json
// @Success 200 {object} serializer.Response
// @Failure 500 {object} serializer.Response
// @Param id path integer true "小说id"
// @Router /api/v1/novel/subscribe/:id [get]
// @Security ApiKeyAuth
func (service *SubscribeService) Create(create uint) *serializer.Response {
	user, err := models.GetUser(create)
	if err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "找不到用户",
			Error:  err.Error(),
		}
	}
	novel, err := models.GetNovel(service.ID)
	s := models.NovelSubscribe{
		User:  user,
		Novel: novel,
	}
	if err := models.DB.Create(&s).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "创建失败",
			Error:  err.Error(),
		}
	}
	return nil
}
func (service *SubscribeService) Response() interface{} {
	return &serializer.Response{
		Status: 200,
		Msg:    "成功",
	}
}

// EroAPI godoc
// @Summary 取消订阅小说
// @Description 必须登陆
// @Tags novel
// @Accept html
// @Produce json
// @Success 200 {object} serializer.Response
// @Failure 500 {object} serializer.Response
// @Param id path integer true "小说id"
// @Router /api/v1/novel/subscribe/:id [delete]
// @Security ApiKeyAuth
func (service *SubscribeService) Delete(create uint) *serializer.Response {
	var s models.NovelSubscribe
	if err := models.DB.Where("user_id = ?", create).Where("novel_id = ?", service.ID).First(&s).Error; err != nil {
		return &serializer.Response{
			Status: 404,
			Msg:    "寻找匹配ID失败",
			Error:  err.Error(),
		}
	}
	if err := models.DB.Delete(&s).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "删除失败",
		}
	}
	return nil
}
