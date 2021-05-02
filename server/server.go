package server

import (
	"blog/db"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	DB     *sql.DB
	Router *gin.Engine
}

func App() (app *Server) {
	database, err := db.DbConfig()
	if err != nil {
		panic(err.Error())
	}
	server := &Server{
		Router: gin.Default(),
		DB:     database,
	}
	return server
}
