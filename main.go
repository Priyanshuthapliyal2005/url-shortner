package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/priyanshuthapliyal/go-url-shortner/database"
	"github.com/priyanshuthapliyal/go-url-shortner/middlewares"
	"github.com/priyanshuthapliyal/go-url-shortner/routes"
)

type Data struct {
	Url string `json:"url"`
}

type UrlSchema struct {
	Url  string
	Hash string
}

func main() {

	// Load .env file
	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found. Relying on environment variables.")
    } else {
        log.Println(".env file loaded successfully.")
    }

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("PORT not set. Using default:", port)
	} else {
		log.Println("PORT set to:", port)
	}

	router := gin.Default()

	mongoClient := database.MongoDB()

	redisClient := database.Redis()

	router.Use(middlewares.DBMiddleWare(mongoClient, redisClient))

	routes.Routes(router)

	log.Fatal(router.Run("0.0.0.0:" + port))

	defer func() {
		ok := mongoClient.Disconnect(context.TODO())
		if ok != nil {
			panic(ok)
		}
	}()
}
