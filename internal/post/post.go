package post

import (
	"github.com/lwcreel/nobots/internal/comment"
	"github.com/lwcreel/nobots/internal/user"
	"time"
)

type Post struct {
	Id       string
	Author   user.User
	Content  string
	Date     time.Time
	Comments []comment.Comment
}

func newPost(Author user.User, Content string, Date time.Time) *Post {
	p := Post{
		Author:   Author,
		Content:  Content,
		Date:     Date,
		Comments: make([]comment.Comment, 0),
	}

	return &p

}
