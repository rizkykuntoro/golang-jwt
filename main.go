package main

import (
	"jwt-golang/controllers"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/logout", controllers.LogoutHandler)
	http.HandleFunc("/register", controllers.RegisterHandler)

	http.HandleFunc("/produk", controllers.ProdukHandler)
	http.HandleFunc("/cart", controllers.CartHandler)

	http.HandleFunc("/checkout", controllers.TransactionHandler)

	http.ListenAndServe(":8080", nil)
}
