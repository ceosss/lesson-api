package authhandlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/ceosss/lesson-api/helper/customerror"
	"github.com/ceosss/lesson-api/helper/jwtkey"
	"github.com/ceosss/lesson-api/models"
	"github.com/dgrijalva/jwt-go"
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
