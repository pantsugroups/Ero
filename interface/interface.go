package _interface

import (
	"eroauz/serializer"
)

type ListInterface interface {
	HaveNextOrLast() (next bool, last bool)
	Pages() (int, *serializer.Response)
	Pull()  *serializer.Response
	Counts() int
	Response() interface{}
}

type GetInterface interface {
	Get() *serializer.Response
	Response() interface{}
}

type CreateInterface interface {
	Create() *serializer.Response
	Response() interface{}
}
type UpdateInterface interface {
	Update() *serializer.Response
	Response() interface{}
}
type DeleteInterface interface {
	Delete() *serializer.Response
	Response() interface{}
}
