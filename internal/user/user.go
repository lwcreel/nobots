package user

import (
	"context"
	"net/http"

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

func PostUsers(conn *pgx.Conn) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var newUser User

		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		query := `INSERT INTO users (name, username, email, passhash) VALUES (@name, @username, @email, @passhash) ON CONFLICT DO NOTHING;`
		args := pgx.NamedArgs{
			"name":     newUser.Name,
			"username": newUser.Username,
			"email":    newUser.Email,
			"passhash": newUser.Passhash,
		}

		row, err := conn.Query(context.Background(), query, args)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		defer row.Close()

		c.JSON(http.StatusCreated, newUser)
	}
	return gin.HandlerFunc(fn)
}

func GetUsers(conn *pgx.Conn) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var users []User
		var params queryParams
		var query string
		var err error

		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
			query = `SELECT * FROM users WHERE id=@id;`
			rows, _ := conn.Query(context.Background(), query, args)
			users, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[User])
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	}
	return gin.HandlerFunc(fn)
}

func PutUsers(conn *pgx.Conn) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var updatedUser User

		if err := c.BindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		query := `UPDATE users SET name=@name, username=@username, email=@email, passhash=@passhash WHERE id=@id;`
		args := pgx.NamedArgs{
			"id":       updatedUser.Id,
			"name":     updatedUser.Name,
			"username": updatedUser.Username,
			"email":    updatedUser.Email,
			"passhash": updatedUser.Passhash,
		}

		row, err := conn.Query(context.Background(), query, args)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		defer row.Close()

		c.JSON(http.StatusOK, updatedUser)
	}
	return gin.HandlerFunc(fn)
}

func DeleteUsers(conn *pgx.Conn) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var params queryParams
		var query string

		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query = `DELETE FROM users WHERE id=@id;`
		args := pgx.NamedArgs{
			"id": params.Id,
		}

		row, err := conn.Query(context.Background(), query, args)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		defer row.Close()

		c.Status(http.StatusNoContent)
	}
	return gin.HandlerFunc(fn)
}
