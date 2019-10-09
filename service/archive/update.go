package archive

import (
	"eroauz/models"
	"eroauz/serializer"
)

type UpdateService struct {
	ID             uint   `json:"id" form:"id" param:"id" null:"false"`
	Title          string `json:"title" form:"title"`
	JapTitle       string `json:"japanese_title" form:"japanese_title"`
	Cover          string `json:"cover" form:"cover"`
	Content        string `json:"content" form:"content"`
	Author         string `json:"author" form:"author"`
	Mail           string `json:"mail" form:"mail"`
	PrimaryContent string `json:"primary_content" form:"primary_content"`
	result         models.Archive
}

// EroAPI godoc
// @Summary 更新文章
// @Description 必须要登陆且创建者为自己或者自己权限是管理员
// @Tags archive
// @Accept html
// @Produce json
// @Success 200 {object} serializer.ArchiveResponse
// @Failure 500 {object} serializer.Response
// @Param ID path integer true "文章ID"
// @Param title formData string true "标题"
// @Param japanese_title formData string false "日文标题"
// @Param author formData string false "作者"
// @Param content formData string false "文章内容"
// @Param primary_content formData string false "隐藏内容"
// @Param cover formData string false "封面，这个是个url地址"
// @Param tag formData string false "标签，推荐使用/分割，例如 纯爱/治愈/等等 "
// @Router /api/v1/archive/:id [put]
func (service *UpdateService) Update(create uint) *serializer.Response {
	var archive models.Archive
	if err := models.DB.Where("ID = ?", service.ID).First(&archive).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "获取失败",
		}
	}
	user, err := models.GetUser(create)
	if err != nil {
		return &serializer.Response{
			Status: 404,
			Msg:    "找不到用户",
			Error:  err.Error(),
		}
	} else {
		if user.ID != archive.CreateID {
			return &serializer.Response{
				Status: 403,
				Msg:    "没有权限",
			}
		}
	}
	if err := models.DB.Model(&archive).Update(models.Archive{
		Title:          service.Title,
		JapTitle:       service.JapTitle,
		Cover:          service.Cover,
		Content:        service.Content,
		Author:         service.Author,
		PrimaryContent: service.PrimaryContent,
	}); err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	service.result = archive
	return nil
}
func (service *UpdateService) Response() interface{} {
	return service.result

}
