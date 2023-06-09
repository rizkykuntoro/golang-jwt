package models

import (
	"context"
	"strconv"
)

func AddCart(username, id_produk, jml string) map[string]string {
	var insert map[string]string
	detail_produk := DetailProduk(id_produk)

	stok, _ := strconv.Atoi(detail_produk["stok"])
	harga, _ := strconv.Atoi(detail_produk["harga"])
	jumlah_beli, _ := strconv.Atoi(jml)

	if jumlah_beli > stok {
		insert = map[string]string{
			"status": "400",
			"msg":    "Stok produk tidak cukup"}

		return insert
	}

	tot := jumlah_beli * harga
	total_belanja := strconv.Itoa(tot)

	query := "INSERT INTO `cart` (`username`, `id_produk`, `jml`, `tot`) VALUES (?, ?, ?, ?)"
	_, err := Init().ExecContext(context.Background(), query, username, id_produk, jml, total_belanja)

	if err != nil {
		insert = map[string]string{
			"status": "400",
			"msg":    err.Error()}

		return insert

	}

	insert = map[string]string{
		"status": "200"}

	return insert
}

func DeleteCart(id_cart string) map[string]string {

	var delete map[string]string
	query := "DELETE FROM `cart` WHERE id = ?"
	_, err := Init().ExecContext(context.Background(), query, id_cart)

	if err != nil {
		delete = map[string]string{
			"status": "400",
			"msg":    err.Error()}

		return delete

	}

	delete = map[string]string{
		"status": "200"}

	return delete
}

func GetCart(username string) string {
	q, _ := getJSON("SELECT * FROM cart WHERE username = '" + username + "'")
	return q
}
