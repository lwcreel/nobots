package user

type User struct {
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
) *User {

	u := User{
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
