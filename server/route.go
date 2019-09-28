package server

import (
	"eroauz/api"
	"eroauz/service/archive"
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
		g.GET("/archive", api.List(&ArchiveList))
		var ArchiveGet archive.GetService
		g.GET("/archive/:id", api.Get(&ArchiveGet))

		r := g.Group("/")
		{
			// 需要登陆的
			config := middleware.JWTConfig{
				Claims:     &utils.JwtCustomClaims{},
				SigningKey: []byte(utils.Secret()),
			}
			r.Use(middleware.JWTWithConfig(config))

			r.GET("/user/:id", api.UserSelf)

			a := r.Group("/")
			{
				// 需要特殊权限
				a.Use(m.AuthRequired)
				var ArchiveCreate archive.CreateService
				a.POST("/archive/", api.Create(&ArchiveCreate))
				var ArchiveDelete archive.DeleteService
				a.DELETE("/archive/:id", api.Delete(&ArchiveDelete))
				var ArchiveUpdate archive.UpdateService
				a.PUT("/archive/:id", api.Update(&ArchiveUpdate))
			}
		}

	}

	return e
}
