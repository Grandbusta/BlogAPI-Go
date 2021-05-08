package main

import (
	"blog/routes"
	"blog/server"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found")
	}
}

func main() {
	app := server.App().Router
	app.GET("/posts", routes.GetPosts)
	app.GET("/posts/:id", routes.GetSinglePost)
	app.Run(":8081")
}
