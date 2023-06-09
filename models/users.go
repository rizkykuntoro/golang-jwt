package models

import (
	"context"
	"jwt-golang/helpers"
	"log"
)

type Users struct {
	Username string
	Password string
	Role     string
}

func Login(username, password string) map[string]string {
	q := "SELECT username, password, role FROM users WHERE username='" + username + "'"

	res, err := Init().Query(q)

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	var data_login map[string]string

	for res.Next() {

		var user Users
		err := res.Scan(&user.Username, &user.Password, &user.Role)

		if err != nil {
			log.Fatal(err)
		}

		match := helpers.CheckPasswordHash(password, user.Password)
		if match {
			data_login = map[string]string{
				"status":   "200",
				"username": user.Username,
				"role":     user.Role}
		} else {
			data_login = map[string]string{
				"status":   "401",
				"username": user.Username}
		}
	}

	return data_login
}

func AddUser(username, password string) map[string]string {

	q := "SELECT username, password, role FROM users WHERE username='" + username + "'"
	res, err := Init().Query(q)
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}

	var insert map[string]string

	if res.Next() == false {

		pass_hashed, _ := helpers.HashPassword(password)
		query := "INSERT INTO `users` (`username`, `password`, `role`) VALUES (?, ?, 'user')"
		_, err = Init().ExecContext(context.Background(), query, username, pass_hashed)

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

	insert = map[string]string{
		"status": "40",
		"msg":    "username sudah terdaftar!"}

	return insert
}

func CheckRole(username string) string {
	q := "SELECT username, password, role FROM users WHERE username='" + username + "'"

	res, err := Init().Query(q)

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	var role string

	for res.Next() {

		var user Users
		err := res.Scan(&user.Username, &user.Password, &user.Role)

		if err != nil {
			log.Fatal(err)
		}

		role = user.Role
	}

	return role
}
