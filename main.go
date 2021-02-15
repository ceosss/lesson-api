package main

import (
	"fmt"
	"net/http"

	"github.com/ceosss/lesson-api/helper/db"
	"github.com/ceosss/lesson-api/router"
	"go.mongodb.org/mongo-driver/mongo"
)

// // ModelCollection ...
// var ModelCollection = db.GetModelCollection()

// // LessonCollection ...
// var LessonCollection = db.GetLessonCollection()

// // UserCollection ...
// var UserCollection = db.GetUserCollection()

var client *mongo.Client

// ModelCollection ...
var ModelCollection *mongo.Collection

// LessonCollection ...
var LessonCollection *mongo.Collection

// UserCollection ...
var UserCollection *mongo.Collection

func main() {
	fmt.Println("START")
	client := db.ConnectToDB()
	ModelCollection = db.GetModelCollection(client)
	LessonCollection = db.GetLessonCollection(client)
	UserCollection = db.GetUserCollection(client)
	// Call router to get a newly configured router
	router := router.NewRouter()
	http.ListenAndServe(":3000", router)
}
