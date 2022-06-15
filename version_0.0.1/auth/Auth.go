package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"todo_service/version_0.0.1/responses"

	jwt "github.com/dgrijalva/jwt-go"
)

// var charlist string= "_${'`'}{|}~123abcde.fmnopqlABCDE@FJKLMNOPQRSTUVWXYZ456789stuvwxyz0!#$%&ijkrgh'*+-/=?^";
var mySigningKey = []byte("captainjacksparrowsayshi")

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Header().Set("Content-Type", "application/json")
				err_resp := responses.AuthError{Status: 401, Message: err.Error()}
				json.NewEncoder(w).Encode(err_resp)
				return
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {

			// fmt.Fprintf(w, "Not Authorized")
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			err_resp := responses.AuthError{Status: 401, Message: "Missing or broken token"}
			json.NewEncoder(w).Encode(err_resp)
			return
		}
	})
}

func GenerateJWT(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = email
	claims["exp"] = time.Now().Add(time.Minute * 0).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
