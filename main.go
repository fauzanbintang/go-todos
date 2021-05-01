package main

import (
	"golang-ptp/go-todos/routers"
	"log"
	"os"

	"golang-ptp/go-todos/docs"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := routers.SetupRouter()
	port := os.Getenv("port")

	if port == "" {
		port = "8080"
	}

	docs.SwaggerInfo.Title = "Example Swagger Todo"
	docs.SwaggerInfo.Description = "Documentation API Todo"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":" + port)
}
