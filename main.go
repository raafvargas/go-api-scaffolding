package main

import (
	"go-api-scaffolding/controllers"
	"go-api-scaffolding/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "go-api-scaffolding/docs"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Some API
// @version 1.0
// @description API para gerênciar a busca de leilões.

// @BasePath /v1
func main() {
	Environment()

	router := gin.Default()
	router.Use(middlewares.CORS)

	v1 := router.Group("/api/models")
	{
		v1.Use(middlewares.Authentication)

		modelController := new(controllers.ModelController)
		v1.GET("/", modelController.Find)
		v1.GET("/:modelID", modelController.Get)
	}

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(404)
	})

	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	router.GET("/swagger/*any", swaggerHandler)

	router.Run(":9000")
}

//Environment variables
func Environment() {
	err := godotenv.Load()
	if err != nil {
		err := godotenv.Load("secrets/.env")
		if err != nil {
			panic(".env wasn't found.")
		}
	}
}
