package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"os"

	db "mini-wallet/database"
	routes "mini-wallet/routes/v1"
)

func init() {
	db.DatabaseConnect()
}

func main() {
	port := os.Getenv("SERVER_PORT")

	if port == "" {
		port = "3000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.IndexRouter(router)

	router.Run("localhost:3000")
}
