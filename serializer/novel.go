package serializer

import (
	"eroauz/models"
)

type Novel struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Subscribed  int    `json:"subscribed"`
	Ended       bool   `json:"ended"`
	Level       int    `json:"level"`
	Tags        string `json:"tags"`
	UpdateAt    int64  `json:"update_at"`
}

type NovelResponse struct {
	Response
	Data Novel `json:"data"`
}

type NovelListResponse struct {
	Response
	Count int     `json:"count"`
	All   int     `json:"all"`
	Data  []Novel `json:"data"`
	Next  bool    `json:"have_next"`
	Last  bool    `json:"have_last"`
	Pages int     `json:"pages"`
}

func BuildNovel(novel models.Novel) Novel {
	return Novel{
		ID:          novel.ID,
		Title:       novel.Title,
		Cover:       novel.Cover,
		Author:      novel.Author,
		Description: novel.Description,
		Subscribed:  novel.Subscribed,
		Ended:       novel.Ended,
		Level:       novel.Level,
		Tags:        novel.Tags,
		UpdateAt:    novel.UpdatedAt.Unix(),
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

func BuildNovelListResponse(novels []models.Novel, all int, count int, next bool, last bool, pages int) NovelListResponse {
	return NovelListResponse{
		Count: count,
		All:   all,
		Data:  BuildNovelList(novels),
		Next:  next,
		Last:  last,
		Pages: pages,
	}
}
