package lessonhandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ceosss/lesson-api/helper/db"
	"github.com/ceosss/lesson-api/models"
)

// CreateLesson ...
func CreateLesson(response http.ResponseWriter, request *http.Request) {
	var lesson models.Lesson
	var err error

	err = json.NewDecoder(request.Body).Decode(&lesson)

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	client, err := db.ConnectToDB()

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	lessonCollection := db.GetLessonCollection(client)

	res, err := lessonCollection.InsertOne(context.TODO(), lesson)

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	res.InsertedID = ""

	response.Header().Set("content-type", "application/json")
	response.WriteHeader(201)
	response.Write([]byte(`{  "response": "Lesson Created Successfully"}`))
}
