package novel

import (
	"eroauz/models"
	"eroauz/serializer"
)

type CreateService struct {
	Title       string `json:"title" form:"title"`
	Author      string `json:"author" form:"title"`
	Cover       string `json:"cover" form:"cover"`
	Description string `json:"description" form:"description"`
	Ended       bool   `json:"ended" form:"ended"`
	Level       int    `json:"level" form:"level"`
	result      models.Novel
}

// EroAPI godoc
// @Summary 创建小说
// @Description 必须登陆
// @Tags novel
// @Accept html
// @Produce json
// @Success 200 {object} serializer.NovelResponse
// @Failure 500 {object} serializer.Response
// @Param title formData string true "小说标题"
// @Param author formData string true "小说作者"
// @Param cover formData string false "小说封面，URL，如果封面为空的话泽会自动替换。默认封面请检查conf.DefaultCover字段"
// @Param description formData integer false "小说简介"
// @Param ended formData boolean true "是否完结"
// @Param level formData integer true "目标等级：还未实现"
// @Router /api/v1/comment/ [post]
// @Security ApiKeyAuth
func (service *CreateService) Create(create uint) *serializer.Response {
	u, _ := models.GetUser(create)
	novel := models.Novel{
		Title:       service.Title,
		Author:      service.Author,
		Cover:       service.Cover,
		Description: service.Description,
		Ended:       service.Ended,
		Level:       service.Level,
		Subscribed:  0,
		Create:      u,
	}
	novel.CheckCover()
	if err := models.DB.Create(&novel).Error; err != nil {
		return &serializer.Response{
			Status: 40007,
			Msg:    "创建失败",
			Error:  err.Error(),
		}
	}

	service.result = novel
	return nil
}
func (service *CreateService) Response() interface{} {
	return serializer.BuildNovelResponse(service.result)
}
