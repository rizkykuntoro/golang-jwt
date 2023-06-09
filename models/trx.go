package models

import (
	"context"
	"log"
	"strconv"
)

type Trx struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Total    string `json:"total"`
	Status   string `json:"status"`
	Datetime string `json:"datetime"`
}

type Cart struct {
	Tot       string `json:"tot"`
	Id_produk string `json:"id_produk"`
	Jml       string `json:"jml"`
}

func AddTrx(username string) map[string]string {

	str_total := GetTotalTrx(username)

	var insert map[string]string
	// query := "INSERT INTO `trx` (`username`, `total`) VALUES (?, ?)"
	// _, err := Init().ExecContext(context.Background(), query, username, str_total)

	query := "INSERT INTO `trx` (`username`, `total`) VALUES (?, ?)"
	insertResult, err := Init().ExecContext(context.Background(), query, username, str_total)
	if err != nil {
		insert = map[string]string{
			"status": "400",
			"msg":    err.Error()}

		return insert

	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}

	str_last_id := string(id)
	InsertTrxDetail(username, str_last_id)
	insert = map[string]string{
		"status": "200"}

	return insert
}

func InsertTrxDetail(username, id_trx string) string {
	q := "SELECT tot, id_produk, jml FROM cart WHERE username='" + username + "'"
	res, err := Init().Query(q)
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}

	var total int
	for res.Next() {
		var cart Cart
		err := res.Scan(&cart.Tot, &cart.Id_produk, &cart.Jml)
		if err != nil {
			log.Fatal(err)
		}

		query := "INSERT INTO `trx_detail` (`tot`, `id_trx`, `id_produk`, `jml`) VALUES (?, ?)"
		_, err = Init().ExecContext(context.Background(), query, cart.Tot, id_trx, cart.Id_produk, cart.Jml)
	}

	str_total := strconv.Itoa(total)

	query := "DELETE FROM `cart` WHERE username = ?"
	_, err = Init().ExecContext(context.Background(), query, username)

	return str_total
}

func GetTotalTrx(username string) string {
	q := "SELECT tot, id_produk, jml FROM cart WHERE username='" + username + "'"
	res, err := Init().Query(q)
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}

	var total int
	for res.Next() {
		var cart Cart
		err := res.Scan(&cart.Tot, &cart.Id_produk, &cart.Jml)
		if err != nil {
			log.Fatal(err)
		}
		int_tot, _ := strconv.Atoi(cart.Tot)
		total += int_tot

		updated_stok := MinusStok(cart.Id_produk, cart.Jml)

		EditProduk(cart.Id_produk, updated_stok["nama_produk"], updated_stok["stok_final"], updated_stok["harga"])
	}

	str_total := strconv.Itoa(total)

	// query := "DELETE FROM `cart` WHERE username = ?"
	// _, err = Init().ExecContext(context.Background(), query, username)

	return str_total
}

func MinusStok(id_produk, jml string) map[string]string {
	var stok string
	var nama_produk string
	var harga string
	var data map[string]string

	sqlStatement := `SELECT stok, nama_produk, harga FROM produk WHERE id="` + id_produk + `"`
	row := Init().QueryRow(sqlStatement)

	err := row.Scan(&stok, &nama_produk, &harga)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("data_db : " + stok + nama_produk + harga)

	stokint, _ := strconv.Atoi(stok)
	jmlint, _ := strconv.Atoi(jml)

	last_stok := stokint - jmlint

	stok_final := strconv.Itoa(last_stok)

	// fmt.Println("stok final " + stok_final)

	data = map[string]string{
		"stok_final":  stok_final,
		"nama_produk": nama_produk,
		"harga":       harga}
	println(data)
	return data
}

func EditProduk_(id_produk, nama_produk, stok, harga string) map[string]string {

	var update map[string]string
	query := "UPDATE `produk` SET "

	if nama_produk != "" && stok != "" && harga != "" {
		query += "nama_produk = '" + nama_produk + "', stok = '" + stok + "' , stok = '" + harga + "' "
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

func GetTrx(username string) string {
	q, _ := getJSON("SELECT * FROM produk")
	return q
}
