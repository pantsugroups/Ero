package user

import "eroauz/serializer"
import model "eroauz/models"

// UserLoginService 管理用户登录的服务
type LoginService struct {
	UserName string `form:"username" json:"username" binding:"required,min=5,max=30"  null:"false"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"  null:"false"`
}

// EroAPI godoc
// @Summary 用户登录
// @Description
// @Tags user
// @Accept html
// @Produce json
// @Success 200 {object} serializer.UserResponse
// @Failure 500 {object} serializer.Response
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Router /api/v1/user/login [get]
func (service *LoginService) Login() (model.User, *serializer.Response) {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return user, &serializer.Response{
			Status: 40001,
			Msg:    "账号或密码错误",
		}
	}

	if user.CheckPassword(service.Password) == false {
		return user, &serializer.Response{
			Status: 40001,
			Msg:    "账号或密码错误",
		}
	}
	return user, nil
}
