package authhandlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ceosss/lesson-api/helper/cookiehandler"
	"github.com/ceosss/lesson-api/helper/customerror"
	"github.com/ceosss/lesson-api/helper/password"
	"github.com/ceosss/lesson-api/helper/successresponse"
	"github.com/ceosss/lesson-api/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

// Login ...
func Login(response http.ResponseWriter, request *http.Request) {
	var User models.User
	var UserFromDB models.User
	var err error

	err = json.NewDecoder(request.Body).Decode(&User)
	if err != nil {
		customerror.BadRequest(&response, err)
		return
	}

	// client, err := db.ConnectToDB()

	// if err != nil {
	// 	customerror.InternalServerError(&response, err)
	// 	return
	// }

	// userCollection := db.GetUserCollection(client)

	filter := bson.M{"email": User.Email}
	UserCollection.FindOne(context.TODO(), filter).Decode(&UserFromDB)

	if !password.DecodePassword(UserFromDB.Password, User.Password) {
		customerror.InternalServerError(&response, errors.New("Invlaid Email or Password"))
		return
	}

	err = cookiehandler.GenerateJWT(&response, User.Email)
	if err != nil {
		return
	}
	successresponse.OK(&response)
}

// Register ...
func Register(response http.ResponseWriter, request *http.Request) {
	var err error
	var user models.User

	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		customerror.BadRequest(&response, err)
		return
	}

	v := validator.New()
	err = v.Struct(user)
	if err != nil {
		customerror.BadRequest(&response, err)
		return
	}

	user.Password, err = password.EncodePassword(user.Password)

	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}

	// client, err := db.ConnectToDB()
	// if err != nil {
	// 	customerror.InternalServerError(&response, err)
	// 	return
	// }

	// userCollection := db.GetUserCollection(client)

	res, err := UserCollection.InsertOne(context.TODO(), user)

	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}

	res.InsertedID = ""

	successresponse.OK(&response)
	response.Write([]byte(`{  "response": "User Registered Successfully"}`))

}
