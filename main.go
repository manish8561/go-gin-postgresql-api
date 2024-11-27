package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/manish8561/go-gin-postgresql-api/pkg/books"
	"github.com/manish8561/go-gin-postgresql-api/pkg/common/db"
)

func init() {
	err := godotenv.Load("./pkg/common/envs/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ginMode := os.Getenv("GIN_MODE")

	gin.SetMode(ginMode)
}
func main() {

	err := godotenv.Load("./pkg/common/envs/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	router := gin.Default()
	dbHandler := db.Init(dbUrl)

	books.RegisterRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

	router.Run(port)
}
