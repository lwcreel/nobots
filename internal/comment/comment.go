package comment

import (
	"github.com/lwcreel/nobots/internal/user"
	"time"
)

type Comment struct {
	Id       string
	Author   user.User
	Content  string
	Date     time.Time
	Comments []Comment
}

func newComment(author user.User, content string, date time.Time) *Comment {
	c := Comment{
		Author:   author,
		Content:  content,
		Date:     date,
		Comments: make([]Comment, 0),
	}

	return &c
}
