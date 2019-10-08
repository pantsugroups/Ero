package api

import (
	"eroauz/serializer"
	"eroauz/utils"
	"github.com/labstack/echo"
)

type Verify struct {
	CodeId string `json:"code_id"`
	Data   string `json:"data"`
}

// EroAPI godoc
// @Summary 验证码获取
// @Description 必须要先从/api/v1/verify 处获取验证码
// @Tags user
// @Accept html
// @Produce json
// @Success 200 {object} api.Verify
// @Router /api/v1/user/verify [get]
func Captcha(c echo.Context) error {
	id, data := utils.CodeCaptchaCreate()
	return c.JSON(200, &serializer.Response{
		Data: Verify{
			CodeId: id,
			Data:   data,
		},
	})
}
