package controllers

import (
	"encoding/json"
	"fmt"
	"jwt-golang/helpers"
	"jwt-golang/models"
	"net/http"
)

type Cart struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Id_produk string `json:"id_produk"`
	Jml       string `json:"jml"`
	Tot       string `json:"tot"`
}

func CartHandler(writer http.ResponseWriter, request *http.Request) {
	err := helpers.Refresh(writer, request)
	if err != nil {
		fmt.Fprintf(writer, err.Error())
		return
	}

	var cart Cart

	c, _ := request.Cookie("username")
	userStr := c.Value

	switch request.Method {
	case "DELETE":

		err = json.NewDecoder(request.Body).Decode(&cart)
		if err != nil {
			fmt.Fprintf(writer, "invalid body")
			return
		}
		if cart.Id == "" {
			fmt.Fprintf(writer, "invalid body")
			return
		}

		add := models.DeleteCart(cart.Id)

		if add["status"] != "200" {
			fmt.Fprintf(writer, add["msg"])
			return
		}

		fmt.Fprintf(writer, "data berhasil dihapus!")
	case "POST":

		err = json.NewDecoder(request.Body).Decode(&cart)
		if err != nil {
			fmt.Fprintf(writer, "invalid body")
			return
		}
		if cart.Id_produk == "" || cart.Jml == "" {
			fmt.Fprintf(writer, "invalid body")
			return
		}

		add := models.AddCart(userStr, cart.Id_produk, cart.Jml)

		if add["status"] != "200" {
			fmt.Fprintf(writer, add["msg"])
			return
		}

		fmt.Fprintf(writer, "data berhasil diinput!")
	case "GET":
		data := models.GetCart(userStr)
		fmt.Fprintf(writer, data)
	}
}
