package api

import (
	"eroauz/conf"
	"eroauz/models"
	"eroauz/serializer"
	"eroauz/utils"
	"github.com/labstack/echo"
	"io"
	"os"
	"path"
)

// EroAPI godoc
// @Summary 上传文件
// @Description 必须要登陆，文件ID就是在这里去拿的，图片等静态 文件可以直接访问/img/*获得，novel必须要经过/download
// @Tags file,volume
// @Accept html
// @Produce json
// @Param type formData string true "类型，分为novel,img两个，如果是img则可以直接去静态文件请求"
// @Param file formData string true "文件"
// @Success 200 {object} serializer.FileResponse
// @Failure 500 {object} serializer.Response
// @Router /api/v1/upload/ [post]
// @Security ApiKeyAuth
func Upload(c echo.Context) error {
	var dir string
	var i int
	uid := utils.GetAuthorID(c)
	user, _ := models.GetUser(uid)
	VerifyCode := c.FormValue("verify_code")
	VerifyCodeId := c.FormValue("verify_id")
	types := c.FormValue("type")
	file, err := c.FormFile("file")
	u, _ := models.GetUser(uid)
	if u.Status != models.Admin {
		if res := utils.VerifyCaptcha(VerifyCodeId, VerifyCode); res == false {
			return c.JSON(200, &serializer.Response{
				Status: 403,
				Msg:    "验证码错误",
			})
		}
	}
	if types == "novel" {
		dir = "novel"
		i = models.Volume_
	} else if types == "img" {
		dir = "img"
		i = models.Image_
	} else {
		dir = "other"
		i = models.Other_
	}

	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func() {
		_ = src.Close()
	}()

	// Destination
	filename := utils.UnixForString() + path.Ext(file.Filename)
	s := path.Join(conf.StaticPath, dir, filename)
	dst, err := os.Create(s)
	if err != nil {
		return err
	}
	defer func() {
		_ = dst.Close()
	}()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	f := models.File{
		User:     user,
		Path:     s,
		FileName: filename,
		Type:     i,
	}
	if err := models.DB.Create(&f).Error; err != nil {
		return c.JSON(200, &serializer.Response{
			Status: 40004,
			Msg:    "创建失败",
		})
	}
	return c.JSON(200, serializer.BuildFileResponse(f))
}

// EroAPI godoc
// @Summary 下载小说分卷
// @Description 必须要登陆，文件ID就是在这里去拿的，图片等静态 文件可以直接访问/img/获得，novel必须要经过/download
// @Tags file,volume
// @Accept html
// @Produce json
// @Param token query string true "从/api/v1/novel/volume/:id得来"
// @Param filename query string true "从/api/v1/novel/volume/:id得来"
// @Param hash query string true "从/api/v1/novel/volume/:id得来"
// @Success 200 {object} serializer.FileResponse
// @Failure 500 {object} serializer.Response
// @Router /api/v1/download/ [get]
// @Security ApiKeyAuth
func Download(c echo.Context) error {

	id := c.QueryParam("id")
	token := c.QueryParam("token")
	file := c.QueryParam("filename")
	hash := c.QueryParam("hash")
	if len(token) != 16 {
		return c.JSON(200, serializer.Response{
			Status: 403,
			Msg:    "验证失败",
		})
	}
	f, err := models.GetFile(id)
	if err != nil {
		return c.JSON(200, serializer.Response{
			Status: 404,
			Msg:    "找不到文件",
			Error:  err.Error(),
		})
	}
	if hash != utils.Generate(file) {
		return c.JSON(200, serializer.Response{
			Status: 403,
			Msg:    "令牌错误",
		})
	} else {
		if f.Type != models.Remote_ {
			return c.Attachment(f.Path, f.FileName)
		} else {
			return c.HTML(302, `<script language="javascript" type="text/javascript">window.location.href=`+f.Path+`';</script>
`)
		}

	}

}
