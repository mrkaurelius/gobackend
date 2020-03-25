package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mrkaurelius/gobackend/database"
	"github.com/mrkaurelius/gobackend/jwt"
)

// TODO
// understand jwt, auth

// not on rest standarts
// get all posts
// get post in detail
// auth

// try to understand abstractions

// connect with database

func userAuth(c *gin.Context) {
	// connect to databaase and check user cridentals
	username := c.PostForm("username")
	password := c.PostForm("password")

	pwbyte := []byte(password)
	sha256 := sha256.Sum256([]byte(pwbyte))

	pwsha256hex := hex.EncodeToString(sha256[:32])

	// fmt.Printf("'%s','%s'\n", username, pwsha256hex)

	if database.ValidateUser(username, pwsha256hex) {
		// return token
		// jwtString := "dummy, Jwt"
		header := gin.H{"token": jwt.CreateToken(username)}
		c.JSON(200, header)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func allUserPosts(c *gin.Context) {
	// connect to databaase and check user cridentals
	posts := database.AllUserPostsJSON()
	//fmt.Println(posts)
	postsString := string(posts)

	c.String(http.StatusOK, postsString)
}

// todo
func verifyUser(c *gin.Context) {

}

// todo
func addPost(c *gin.Context) {

}

func main() {
	router := gin.Default()

	router.POST("/auth", userAuth)
	router.GET("/posts", allUserPosts)
	//router.GET("/verify", verifyUser)

	// router.GET("/user/:name", userPost)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// func userPost(c *gin.Context) {
// 	// name := c.Param("name")
// 	// c.String(http.StatusOK, "Hello %s", name)
// }
