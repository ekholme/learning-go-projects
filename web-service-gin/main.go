package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
//this creates a struct with the following conventions
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
//the syntax []album indicates a slice with elements of type album
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID) //note that the colon preceding the path signifies that id is a parameter
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

//getAlbums creates JSON from the slice of albums, writing out a JSON response
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

//postAlbums adds an album from JSON received in the request body.
//other note -- go doesn't really care about the order in which we declare functions; just that we use them if we declare them
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//getAlbumByID locates the album whose ID value matches the id
//parameter sent by the client, then returns that album as a response
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	//loop over list of albums, looking for an album that matches the parameter
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
