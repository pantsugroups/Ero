package server

import (
	"eroauz/api"
	"eroauz/service/archive"
	"eroauz/service/comment"
	"eroauz/service/novel"
	"eroauz/utils"
	"github.com/labstack/echo"
)
import "github.com/labstack/echo/middleware"
import m "eroauz/middleware"

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.Use(middleware.Recover())

	g := e.Group("/api/v1")
	{
		//普通等级路由
		g.POST("/user/login", api.UserLogin)
		g.POST("/user/register", api.UserRegister)
		var ArchiveList archive.ListService
		g.GET("/archive/", api.List(&ArchiveList))
		var NovelList novel.ListService
		g.GET("/novel/", api.List(&NovelList))
		var CommentList comment.ListService
		g.GET("/comments/:type/:id", api.List(&CommentList))
		var ArchiveGet archive.GetService
		g.GET("/archive/:id", api.Get(&ArchiveGet))
		var NovelGet novel.GetService
		g.GET("/novel/:id", api.Get(&NovelGet))
		var CommentGet comment.GetService
		g.GET("/comment/:id", api.Get(&CommentGet))
		r := g.Group("")
		{
			// 需要登陆的
			config := middleware.JWTConfig{
				Claims:     &utils.JwtCustomClaims{},
				SigningKey: []byte(utils.Secret()),
			}
			r.Use(middleware.JWTWithConfig(config))

			r.GET("/user/:id", api.UserSelf)
			var ArchiveCreate archive.CreateService
			r.POST("/archive/", api.Create(&ArchiveCreate))
			var NovelCreate novel.CreateService
			r.POST("/novel/", api.Create(&NovelCreate))
			var CommentCreate comment.CreateService
			r.POST("/comment/", api.Create(&CommentCreate))
			a := r.Group("")
			{
				// 需要特殊权限
				a.Use(m.AuthRequired)

				var ArchiveDelete archive.DeleteService
				a.DELETE("/archive/:id", api.Delete(&ArchiveDelete))
				var NovelDelete novel.DeleteService
				a.DELETE("/novel/:id", api.Delete(&NovelDelete))
				var CommentDelete comment.DeleteService
				a.DELETE("/comment/:id", api.Delete(&CommentDelete))
				var ArchiveUpdate archive.UpdateService
				a.PUT("/archive/:id", api.Update(&ArchiveUpdate))
				var NovelUpdate archive.UpdateService
				a.PUT("/novel/:id", api.Update(&NovelUpdate))
			}
		}

	}

	return e
}
