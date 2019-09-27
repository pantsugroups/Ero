package server
import (
	"eroauz/api"
	"eroauz/utils"

	"github.com/labstack/echo"
)
import "github.com/labstack/echo/middleware"
import m "eroauz/middleware"


func NewRouter()*echo.Echo{
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.Use(middleware.Recover())

	g:= e.Group("/api/v1")
	{
		//普通等级路由
		g.POST("/user/login",api.UserLogin)
		g.POST("/user/register",api.UserRegister)
		g.GET("/archive",api.Archive)
		g.GET("/archive/:id",api.ArchiveGet)

		r:= g.Group("/")
		{
			// 需要登陆的
			config := middleware.JWTConfig{
				Claims:     &utils.JwtCustomClaims{},
				SigningKey: []byte(utils.Secret()),
			}
			r.Use(middleware.JWTWithConfig(config))


			r.GET("/user/:id",api.UserSelf)




			a:=r.Group("/")
			{
				// 需要特殊权限
				a.Use(m.AuthRequired)
				a.POST("/archive/",api.ArchiveNew)
				a.DELETE("/archive/:id",api.ArchiveDelete)
				a.PUT("/archive/:id",api.ArchiveUpdate)
			}
		}

	}


	return e
}