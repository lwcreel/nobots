package user

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

// ====== typedef ======
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

// ====== HTTP Requests ======

type queryParams struct {
	Id string `form:"id" query:"id"`
}

func GetUsers(conn *pgx.Conn) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var users []User
		var params queryParams
		var query string
		var err error

		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		args := pgx.NamedArgs{
			"id": params.Id,
		}

		if params.Id == "" {
			query = `SELECT * FROM users;`
			rows, _ := conn.Query(context.Background(), query)
			users, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[User])
		} else {
			query = `SELECT * FROM users WHERE id=@id`
			rows, _ := conn.Query(context.Background(), query, args)
			users, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[User])
		}

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, users)
	}
	return gin.HandlerFunc(fn)
}

func PostUsers(conn *pgx.Conn) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var newUser User

		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}

		query := `INSERT INTO users (name, username, email, passhash) VALUES (@name, @username, @email, @passhash) ON CONFLICT DO NOTHING`
		args := pgx.NamedArgs{
			"name":     newUser.Name,
			"username": newUser.Username,
			"email":    newUser.Email,
			"passhash": newUser.Passhash,
		}

		_, err := conn.Query(context.Background(), query, args)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}

		c.IndentedJSON(http.StatusCreated, newUser)
	}
	return gin.HandlerFunc(fn)
}
