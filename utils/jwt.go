package utils

import (
	"eroauz/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"time"
)

type JwtCustomClaims struct {
	Name string `json:"username"`
	ID   uint   `json:"id"`
	jwt.StandardClaims
}

var secret = "secret"

// 创建token
func Secret() string {
	return secret
}
func CreateToken(user models.User) (string, error) {
	claims := &JwtCustomClaims{
		user.UserName,
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil

}

func GetAutherID(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	id := claims.ID
	return id
}
