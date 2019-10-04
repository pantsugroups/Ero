package middleware

import (
	"eroauz/models"
	"eroauz/serializer"
	"eroauz/utils"
	"fmt"
	"github.com/labstack/echo"
)

// 检测特殊权限

func BaseRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Path())
		uid := utils.GetAutherID(c)
		u, err := models.GetUser(uid)
		if err != nil {

			return c.JSON(200, serializer.Response{
				Status: 500,
				Msg:    "找不到用户",
				Error:  err.Error(),
			})
		}
		if u.Status == models.Inactive || u.Status == models.Suspend {
			return c.JSON(200, serializer.Response{
				Status: 403,
				Msg:    "没有权限",
			})
		}
		return next(c)
	}
}
func AuthRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Path())
		uid := utils.GetAutherID(c)
		u, err := models.GetUser(uid)
		if err != nil {

			return c.JSON(200, serializer.Response{
				Status: 500,
				Msg:    "找不到用户",
				Error:  err.Error(),
			})
		}
		if u.Status != models.Admin {
			return c.JSON(200, serializer.Response{
				Status: 403,
				Msg:    "没有权限",
			})
		}
		return next(c)
	}
}
