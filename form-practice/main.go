package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type dare struct {
	ID       string `json:"id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Text     string `json:"text" binding:"required"`
	Savagery uint   `json:"savagery" binding:"required,gte=1,lte=10"`
}

var dares = []dare{
	{ID: "1", Title: "Sockhands", Text: "wear socks on your hands", Savagery: 3},
	{ID: "2", Title: "Bustin Makes me Feel Good", Text: "drink lots of beer while listening to Ghostbusters", Savagery: 6},
}

func main() {
	r := gin.Default()

	r.GET("/dares", getDares)
	r.POST("/dares", postDares)

	r.Run("localhost:8080")

}

//function to get all dares
func getDares(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dares)
}

func postDares(c *gin.Context) {
	var newDare dare

	if err := c.BindJSON(&newDare); err != nil {
		return
	}

	dares = append(dares, newDare)

	c.IndentedJSON(http.StatusCreated, newDare)
}

//RESUME -- CHECK OUT THIS PAGE
//https://stackoverflow.com/questions/39215772/marshalling-html-form-input-to-json-and-writing-to-file-golang
