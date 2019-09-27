package utils

import (
	"errors"
	"github.com/labstack/echo"
	"reflect"
)

func IsNull(kind reflect.Kind,value reflect.Value) error{
	switch kind{
	case reflect.Int:
		if value.Int() == 0{
			return errors.New("Non-empty field is none ")
		}
	case reflect.String:
		if value.String() == ""{
			return errors.New("Non-empty field is none ")
		}
	case reflect.Int64:
		if value.Int() == 0{
			return errors.New("Non-empty field is none ")
		}
	case reflect.Uint:
		if value.Int() == 0{
			return errors.New("Non-empty field is none ")
		}
	case reflect.Uint8:
		if value.Int() == 0{
			return errors.New("Non-empty field is none ")
		}
	case reflect.Uint32:
		if value.Int() == 0{
			return errors.New("Non-empty field is none ")
		}
	}
	return nil
}

// 用户判断字段是否为空
func Bind(i interface{}, c echo.Context) (err error) {

	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
		return
	}
	t := reflect.TypeOf(i).Elem()
	val := reflect.ValueOf(i).Elem()
	for i := 0; i < t.NumField(); i++ {
		free := t.Field(i).Tag.Get("null")
		valueObj := val.Field(i)

		if free == "false"{
			if err := IsNull(valueObj.Kind(),valueObj);err != nil{
				return err
			}
		}
	}

	return
}
