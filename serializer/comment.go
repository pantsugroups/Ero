package serializer

import (
	"eroauz/models"
	"time"
)

type Comment struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	AuthorID     uint      `json:"author_id"`
	AuthorName   string    `json:"author_name"`
	AuthorAvatar string    `json:"author_avatar"`
	Type         int       `json:"type"`
	RId          uint      `json:"raw"`
	RTitle       string    `json:"raw_title"`
	RCid         uint      `json:"reply"`
	RCTitle      string    `json:"reply_title"`
	CreateAt     time.Time `json:"time"`
}

type CommentResponse struct {
	Response
	Data Comment `json:"data"`
}

type CommentListResponse struct {
	Response
	Count int       `json:"count"`
	All   int       `json:"all"`
	Data  []Comment `json:"data"`
	Next  bool      `json:"have_next"`
	Last  bool      `json:"have_last"`
	Pages int       `json:"pages"`
}

func BuildComment(comment models.Comment) Comment {
	c := Comment{
		ID:           comment.ID,
		Title:        comment.Title,
		AuthorID:     comment.Author.ID,
		AuthorName:   comment.Author.Nickname,
		AuthorAvatar: comment.Author.Avatar,
		Type:         comment.Type,
		RId:          comment.RId,
		RCid:         comment.RCid,
		CreateAt:     comment.CreatedAt,
	}
	if c.RCid != 0 {
		rc, err := models.GetComment(c.RCid)
		if err == nil {
			c.RCTitle = rc.Title
		}
	}
	if c.Type == models.Archive_ {
		a, err := models.GetArchive(c.RId)
		if err != nil {
			c.RTitle = "被删除文章"
		} else {
			c.RTitle = a.Title
		}
	} else {
		n, err := models.GetNovel(c.RId)
		if err != nil {
			c.RTitle = "被删除小说"
		} else {
			c.RTitle = n.Title
		}
	}

	return c
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

func BuildCommentListResponse(comments []models.Comment, all int, count int, next bool, last bool, pages int) CommentListResponse {
	return CommentListResponse{
		Count: count,
		All:   all,
		Data:  BuildCommentList(comments),
		Next:  next,
		Last:  last,
		Pages: pages,
	}
}
