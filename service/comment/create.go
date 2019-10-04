package comment

import (
	"eroauz/models"
	"eroauz/serializer"
	"eroauz/service/message"
)

type CreateService struct {
	Title  string `json:"title" form:"title"`
	Type   int    `json:"type" form:"type"`
	RId    uint   `json:"raw"  form:"raw"`
	RCid   uint   `json:"reply" form:"reply"`
	result models.Comment
}

// EroAPI godoc
// @Summary 创建评论
// @Description 必须登陆
// @Tags comment
// @Accept html
// @Produce json
// @Success 200 {object} serializer.CommentResponse
// @Failure 500 {object} serializer.Response
// @Param title formData string true "评论内容"
// @Param type formData int true "评论类型。小说为：2 文章为：1 其余泽忽略"
// @Param raw formData int true "如果type参数为1泽这是文章ID，如果type参数为2则是小说ID"
// @Param reply formData int false "要返回用户的评论的id"
// @Router /api/v1/comment/ [post]
// @Security ApiKeyAuth
func (service *CreateService) Create(create uint) *serializer.Response {
	u, _ := models.GetUser(create)
	if service.Type == models.Archive_ {
		var archive models.Archive
		archive.ID = service.RId
		if err := models.DB.Where("ID = ?", archive.ID).First(&archive).Error; err != nil {
			return &serializer.Response{
				Status: 40005,
				Msg:    "寻找匹配ID失败",
				Error:  err.Error(),
			}
		}
	} else if service.Type == models.Novel_ {
		var novel models.Novel
		novel.ID = service.RId
		if err := models.DB.Where("ID = ?", service.RId).First(&novel).Error; err != nil {
			return &serializer.Response{
				Status: 40005,
				Msg:    "寻找匹配ID失败",
			}
		}
	}
	archive := models.Comment{
		Title:  service.Title,
		Author: u,
		Type:   service.Type,
		RId:    service.RId,
		RCid:   service.RCid,
	}
	if err := models.DB.Create(&archive).Error; err != nil {
		return &serializer.Response{
			Status: 40004,
			Msg:    "创建失败",
		}
	}
	if service.RCid != 0 {
		c, _ := models.GetComment(service.RCid)
		if c.AuthorID != 0 {
			M := message.CreateService{
				Title: "您的消息有回复啦！<a>查看回复</a>",
				Recv:  c.AuthorID,
			}
			if err := M.Create(create); err != nil {
				return err
			}
		}

	}
	service.result = archive
	return nil
}
func (service *CreateService) Response() interface{} {
	return serializer.BuildCommentResponse(service.result)

}
