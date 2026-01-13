package main

import "time"

type user struct {
	Id       string `db:"id"`
	Name     string `db:"name"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Passhash string `db:"passhash"` // change to hash
	// Followers []user
	// Following []user
	// Posts     []post
	// Comments  []comment
}

func newUser(
	Name string,
	Username string,
	Email string,
	Passhash string,
) *user {

	u := user{
		Name:     Name,
		Username: Username,
		Email:    Email,
		Passhash: Passhash,
		//		Followers: make([]user, 0),
		//		Following: make([]user, 0),
		//		Posts:     make([]post, 0),
		//		Comments:  make([]comment, 0),
	}

	return &u
}

type post struct {
	Id       string
	Author   user
	Content  string
	Date     time.Time
	Comments []comment
}

func newPost(Author user, Content string, Date time.Time) *post {
	p := post{
		Author:   Author,
		Content:  Content,
		Date:     Date,
		Comments: make([]comment, 0),
	}

	return &p

}

type comment struct {
	Id       string
	Author   user
	Content  string
	Date     time.Time
	Comments []comment
}

func newComment(author user, content string, date time.Time) *comment {
	c := comment{
		Author:   author,
		Content:  content,
		Date:     date,
		Comments: make([]comment, 0),
	}

	return &c
}
