package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mrkaurelius/gobackend/database"
	"github.com/mrkaurelius/gobackend/jwt"
)

// TODO

func userAuth(c *gin.Context) {
	// connect to databaase and check user cridentals
	username := c.PostForm("username")
	password := c.PostForm("password")

	body, _ := ioutil.ReadAll(c.Request.Body)
	println(string(body))
	fmt.Println(username, password)

	pwbyte := []byte(password)
	sha256 := sha256.Sum256([]byte(pwbyte))

	pwsha256hex := hex.EncodeToString(sha256[:32])
	// fmt.Printf("'%s','%s'\n", username, pwsha256hex)

	if database.ValidateUser(username, pwsha256hex) {
		header := gin.H{"token": jwt.CreateToken(username)}
		c.JSON(200, header)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func allUserPosts(c *gin.Context) {
	// connect to database and check user cridentals
	posts := database.AllUserPostsJSON()
	//fmt.Println(posts)
	postsString := string(posts)

	c.String(http.StatusOK, postsString)
}

func addPost(c *gin.Context) {
	token := c.PostForm("token")

	var post database.Post
	username, ret := jwt.VerifyToken(token)

	if ret {
		//fmt.Println(post)

		post.Title = c.PostForm("title")
		post.Post = c.PostForm("post")
		post.Date = time.Now().Format("2006-01-02")
		post.User = username
		fmt.Println(username)
		database.AddPostDB(post)

		c.Status(http.StatusOK)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}

func main() {
	router := gin.Default()

	router.POST("/auth", userAuth)
	router.GET("/posts", allUserPosts)
	router.POST("/posts", addPost)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
