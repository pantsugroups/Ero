package serializer

import (
	"eroauz/models"
)

type Comment struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	AuthorID   uint   `json:"author_id"`
	AuthorName string `json:"author_name"`
	Type       int    `json:"type"`
	RId        uint   `json:"raw"`
	RUid       uint   `json:"reply"`
}

type CommentResponse struct {
	Response
	Data Comment `json:"data"`
}

type CommentListResponse struct {
	Response
	Count int       `json:"count"`
	Data  []Comment `json:"data"`
	Next  bool      `json:"have_next"`
	Last  bool      `json:"have_last"`
	Pages int       `json:"pages"`
}

func BuildComment(comment models.Comment) Comment {
	return Comment{
		ID:         comment.ID,
		Title:      comment.Title,
		AuthorID:   comment.Author.ID,
		AuthorName: comment.Author.Nickname,
		Type:       comment.Type,
		RId:        comment.RId,
	}
}

func BuildCommentList(comments []models.Comment) []Comment {
	var commentList []Comment
	for _, a := range comments {
		i := BuildComment(a)
		commentList = append(commentList, i)
	}
	return commentList
}

func BuildCommentResponse(comment models.Comment) CommentResponse {
	return CommentResponse{
		Data: BuildComment(comment),
	}
}

func BuildCommentListResponse(commments []models.Comment, count int, next bool, last bool, pages int) CommentListResponse {
	return CommentListResponse{
		Count: count,
		Data:  BuildCommentList(commments),
		Next:  next,
		Last:  last,
		Pages: pages,
	}
}
