package routes

import (
	"blog/server"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Author    string `json:"author"`
	CreatedAt string `json:"createdAt"`
}

func GetPosts(ctx *gin.Context) {
	db := server.App().DB
	res, err := db.Query("SELECT * FROM posts")
	if err != nil {
		panic(err.Error())
	}
	var allposts []Post
	for res.Next() {
		var posts Post
		err = res.Scan(&posts.Id, &posts.Title, &posts.Body, &posts.Author, &posts.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		allposts = append(allposts, posts)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": allposts})
}

func GetSinglePost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
