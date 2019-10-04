package api

import (
	"eroauz/interface"
	"eroauz/serializer"
	"eroauz/utils"
	"fmt"
	"github.com/labstack/echo"
)

// 获取列表

func List(service _interface.ListInterface) echo.HandlerFunc {

	return func(c echo.Context) (err error) {

		if err := utils.Bind(service, c); err == nil {
			uid := utils.GetAutherID(c)
			if err := service.Pull(uid); err != nil {
				return c.JSON(200, err)
			} else {

				res := service.Response()
				return c.JSON(200, res)
			}
		} else {
			return c.JSON(200, &serializer.Response{
				Status: 40001,
				Msg:    "参数错误",
				Error:  fmt.Sprint(err),
			})
		}
	}
}

// 创建
func Create(service _interface.CreateInterface) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		if err := utils.Bind(service, c); err == nil {
			uid := utils.GetAutherID(c)
			if err := service.Create(uid); err != nil {
				return c.JSON(200, err)
			} else {
				res := service.Response()
				return c.JSON(200, res)
			}
		} else {
			return c.JSON(200, &serializer.Response{
				Status: 40001,
				Msg:    "参数错误",
				Error:  fmt.Sprint(err),
			})
		}

	}
}

// 获取单个
func Get(service _interface.GetInterface) echo.HandlerFunc {

	return func(c echo.Context) (err error) {
		uid := utils.GetAutherID(c)
		if err := utils.Bind(service, c); err == nil {
			if err := service.Get(uid); err != nil {
				return c.JSON(200, err)
			} else {
				res := service.Response()
				return c.JSON(200, res)
			}
		} else {
			return c.JSON(200, &serializer.Response{
				Status: 40001,
				Msg:    "参数错误",
				Error:  fmt.Sprint(err),
			})
		}
	}
}

//更新
func Update(service _interface.UpdateInterface) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		uid := utils.GetAutherID(c)
		if err := utils.Bind(service, c); err == nil {
			if err := service.Update(uid); err != nil {
				return c.JSON(200, err)
			} else {
				res := service.Response()
				return c.JSON(200, res)
			}
		} else {
			return c.JSON(200, &serializer.Response{
				Status: 40001,
				Msg:    "参数错误",
				Error:  fmt.Sprint(err),
			})
		}
	}
}

//删除
func Delete(service _interface.DeleteInterface) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		if err := utils.Bind(service, c); err == nil {
			uid := utils.GetAutherID(c)
			if err := service.Delete(uid); err != nil {
				return c.JSON(200, err)
			} else {
				res := service.Response()
				return c.JSON(200, res)
			}
		} else {
			return c.JSON(200, &serializer.Response{
				Status: 40001,
				Msg:    "参数错误",
				Error:  fmt.Sprint(err),
			})
		}
	}
}
