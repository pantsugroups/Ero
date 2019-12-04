package archive

import (
	"eroauz/models"
	"eroauz/serializer"
	"eroauz/utils"
)

type CreateService struct {
	Title          string `json:"title" form:"title" null:"false"`
	JapTitle       string `json:"japanese_title" form:"japanese_title"`
	Author         string `json:"author" form:"author"`
	Content        string `json:"content" form:"content"`
	PrimaryContent string `json:"primary_content" form:"primary_content"`
	Cover          string `json:"cover" form:"cover"`
	Tag            string `json:"tag" form:"tag"`
	VerifyCode     string `json:"verify_code" form:"verify_code" null:"false"`
	VerifyCodeId   string `json:"verify_id" form:"verify_id" null:"false"`
	result         models.Archive
}

// EroAPI godoc
// @Summary 创建文章
// @Description 必须要登陆
// @Tags archive
// @Accept html
// @Produce json
// @Param title formData string true "标题"
// @Param japanese_title formData string false "日文标题"
// @Param author formData string false "作者"
// @Param content formData string false "文章内容"
// @Param primary_content formData string false "隐藏内容"
// @Param cover formData string false "封面，这个是个url地址"
// @Param tag formData string false "标签，推荐使用/分割，例如 纯爱/治愈/等等 "
// @Success 200 {object} serializer.ArchiveResponse
// @Failure 500 {object} serializer.Response
// @Router /api/v1/archive/ [post]
// @Security ApiKeyAuth
func (service *CreateService) Create(create uint) *serializer.Response {
	user, _ := models.GetUser(create)
	if user.Status != models.Admin {
		if res := utils.VerifyCaptcha(service.VerifyCodeId, service.VerifyCode); res == false {
			return &serializer.Response{
				Status: 403,
				Msg:    "验证码错误",
			}
		}
	}
	archive := models.Archive{
		Title:          service.Title,
		JapTitle:       service.JapTitle,
		Content:        service.Content,
		Author:         service.Author,
		PrimaryContent: service.PrimaryContent,
		Cover:          service.Cover,
		Create:         user,
		Tag:            service.Tag,
		Pass:           false,
	}

	if user.Status == models.Admin {
		archive.Pass = true
	}
	if err := models.DB.Create(&archive).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "创建失败",
		}
	}
	service.result = archive

	return nil
}
func (service *CreateService) Response() interface{} {
	return serializer.BuildArchiveResponse(service.result)

}
