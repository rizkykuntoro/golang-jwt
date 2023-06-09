package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"jwt-golang/models"
	"net/http"
)

func RegisterHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		var user User

		err := json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			fmt.Fprintf(writer, "invalid body")
			return
		}

		add := models.AddUser(user.Username, user.Password)

		if add["status"] != "200" {
			fmt.Fprintf(writer, add["msg"])
			return
		}

		fmt.Fprintf(writer, "silahkan Login")

	case "GET":
		fmt.Fprintf(writer, "only POST methods is allowed.")
		return
	}
}

func ifNotAdmin(w http.ResponseWriter, r *http.Request) (err error) {
	c, err := r.Cookie("username")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintf(w, "can not find token in cookies")
			w.WriteHeader(http.StatusUnauthorized)
			return

		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userStr := c.Value

	user_role := models.CheckRole(userStr)
	if user_role != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		return errors.New("Anda tidak diizinkan")
	}

	return nil
}
