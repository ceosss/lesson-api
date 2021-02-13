package cookiehandler

import (
	"net/http"
	"time"

	"github.com/ceosss/lesson-api/helper/customerror"
	"github.com/ceosss/lesson-api/helper/jwtkey"
	"github.com/ceosss/lesson-api/models"
	"github.com/dgrijalva/jwt-go"
)

// VerifyCookie ...
func VerifyCookie(response http.ResponseWriter, request *http.Request) error {
	cookie, err := request.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {

			customerror.Unauthorized(&response, err)
			return err
		}
		customerror.BadRequest(&response, err)
		return err
	}

	tokenString := cookie.Value

	claims := &models.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtkey.GetJwtKey(), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			customerror.Unauthorized(&response, err)
			return err
		}
		customerror.BadRequest(&response, err)
		return err
	}
	if !token.Valid {
		customerror.Unauthorized(&response, err)
		return err
	}
	return nil
}

// GenerateJWT ...
func GenerateJWT(response *http.ResponseWriter, email string) error {

	expirationTime := time.Now().Add(10 * time.Minute)

	claims := models.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtkey.GetJwtKey())

	if err != nil {
		customerror.InternalServerError(response, err)
		return err
	}

	http.SetCookie(*response, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	return nil
}
