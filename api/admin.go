package api

import (
	"eroauz/conf"
	"eroauz/models"
	"eroauz/serializer"
	"eroauz/utils"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"time"
)

// 切换注册模式
func DerailRegister(c echo.Context) error {
	var result string
	if conf.AllowRegister == true {
		conf.AllowRegister = false
		result = "false"
	} else {
		conf.AllowRegister = true
		result = "true"
	}
	return c.HTML(200, result)
}

// 创建邀请码
func CreateInviteCode(c echo.Context) error {
	var invite models.Invite

	uid := utils.GetAuthorID(c)
	code := utils.RandStringRunes(6)
	dd, _ := time.ParseDuration("24h")
	limit := time.Now().Add(dd)

	if err := models.DB.Where(&models.Invite{Create: uid}).First(&invite).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return c.JSON(200, serializer.Response{
				Status: 500,
				Msg:    "读取失败",
				Error:  err.Error(),
			})
		}
	}
	if invite.ID != 0 {
		if err := models.DB.Delete(&invite).Error; err != nil {
			return c.JSON(200, serializer.Response{
				Status: 500,
				Msg:    "删除旧凭据失败",
				Error:  err.Error(),
			})
		}
	}
	//if time.Now().Before(invite.TimeLimit){
	//
	//}
	invite = models.Invite{Code: code, TimeLimit: limit, Create: uid}
	if err := models.DB.Create(&invite).Error; err != nil {
		return c.JSON(200, serializer.Response{
			Status: 500,
			Msg:    "创建失败",
			Error:  err.Error(),
		})
	}

	return c.JSON(200, serializer.Response{
		Status: 0,
		Data:   invite,
		Msg:    "",
		Error:  "",
	})
}
