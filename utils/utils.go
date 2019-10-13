package utils

import (
	"crypto/md5"
	"eroauz/conf"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"math/rand"
	"net/smtp"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func UnixForString() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func Generate(hash string) string {
	ts := time.Now().Unix()
	long := int(ts) / 600
	t := strconv.Itoa(long) + hash + conf.Secret
	data := []byte(t)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	return md5str1
}
func IsNull(kind reflect.Kind, value reflect.Value) error {
	switch kind {
	case reflect.Int:
		if value.Int() == 0 {
			return errors.New("Non-empty field is none ")
		}
	case reflect.String:
		if value.String() == "" {
			return errors.New("Non-empty field is none ")
		}
	case reflect.Int64:
		if value.Int() == 0 {
			return errors.New("Non-empty field is none ")
		}
	case reflect.Uint:
		if value.Uint() == 0 {
			return errors.New("Non-empty field is none ")
		}
	case reflect.Uint8:
		if value.Uint() == 0 {
			return errors.New("Non-empty field is none ")
		}
	case reflect.Uint32:
		if value.Uint() == 0 {
			return errors.New("Non-empty field is none ")
		}
	}
	return nil
}

// 用户判断字段是否为空
func Bind(i interface{}, c echo.Context) (err error) {

	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != nil {
		return err
	}
	t := reflect.TypeOf(i).Elem()
	val := reflect.ValueOf(i).Elem()
	for i := 0; i < t.NumField(); i++ {
		free := t.Field(i).Tag.Get("null")
		param := t.Field(i).Tag.Get("param")

		value := c.Param(param)

		valueObj := val.Field(i)
		if param != "" || value != "" {
			switch valueObj.Kind() {
			case reflect.Int:
				v, _ := strconv.ParseInt(value, 0, 64)
				valueObj.SetInt(v)
			case reflect.Int8:
				v, _ := strconv.ParseInt(value, 0, 8)
				valueObj.SetInt(v)
			case reflect.Int32:
				v, _ := strconv.ParseInt(value, 0, 32)
				valueObj.SetInt(v)
			case reflect.Uint:
				v, _ := strconv.ParseUint(value, 0, 64)
				valueObj.SetUint(v)
			case reflect.Uint32:
				v, _ := strconv.ParseUint(value, 0, 32)
				valueObj.SetUint(v)
			case reflect.Uint8:
				v, _ := strconv.ParseUint(value, 0, 8)
				valueObj.SetUint(v)
			case reflect.String:
				valueObj.SetString(value)

			}
		}

		if free == "false" {
			if err := IsNull(valueObj.Kind(), valueObj); err != nil {
				return err
			}
		}
	}

	return
}

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + "\r\n" + contentType + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}
