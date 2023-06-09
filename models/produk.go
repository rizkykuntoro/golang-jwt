package models

import (
	"context"
	"database/sql"
)

type Produk struct {
	Id          string `json:"id"`
	Nama_produk string `json:"nama_produk"`
	Stok        string `json:"stok"`
	Harga       string `json:"harga"`
}

func AddProduk(nama_produk, stok, harga string) map[string]string {

	var insert map[string]string
	query := "INSERT INTO `produk` (`nama_produk`, `stok`, `harga`) VALUES (?, ?, ?)"
	_, err := Init().ExecContext(context.Background(), query, nama_produk, stok, harga)

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

func DeleteProduk(id_produk string) map[string]string {

	var delete map[string]string
	query := "DELETE FROM `produk` WHERE id = ?"
	_, err := Init().ExecContext(context.Background(), query, id_produk)

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

func EditProduk(id_produk, nama_produk, stok, harga string) map[string]string {

	var update map[string]string
	query := "UPDATE `produk` SET "

	if nama_produk != "" && stok != "" && harga != "" {
		query += "nama_produk = '" + nama_produk + "', stok = '" + stok + "' , harga = '" + harga + "' "
	} else if nama_produk != "" {
		query += "nama_produk = '" + nama_produk + "' "
	} else if stok != "" {
		query += "stok = '" + stok + "' "
	} else if harga != "" {
		query += "harga = '" + harga + "' "
	}
	query += "WHERE id = '" + id_produk + "'"
	_, err := Init().ExecContext(context.Background(), query)

	if err != nil {
		update = map[string]string{
			"status": "400",
			"msg":    err.Error()}

		return update

	}

	update = map[string]string{
		"status": "200"}

	return update
}

func GetProduk() string {
	q, _ := getJSON("SELECT * FROM produk")
	return q
}

func DetailProduk(id_produk string) map[string]string {
	var id string
	var nama_produk string
	var stok string
	var harga string
	var detail map[string]string
	sqlStatement := `SELECT * FROM produk WHERE id="` + id_produk + `"`
	row := Init().QueryRow(sqlStatement)

	switch err := row.Scan(&id, &nama_produk, &stok, &harga); err {
	case sql.ErrNoRows:
		detail = map[string]string{
			"status": "400",
			"msg":    "No rows were returned!"}
	case nil:
		detail = map[string]string{
			"status": "200",
			"stok":   stok,
			"harga":  harga}
	default:
		detail = map[string]string{
			"status": "400",
			"msg":    err.Error()}
	}
	return detail
}
