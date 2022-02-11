package controllers

import (
	"net/http"

	"github.com/ekholme/learning-go-projects/gin-gorm-api/models"
	"github.com/gin-gonic/gin"
)

//GET /books
//function to return all books in db
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}
