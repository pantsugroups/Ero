package server

import (
	"encoding/json"
	"eroauz/api"
	"eroauz/conf"
	_ "eroauz/docs"
	"eroauz/service/archive"
	"eroauz/service/category"
	"eroauz/service/comment"
	"eroauz/service/message"
	"eroauz/service/novel"
	"eroauz/service/relationship"
	"eroauz/service/search"
	"eroauz/service/user"
	"eroauz/service/volume"
	"eroauz/utils"
	"github.com/labstack/echo"
	"io/ioutil"
	"path"
)
import "github.com/labstack/echo/middleware"
import m "eroauz/middleware"
import "github.com/swaggo/echo-swagger"

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowCredentials, echo.HeaderAccessControlAllowHeaders},
		AllowCredentials: true,

		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}))
	e.Use(middleware.Logger())

	e.Use(middleware.Recover())
	e.GET("/swagger/*", echoSwagger.WrapHandler).Name = "API"
	e.File("/routes.json", "routes.json").Name = "路由表"
	e.Static("/img", path.Join(conf.StaticPath, "img")).Name = "静态图片文件"
	e.Static("/other", path.Join(conf.StaticPath, "other")).Name = "静态文件"
	g := e.Group("/api/v1")
	{
		//普通等级路由
		g.GET("/verify", api.Captcha).Name = "验证码获取"

		g.POST("/user/login", api.UserLogin).Name = "用户登陆"

		g.POST("/user/register", api.UserRegister).Name = "用户注册"

		g.GET("/user/register", api.VerifyMail).Name = "用户邮箱验证"

		var SearchNovel search.NovelListService
		g.POST("/search/novel/", api.List(&SearchNovel)).Name = "小说搜索"

		var SearchArchive search.ArchiveListService
		g.POST("/search/archive/", api.List(&SearchArchive)).Name = "文章搜索"

		var ArchiveList archive.ListService
		g.GET("/archive/", api.List(&ArchiveList)).Name = "文章列表"

		var NovelList novel.ListService
		g.GET("/novel/", api.List(&NovelList)).Name = "小说列表"

		var CommentList comment.ListService
		g.GET("/comments/:id", api.List(&CommentList)).Name = "评论列表"

		var CategoryList category.ListService
		g.GET("/category/", api.List(&CategoryList)).Name = "分类列表"

		var ArchiveGet archive.GetService
		g.GET("/archive/:id", api.Get(&ArchiveGet)).Name = "文章查看"

		var NovelGet novel.GetService
		g.GET("/novel/:id", api.Get(&NovelGet)).Name = "小说查看"

		var CommentGet comment.GetService
		g.GET("/comment/:id", api.Get(&CommentGet)).Name = "评论查看"

		var VolumeList volume.ListService
		g.GET("/novel/volume/:id", api.List(&VolumeList)).Name = "查看小说分卷"

		r := g.Group("")
		{
			// 需要登陆的
			config := middleware.JWTConfig{
				Claims:     &utils.JwtCustomClaims{},
				SigningKey: []byte(conf.Secret),
			}
			r.Use(middleware.JWTWithConfig(config))
			r.Use(m.BaseRequired)

			r.POST("/upload/", api.Upload).Name = "上传文件"
			r.GET("/download", api.Download).Name = "下载小说"
			r.GET("/user/sendmail", api.SendMail).Name = "发送验证邮件"

			var UserBook user.SubscribeListService
			r.GET("/user/book", api.List(&UserBook)).Name = "用户书架列表"

			var UserComments user.CommentListService
			r.GET("/user/comments", api.List(&UserComments)).Name = "用户发送的评论列表"

			var UserSelf user.GetService
			r.GET("/user/", api.Get(&UserSelf)).Name = "查看用户自己信息"

			var UserGet user.GetService
			r.GET("/user/:id", api.Get(&UserGet)).Name = "查看用户信息"

			var ArchiveCreate archive.CreateService
			r.POST("/archive/", api.Create(&ArchiveCreate)).Name = "创建文章"

			var NovelCreate novel.CreateService
			r.POST("/novel/", api.Create(&NovelCreate)).Name = "创建小说"

			var CommentCreate comment.CreateService
			r.POST("/comment/", api.Create(&CommentCreate)).Name = "创建评论"

			var Novel2Category relationship.AppendN2CService
			r.POST("/category/novel/", api.Create(&Novel2Category)).Name = "关联小说分类"

			var Archive2Category relationship.AppendA2CService
			r.POST("/category/archive/", api.Create(&Archive2Category)).Name = "关联文章分类"

			var NovelSubscribe novel.SubscribeService
			r.GET("/novel/subscribe/:id", api.Create(&NovelSubscribe)).Name = "订阅小说"

			var NovelDeSubscribe novel.SubscribeService
			r.DELETE("/novel/subscribe/:id", api.Delete(&NovelDeSubscribe)).Name = "取消订阅小说"

			a := r.Group("")
			{
				// 需要特殊权限(自己为创建者或管理员)
				// todo:鉴权
				var MessageList message.ListService
				a.GET("/message/", api.List(&MessageList)).Name = "查看消息列表"

				var MessageUpdate message.UpdateService
				a.GET("/message/:id", api.Update(&MessageUpdate)).Name = "查看单个消息"

				var VolumeDown volume.GetService
				a.GET("/volume/:id", api.Get(&VolumeDown)).Name = "分卷下载"

				var CommentDelete comment.DeleteService
				a.DELETE("/comment/:id", api.Delete(&CommentDelete)).Name = "删除评论"

				var MessageDelete message.DeleteService
				a.DELETE("/message/:id", api.Delete(&MessageDelete)).Name = "删除消息"

				var ArchiveUpdate archive.UpdateService
				a.PUT("/archive/:id", api.Update(&ArchiveUpdate)).Name = "更新文章"

				var NovelUpdate novel.UpdateService
				a.PUT("/novel/:id", api.Update(&NovelUpdate)).Name = "更新小说"

				var UserUpdate user.UpdateService
				a.PUT("/user/", api.Update(&UserUpdate)).Name = "更新用户信息"

				s := a.Group("")
				{
					// 超级权限
					s.Use(m.AuthRequired)

					var VolumeCreate volume.CreateService
					s.POST("/volume/:id", api.Create(&VolumeCreate)).Name = "创建小说分卷"

					var CategoryCreate category.CreateService
					r.POST("/category/", api.Create(&CategoryCreate)).Name = "创建分类"

					var VolumeUpdate volume.UpdateService
					s.PUT("/volume/:id", api.Update(&VolumeUpdate)).Name = "更新小说分卷信息"

					var CategoryUpdate category.UpdateService
					s.PUT("/category/:id", api.Update(&CategoryUpdate)).Name = "更新分类"

					var ArchiveDelete archive.DeleteService
					s.DELETE("/archive/:id", api.Delete(&ArchiveDelete)).Name = "删除文章"

					var UserDelete user.DeleteService
					s.DELETE("/user/:id", api.Delete(&UserDelete)).Name = "删除用户"

					var SuperUserUpdate user.SuperUpdateService
					s.PUT("/admin/user/:id", api.Update(&SuperUserUpdate))

					var CategoryDelete category.DeleteService
					a.DELETE("/category/:id", api.Delete(&CategoryDelete)).Name = "删除分类"

					var Novel2CategoryDelete relationship.DeleteN2CService
					a.DELETE("/category/novel/", api.Delete(&Novel2CategoryDelete)).Name = "取消小说分类关联"

					var Archive2CategoryDelete relationship.DeleteA2CService
					a.DELETE("/category/archive/", api.Delete(&Archive2CategoryDelete)).Name = "取消文章分类关联"

					var NovelDelete novel.DeleteService
					a.DELETE("/novel/:id", api.Delete(&NovelDelete)).Name = "删除小说"

					var VolumeDelete volume.DeleteService
					a.DELETE("/volume/:id", api.Delete(&VolumeDelete)).Name = "删除小说分卷"

				}
			}
		}

	}
	data, _ := json.MarshalIndent(e.Routes(), "", "  ")

	_ = ioutil.WriteFile("routes.json", data, 0644)
	return e
}
