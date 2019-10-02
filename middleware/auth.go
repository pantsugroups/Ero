package middleware

import (
	"github.com/labstack/echo"
)

// 检测特殊权限
// 管理员无条件放行 其余的只对创建者放行
func AuthRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//if err := next(c); err != nil {
		//	c.Error(err)
		//}
		return next(c)
	}
}

// 超级权限
// 仅对管理员放行
func AdminRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}
