package main

import (
	"github.com/ekholme/learning-go-projects/gin-gorm-api/controllers"
	"github.com/ekholme/learning-go-projects/gin-gorm-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)

	r.Run()
}

//RESUME AT "TO CREATE A BOOK..."
