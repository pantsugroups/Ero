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

// ArchiveResponse 单个用户序列化
type NovelResponse struct {
	Response
	Data Novel `json:"data"`
}

// 相应列表
type NovelListResponse struct {
	Response
	Count int     `json:"count"`
	Data  []Novel `json:"data"`
	Next  bool    `json:"have_next"`
	Last  bool    `json:"have_last"`
	Pages int     `json:"pages"`
}

// BuildArchive 单个序列化文章
func BuildNovel(novel models.Novel) Novel {
	return Novel{
		Title:      novel.Title,
		Cover:      novel.Cover,
		Author:     novel.Author,
		Subscribed: novel.Subscribed,
		Ended:      novel.Ended,
		Level:      novel.Level,
	}
}

// BuildArchiveList 序列化文章列表
func BuildNovelList(novels []models.Novel) []Novel {
	var novelList []Novel
	for _, a := range novels {
		i := Novel{
			ID:         a.ID,
			Title:      a.Title,
			Cover:      a.Cover,
			Author:     a.Author,
			Subscribed: a.Subscribed,
			Ended:      a.Ended,
			Level:      a.Level,
		}
		novelList = append(novelList, i)
	}
	return novelList
}

// BuildArchiveResponse 序列化文章响应
func BuildNovelResponse(novel models.Novel) NovelResponse {
	return NovelResponse{
		Data: BuildNovel(novel),
	}
}

// BuildArchiveResponse 序列化文章列表响应
func BuildNovelListResponse(novels []models.Novel, count int, next bool, last bool, pages int) NovelListResponse {
	return NovelListResponse{
		Count: count,
		Data:  BuildNovelList(novels),
		Next:  next,
		Last:  last,
		Pages: pages,
	}
}
