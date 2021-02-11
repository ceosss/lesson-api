package authhandlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/ceosss/lesson-api/helper/customerror"
	"github.com/ceosss/lesson-api/helper/db"
	"github.com/ceosss/lesson-api/helper/jwtkey"
	"github.com/ceosss/lesson-api/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
)

// Login ...
func Login(response http.ResponseWriter, request *http.Request) {
	var User models.User
	var err error

	err = json.NewDecoder(request.Body).Decode(&User)
	if err != nil {
		customerror.BadRequest(&response, err)
		return
	}

	if User.Email != "swaraj@inspiritvr.com" || User.Password != "swaraj12" {
		customerror.Unauthorized(&response, errors.New("Invalid Email or Password"))
		return
	}

	expirationTime := time.Now().Add(10 * time.Minute)

	claims := models.Claims{
		Email: User.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtkey.GetJwtKey())

	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}

	http.SetCookie(response, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

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

	client, err := db.ConnectToDB()
	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}

	userCollection := db.GetUserCollection(client)

	res, err := userCollection.InsertOne(context.TODO(), user)

	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}

	res.InsertedID = ""

	response.Header().Set("content-type", "application/json")
	response.WriteHeader(200)
	response.Write([]byte(`{  "response": "User Registered Successfully"}`))

}
