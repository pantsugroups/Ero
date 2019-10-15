package serializer

import (
	"eroauz/models"
)

// Archive Info
//
// swagger:response ArchiveResponse
type Archive struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	JapTitle       string `json:"japanese_title"`
	Content        string `json:"content"`
	Cover          string `json:"cover"`
	Author         string `json:"author"`
	PrimaryContent string `json:"primary_content"`
	CreatedAt      int64  `json:"created_at"`
	UpdateAt       int64  `json:"update_at"`
	CreateId       uint   `json:"create_id"`
	CreateName     string `json:"create_name"`
	Tag            string `json:"tag"`
}

// ArchiveResponse 单个用户序列化
type ArchiveResponse struct {
	Response
	Data Archive `json:"data"`
}

// ArchiveResponse 单个用户序列化
type ArchiveListResponse struct {
	Response
	Count int       `json:"count"`
	All   int       `json:"all"`
	Data  []Archive `json:"data"`
	Next  bool      `json:"have_next"`
	Last  bool      `json:"have_last"`
	Pages int       `json:"pages"`
}

// BuildArchive 单个序列化文章
func BuildArchive(archive models.Archive) Archive {
	return Archive{
		ID:             archive.ID,
		Title:          archive.Title,
		JapTitle:       archive.JapTitle,
		Cover:          archive.Cover,
		Content:        archive.Content,
		Author:         archive.Author,
		PrimaryContent: archive.PrimaryContent,
		CreatedAt:      archive.CreatedAt.Unix(),
		CreateId:       archive.Create.ID,
		UpdateAt:       archive.UpdatedAt.Unix(),
		CreateName:     archive.Create.Nickname,
		Tag:            archive.Tag,
	}
}

// BuildArchiveList 序列化文章列表
func BuildArchiveList(archives []models.Archive) []Archive {
	var archiveList []Archive
	for _, a := range archives {
		i := BuildArchive(a)
		archiveList = append(archiveList, i)
	}
	return archiveList
}

// BuildArchiveResponse 序列化文章响应
func BuildArchiveResponse(archive models.Archive) ArchiveResponse {
	return ArchiveResponse{
		Data: BuildArchive(archive),
	}
}

// BuildArchiveResponse 序列化文章列表响应
func BuildArchiveListResponse(archives []models.Archive, all int, count int, next bool, last bool, pages int) ArchiveListResponse {
	return ArchiveListResponse{
		Count: count,
		All:   all,
		Data:  BuildArchiveList(archives),
		Next:  next,
		Last:  last,
		Pages: pages,
	}
}
