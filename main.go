package main

import (
	"fmt"
	"time"

	//"userapi/app"

	"github.com/dgrijalva/jwt-go"
)

var mykey = []byte("mysecretkey")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Arun Kumar"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mykey)
	if err != nil {
		fmt.Println("Something went wrong", err.Error())
		return "", err
	}

	return tokenString, nil
}

func main() {
	//app.StartApplication()

	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Println("Error generating token string")
	}
	fmt.Println(tokenString)
}
