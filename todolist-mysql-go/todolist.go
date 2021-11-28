package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//I think this is currently the problem
//var db, _ := gorm.Open("mysql", "root:root@/todolist?charset=utf8&parseTime=True&loc=Local")

//create a struct that will contain all of the info in our database table
type TodoItemModel struct {
	Id          int `gorm:"primary_key"`
	Description string
	Completed   bool
}

//function to respond to the /healthz endpoint
func Healthz(w http.ResponseWriter, r *http.Request) {
	log.Info("API Health is Ok")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	db, _ := gorm.Open(mysql.Open("root:root@/todolist?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})

	defer db.Close() //close the database connection when the main function stops running (when it's returned)

	db.Debug().DropTableIfExists(&TodoItemModel{})
	db.Debug().AutoMigrate(&TodoItemModel{}) //will automigrate the struct into a table in the database

	log.Info("Starting Todolist API server")
	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz).Methods("GET") //create an endpoint that will call the Healthz function
	http.ListenAndServe(":8000", router)
}
