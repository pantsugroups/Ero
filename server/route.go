package server

import (
	"eroauz/api"
	"eroauz/service/archive"
	"eroauz/service/category"
	"eroauz/service/comment"
	"eroauz/service/novel"
	"eroauz/service/relationship"
	"eroauz/service/user"
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

		var CategoryList category.ListService
		g.GET("/category/", api.List(&CategoryList))

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

			var CategoryCreate category.CreateService
			r.POST("/category/", api.Create(&CategoryCreate))

			var Novel2Category relationship.AppendN2CService
			r.POST("/category/", api.Create(&Novel2Category))

			a := r.Group("")
			{
				// 需要特殊权限(自己为创建者或管理员)
				a.Use(m.AuthRequired)

				var CommentDelete comment.DeleteService
				a.DELETE("/comment/:id", api.Delete(&CommentDelete))

				var ArchiveUpdate archive.UpdateService
				a.PUT("/archive/:id", api.Update(&ArchiveUpdate))

				var NovelUpdate archive.UpdateService
				a.PUT("/novel/:id", api.Update(&NovelUpdate))

				var UserUpdate user.UpdateService
				a.PUT("/user/:id", api.Update(&UserUpdate))

				s := a.Group("")
				{
					// 超级权限
					var CategoryUpdate category.UpdateService
					a.PUT("/category/:id", api.Update(&CategoryUpdate))

					var ArchiveDelete archive.DeleteService
					a.DELETE("/archive/:id", api.Delete(&ArchiveDelete))

					var UserDelete user.DeleteService
					s.DELETE("/user/:id", api.Delete(&UserDelete))

					var CategoryDelete category.DeleteService
					a.DELETE("/category/:id", api.Delete(&CategoryDelete))

					var Novel2CategoryDelete relationship.DeleteN2CService
					a.DELETE("/category/", api.Delete(&Novel2CategoryDelete))

					var NovelDelete novel.DeleteService
					a.DELETE("/novel/:id", api.Delete(&NovelDelete))
				}
			}
		}

	}

	return e
}
