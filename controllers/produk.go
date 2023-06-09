package controllers

import (
	"encoding/json"
	"fmt"
	"jwt-golang/helpers"
	"jwt-golang/models"
	"net/http"
)

type Produk struct {
	Id          string `json:"id"`
	Nama_produk string `json:"nama_produk"`
	Stok        string `json:"stok"`
	Harga       string `json:"harga"`
}

func ProdukHandler(writer http.ResponseWriter, request *http.Request) {
	err := helpers.Refresh(writer, request)
	if err != nil {
		fmt.Fprintf(writer, err.Error())
		return
	}

	switch request.Method {
	case "DELETE":
		err := ifNotAdmin(writer, request)
		if err != nil {
			fmt.Fprintf(writer, err.Error())
			return
		}

		var produk Produk

		err = json.NewDecoder(request.Body).Decode(&produk)
		if err != nil {
			fmt.Fprintf(writer, "invalid body")
			return
		}
		if produk.Id == "" {
			fmt.Fprintf(writer, "invalid body")
			return
		}

		add := models.DeleteProduk(produk.Id)

		if add["status"] != "200" {
			fmt.Fprintf(writer, add["msg"])
			return
		}

		fmt.Fprintf(writer, "data berhasil dihapus!")
	case "PUT":
		err := ifNotAdmin(writer, request)
		if err != nil {
			fmt.Fprintf(writer, err.Error())
			return
		}

		var produk Produk

		err = json.NewDecoder(request.Body).Decode(&produk)
		if err != nil {
			fmt.Fprintf(writer, "invalid body")
			return
		}
		if produk.Id == "" {
			fmt.Fprintf(writer, "invalid body")
			return
		}

		add := models.EditProduk(produk.Id, produk.Nama_produk, produk.Stok, produk.Harga)

		if add["status"] != "200" {
			fmt.Fprintf(writer, add["msg"])
			return
		}

		fmt.Fprintf(writer, "data berhasil diupdate!")
	case "POST":
		err := ifNotAdmin(writer, request)
		if err != nil {
			fmt.Fprintf(writer, err.Error())
			return
		}

		var produk Produk

		err = json.NewDecoder(request.Body).Decode(&produk)
		if err != nil {
			fmt.Fprintf(writer, "invalid body")
			return
		}
		if produk.Nama_produk == "" || produk.Stok == "" || produk.Harga == "" {
			fmt.Fprintf(writer, "invalid body")
			return
		}

		add := models.AddProduk(produk.Nama_produk, produk.Stok, produk.Harga)

		if add["status"] != "200" {
			fmt.Fprintf(writer, add["msg"])
			return
		}

		fmt.Fprintf(writer, "data berhasil diinput!")
	case "GET":
		data := models.GetProduk()

		fmt.Fprintf(writer, data)
	}
}
