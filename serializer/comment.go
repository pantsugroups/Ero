package serializer

import (
	"eroauz/models"
)

type Comment struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Author_id   uint   `json:"author_id"`
	Author_Name string `json:"author_name"`
	Type        int    `json:"type"`
	RId         uint   `json:"raw"`
}

// ArchiveResponse 单个用户序列化
type CommentResponse struct {
	Response
	Data Comment `json:"data"`
}

// ArchiveResponse 单个用户序列化
type CommentListResponse struct {
	Response
	Count int       `json:"count"`
	Data  []Comment `json:"data"`
	Next  bool      `json:"have_next"`
	Last  bool      `json:"have_last"`
	Pages int       `json:"pages"`
}

// BuildArchive 单个序列化文章
func BuildComment(comment models.Comment) Comment {
	return Comment{
		ID:          comment.ID,
		Title:       comment.Title,
		Author_id:   comment.Author.ID,
		Author_Name: comment.Author.Nickname,
		Type:        comment.Type,
		RId:         comment.RId,
	}
}

// BuildArchiveList 序列化文章列表
func BuildCommentList(comments []models.Comment) []Comment {
	var commentList []Comment
	for _, a := range comments {
		i := BuildComment(a)
		commentList = append(commentList, i)
	}
	return commentList
}

// BuildArchiveResponse 序列化文章响应
func BuildCommentResponse(comment models.Comment) CommentResponse {
	return CommentResponse{
		Data: BuildComment(comment),
	}
}

// BuildArchiveResponse 序列化文章列表响应
func BuildCommentListResponse(commments []models.Comment, count int, next bool, last bool, pages int) CommentListResponse {
	return CommentListResponse{
		Count: count,
		Data:  BuildCommentList(commments),
		Next:  next,
		Last:  last,
		Pages: pages,
	}
}
