package _interface

import (
	"eroauz/models"
	"eroauz/serializer"
)

type ListInterface interface {
	HaveNextOrLast() (next bool, last bool)
	Pages() (int, *serializer.Response)
	Pull() ([]models.Archive, *serializer.Response)
	Counts() int
}

type GetInterface interface {
	Get() *serializer.Response
	Response() *serializer.Response
}

type CreateInterface interface {
	Create() *serializer.Response
	Response() *serializer.Response
}
type UpdateInterface interface {
	Update() *serializer.Response
	Response() *serializer.Response
}
type DeleteInterface interface {
	Delete() *serializer.Response
	Response() *serializer.Response
}
