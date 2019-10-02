package api

import (
	"eroauz/serializer"
	"eroauz/service/user"
	"eroauz/utils"
	"fmt"
	"github.com/labstack/echo"
)

func UserRegister(c echo.Context) (err error) {

	var service user.RegisterService
	if err := c.Bind(&service); err == nil {
		if users, err := service.Register(); err != nil {
			return c.JSON(200, err)
		} else {
			res := serializer.BuildUserResponse(users)
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
func UserLogin(c echo.Context) (err error) {
	var service user.LoginService
	if err := c.Bind(&service); err == nil {
		if users, err := service.Login(); err != nil {
			return c.JSON(200, err)
		} else {
			// 设置Session
			token, err := utils.CreateToken(users)
			if err != nil {
				return c.JSON(200, &serializer.Response{
					Status: 40003,
					Msg:    "token生成失败",
					Error:  fmt.Sprint(err)})
			}
			res := serializer.BuildTokenResponse(users, token)
			return c.JSON(200, res)
		}
	} else {
		return c.JSON(200, &serializer.Response{
			Status: 40001,
			Msg:    "参数错误",
			Error:  fmt.Sprint(err)})
	}
}
func UserSelf(c echo.Context) (err error) {

	return nil
}
