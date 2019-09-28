package tag

type TagService struct {
	Title uint   `json:"title" form:"title"`
	Type  string `json:"type" form:"type"`
	Rid   int    `json:"rid" form:"rid"`
}

//func (service *TagService)Get()(models.Tag,*serializer.Response) {
//
//}
