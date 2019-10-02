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
	RUid   uint   `json:"reply" form:"reply"`
	result models.Comment
}

func (service *CreateService) Create(creater uint) *serializer.Response {
	u, _ := models.GetUser(creater)
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
		RUid:   service.RUid,
	}
	if err := models.DB.Create(&archive).Error; err != nil {
		return &serializer.Response{
			Status: 40004,
			Msg:    "创建失败",
		}
	}
	if service.RUid != 0 {
		M := message.CreateService{
			Title: "您的消息有回复啦！<a>查看回复</a>",
			Recv:  u.ID,
		}
		if err := M.Create(creater); err != nil {
			return err
		}
	}
	service.result = archive
	return nil
}
func (service *CreateService) Response() interface{} {
	return serializer.BuildCommentResponse(service.result)

}
