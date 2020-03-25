package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// TODO

// Post json scheme
type Post struct {
	Title string
	User  string
	Post  string
	Date  string
}

var db *sql.DB = nil

// ValidateUser
func ValidateUser(username, password string) bool {
	// bind user
	if db == nil {
		initDatabase("/home/mrk1/go/src/github.com/mrkaurelius/gobackend/database/db.db")
		log.Println("Database initialized")
	} else {
		log.Println("Database already initialized")
	}

	rows, err := db.Query("SELECT username, password  FROM users WHERE username=(?)", username)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// defer rows.Close()
	var usernameDB string
	var passwordDB string

	for rows.Next() {
		err = rows.Scan(&usernameDB, &passwordDB)
		if err != nil {
			fmt.Println(err)
			return false
		}

		if password == passwordDB {
			fmt.Println(usernameDB, passwordDB)
			return true
		}
	}

	return false
}

// UserPostsJSON, handle errors
func UserPostsJSON(username string) []byte {
	if db == nil {
		initDatabase("/home/mrk1/go/src/github.com/mrkaurelius/gobackend/database/db.db")
	}

	posts := make([]Post, 0)

	// bind user
	rows, _ := db.Query("SELECT user, date, post, title FROM posts WHERE user=(?)", username)
	defer rows.Close()

	for rows.Next() {
		var tmpPost Post
		rows.Scan(&tmpPost.User, &tmpPost.Date, &tmpPost.Post, &tmpPost.Title)
		posts = append(posts, tmpPost)
	}
	// fmt.Println(posts)

	postsJSON, err := json.Marshal(posts)
	if err != nil {
		fmt.Println(err)
	}

	return postsJSON
}

// AllUserPostsJSON returns all posts for home page
func AllUserPostsJSON() []byte {
	if db == nil {
		initDatabase("/home/mrk1/go/src/github.com/mrkaurelius/gobackend/database/db.db")
	}

	posts := make([]Post, 0)

	// order by desc
	rows, err := db.Query("SELECT user, date, post, title FROM posts ORDER BY date DESC")
	//defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var tmpPost Post
		err := rows.Scan(&tmpPost.User, &tmpPost.Date, &tmpPost.Post, &tmpPost.Title)
		if err != nil {
			//log.Fatal(err)
		}

		posts = append(posts, tmpPost)
	}
	// fmt.Println(posts)

	PostsJSON, err := json.Marshal(posts)
	if err != nil {
		fmt.Println(err)
	}

	return PostsJSON
}

func initDatabase(path string) {
	db, _ = sql.Open("sqlite3", path)
}
