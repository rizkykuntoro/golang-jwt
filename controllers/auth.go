package controllers

import (
	"encoding/json"
	"fmt"
	"jwt-golang/helpers"
	"jwt-golang/models"
	"net/http"
)

// Uer type
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		var user User

		err := json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			fmt.Fprintf(writer, "invalid body")
			return
		}

		login := models.Login(user.Username, user.Password)

		if login["status"] != "200" {
			fmt.Fprintf(writer, "can not authenticate this user")
			return
		}

		token, err := helpers.GenerateJWT(writer, request, user.Username)
		if err != nil {
			fmt.Fprintf(writer, "error in generating token")
		}

		fmt.Fprintf(writer, token)

	case "GET":
		fmt.Fprintf(writer, "only POST methods is allowed.")
		return
	}
}

func LogoutHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		helpers.Logout(writer, request)
	}
}
