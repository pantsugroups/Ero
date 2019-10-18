package volume

import (
	"eroauz/conf"
	"eroauz/models"
	"eroauz/serializer"
)

type CreateService struct {
	ID     uint   `json:"id" param:"id" form:"id" query:"id"  null:"false"`
	Title  string `json:"title" form:"title" query:"title"`
	Cover  string `json:"cover" form:"cover" query:"cover"`
	File   int    `json:"file" form:"file"  null:"false"`
	ZIndex int    `json:"z_index" form:"z_index"`
	Type   int    `json:"type" form:"type"`
	result models.Volume
}

// EroAPI godoc
// @Summary 创建小说分卷
// @Description 必须登陆
// @Tags volume,admin
// @Accept html
// @Produce json
// @Success 200 {object} serializer.VolumeResponse
// @Failure 500 {object} serializer.Response
// @Param id path int true "小说ID"
// @Param title formData string true "分卷标题"
// @Param cover formData string false "分卷封面，URL，如果封面为空的话泽会自动替换。默认封面请检查conf.DefaultCover字段"
// @Param file formData integer false "文件ID"
// @Param z_index formData integer false "分卷顺序"
// @Param type formData integer false "类型，0为本地，1为远程"
// @Router /api/v1/volume/ [post]
// @Security ApiKeyAuth
func (service *CreateService) Create(create uint) *serializer.Response {
	var file models.File

	n, err := models.GetNovel(service.ID)
	if err != nil {
		return &serializer.Response{
			Status: 404,
			Msg:    "寻找匹配ID失败",
		}
	}
	if service.File != 0 {
		f, err := models.GetFile(service.File)
		if err != nil {
			return &serializer.Response{
				Status: 404,
				Msg:    "寻找匹配ID失败",
			}
		}
		file = f
	}
	if service.Cover == "" {
		service.Cover = conf.DefaultCover
	}
	volume := models.Volume{
		Title:  service.Title,
		Cover:  service.Cover,
		Novel:  n,
		File:   file,
		ZIndex: service.ZIndex,
		Type:   service.Type,
	}
	if err := models.DB.Create(&volume).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "创建失败",
		}
	}
	service.result = volume
	return nil
}
func (service *CreateService) Response() interface{} {
	return serializer.BuildVolumeResponse(service.result)

}
