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

// todo:偷懒直接写API层了
func Upload(c echo.Context) error {
	var dir string
	var i int
	uid := utils.GetAutherID(c)
	user, _ := models.GetUser(uid)
	types := c.FormValue("type")

	if types == "novel" {
		dir = "/novel/"
		i = models.Volume_
	} else if types == "img" {
		dir = "/img/"
		i = models.Image_
	} else {
		dir = "/other/"
		i = models.Other_
	}
	file, err := c.FormFile("file")
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
	s := path.Join(conf.StaticPath, dir, file.Filename)
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
		FileName: file.Filename,
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
func Download(c echo.Context) error {
	token := c.QueryParam("token")
	file := c.QueryParam("filename")
	hash := c.QueryParam("hash")
	if len(token) != 16 {
		return c.JSON(200, serializer.Response{
			Status: 403,
			Msg:    "验证失败",
		})
	}
	f := models.File{
		FileName: file,
	}
	if err := models.DB.First(f).Error; err != nil {
		return c.JSON(200, serializer.Response{
			Status: 404,
			Msg:    "找不到文件",
		})
	}
	if hash != utils.Generate(file) {
		return c.JSON(200, serializer.Response{
			Status: 403,
			Msg:    "令牌错误",
		})
	} else {
		return c.File(f.Path)
	}

}
