package helpers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateJWT(writer http.ResponseWriter, request *http.Request, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	expirationTime := time.Now().Add(time.Minute * 1)
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = expirationTime.Unix()

	tokenString, err := token.SignedString(sampleSecretKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	http.SetCookie(writer, &http.Cookie{
		Name:    "username",
		Value:   username,
		Expires: expirationTime,
	})
	return "selamat datang " + username, nil
}

func ValidateToken(w http.ResponseWriter, r *http.Request) (err error) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintf(w, "can not find token in cookies")
			w.WriteHeader(http.StatusUnauthorized)
			return

		}
		fmt.Fprintf(w, "can not find token in cookies")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tknStr := c.Value

	token, err := jwt.Parse(tknStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return sampleSecretKey, nil
	})

	if token == nil {
		fmt.Fprintf(w, "invalid token")
		w.WriteHeader(http.StatusUnauthorized)
		return errors.New("Token error")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Fprintf(w, "couldn't parse claims")
		return errors.New("Token error")
	}

	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		fmt.Fprintf(w, "token expired")
		return errors.New("Token error")
	}

	return nil
}

func Refresh(w http.ResponseWriter, r *http.Request) (err error) {
	err = ValidateToken(w, r)
	if err != nil {
		return err
	}

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

	_, err = GenerateJWT(w, r, userStr)
	return err
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: "",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "username",
		Value: "",
	})
}
