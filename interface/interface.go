package _interface

import (
	"eroauz/serializer"
)

type ListInterface interface {
	HaveNextOrLast() (next bool, last bool)
	Pages() (int, *serializer.Response)
	Pull(create uint) *serializer.Response
	Counts() int
	Response() interface{}
}

type GetInterface interface {
	Get(create uint) *serializer.Response
	Response() interface{}
}

type CreateInterface interface {
	Create(create uint) *serializer.Response
	Response() interface{}
}
type UpdateInterface interface {
	Update(create uint) *serializer.Response
	Response() interface{}
}
type DeleteInterface interface {
	Delete(create uint) *serializer.Response
	Response() interface{}
}
