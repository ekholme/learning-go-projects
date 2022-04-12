//holding onto code so i can play around with firebase -- should just do some version control, but w/e
// package tmp

// import (
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// type dare struct {
// 	ID       string `json:"id" binding:"required"`
// 	Title    string `json:"title" binding:"required"`
// 	Text     string `json:"text" binding:"required"`
// 	Savagery int    `json:"savagery" binding:"required,gte=1,lte=10"`
// }

// var dares = []dare{
// 	{ID: "1", Title: "Sockhands", Text: "wear socks on your hands", Savagery: 3},
// 	{ID: "2", Title: "Bustin Makes me Feel Good", Text: "drink lots of beer while listening to Ghostbusters", Savagery: 6},
// }

// func main() {
// 	r := gin.Default()

// 	//serve form & post form data
// 	r.Static("/forms", "./forms")
// 	r.POST("/forms/testform", formDare)

// 	r.GET("/dares", getDares)
// 	r.POST("/dares", postDares)

// 	r.Run("localhost:8080")

// }

// //function to get all dares
// func getDares(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, dares)
// }

// func postDares(c *gin.Context) {
// 	var newDare dare

// 	if err := c.BindJSON(&newDare); err != nil {
// 		return
// 	}

// 	dares = append(dares, newDare)

// 	c.IndentedJSON(http.StatusCreated, newDare)
// }

// func formDare(c *gin.Context) {

// 	idVal := c.PostForm("id")
// 	titleVal := c.PostForm("title")
// 	textVal := c.PostForm("text")
// 	savVal, err := strconv.Atoi(c.PostForm("savagery"))

// 	if err != nil {
// 		log.Fatal("couldn't convert savagery value")
// 	}

// 	newDare := &dare{idVal, titleVal, textVal, savVal}

// 	dares = append(dares, *newDare)

// 	c.JSON(http.StatusCreated, gin.H{"message": "Dare submitted!"})
// }

//need to render the form and then try out post. process might be to have a page
//that renders all of the current data, then have the form at the bottom?

//possible form help stuff
//https://stackoverflow.com/questions/39215772/marshalling-html-form-input-to-json-and-writing-to-file-golang
//https://stackoverflow.com/questions/48909476/post-html-form-with-golang-gin-backend