package main

import "time"

type user struct {
	name      string
	username  string
	email     string
	password  string // change to hash
	followers []user
	following []user
	posts     []post
	comments  []comment
}

func newUser(
	name string,
	username string,
	email string,
	password string,
) *user {

	u := user{
		name:      name,
		username:  username,
		email:     email,
		password:  password,
		followers: make([]user, 0),
		following: make([]user, 0),
		posts:     make([]post, 0),
		comments:  make([]comment, 0),
	}

	return &u
}

type post struct {
	author   user
	content  string
	date     time.Time
	comments []comment
}

func newPost(author user, content string, date time.Time) *post {
	p := post{
		author:   author,
		content:  content,
		date:     date,
		comments: make([]comment, 0),
	}

	return &p

}

type comment struct {
	author   user
	content  string
	date     time.Time
	comments []comment
}

func newComment(author user, content string, date time.Time) *comment {
	c := comment{
		author:   author,
		content:  content,
		date:     date,
		comments: make([]comment, 0),
	}

	return &c
}
