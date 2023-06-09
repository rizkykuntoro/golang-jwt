package controllers

import (
	"fmt"
	"jwt-golang/helpers"
	"jwt-golang/models"
	"net/http"
)

func TransactionHandler(writer http.ResponseWriter, request *http.Request) {
	err := helpers.Refresh(writer, request)
	if err != nil {
		fmt.Fprintf(writer, err.Error())
		return
	}

	c, _ := request.Cookie("username")
	userStr := c.Value

	switch request.Method {
	case "POST":

		add := models.AddTrx(userStr)

		if add["status"] != "200" {
			fmt.Fprintf(writer, add["msg"])
			return
		}

		fmt.Fprintf(writer, "data berhasil diinput!")
	case "GET":
		data := models.GetTrx(userStr)
		fmt.Fprintf(writer, data)
	}
}
