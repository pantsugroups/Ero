package utils

import (
	"eroauz/conf"
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

func CreateToken(user models.User) (string, error) {
	claims := &JwtCustomClaims{
		user.UserName,
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(conf.Secret))
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
