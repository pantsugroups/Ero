package serializer

import (
	"eroauz/models"
)

type Novel struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Cover      string `json:"cover"`
	Author     string `json:"author"`
	Subscribed int    `json:"subscribed"`
	Ended      bool   `json:"ended"`
	Level      int    `json:"level"`
}

type NovelResponse struct {
	Response
	Data Novel `json:"data"`
}

type NovelListResponse struct {
	Response
	Count int     `json:"count"`
	Data  []Novel `json:"data"`
	Next  bool    `json:"have_next"`
	Last  bool    `json:"have_last"`
	Pages int     `json:"pages"`
}

func BuildNovel(novel models.Novel) Novel {
	return Novel{
		ID:         novel.ID,
		Title:      novel.Title,
		Cover:      novel.Cover,
		Author:     novel.Author,
		Subscribed: novel.Subscribed,
		Ended:      novel.Ended,
		Level:      novel.Level,
	}
}

func BuildNovelList(novels []models.Novel) []Novel {
	var novelList []Novel
	for _, a := range novels {
		i := BuildNovel(a)
		novelList = append(novelList, i)
	}
	return novelList
}

func BuildNovelResponse(novel models.Novel) NovelResponse {
	return NovelResponse{
		Data: BuildNovel(novel),
	}
}

func BuildNovelListResponse(novels []models.Novel, count int, next bool, last bool, pages int) NovelListResponse {
	return NovelListResponse{
		Count: count,
		Data:  BuildNovelList(novels),
		Next:  next,
		Last:  last,
		Pages: pages,
	}
}
