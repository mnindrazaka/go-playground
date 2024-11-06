package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id       int
	Username string
	Password string
	Posts    []Post
}

var jwtSecret = []byte("INI_RAHASIA")

func Login(w http.ResponseWriter, r *http.Request) {

	var userBody User

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal(bodyBytes, &userBody)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var userResult User
	result := DB.Table("users").Where("username = ? AND password = ?", userBody.Username, userBody.Password).First(&userResult)
	if result.Error != nil {
		w.Write([]byte("wrong username or password"))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   userResult.Id,
		"username": userResult.Username,
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	cookie := http.Cookie{
		Name:     "jwt_token",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	w.Write([]byte("success"))
}

func CheckLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtTokenCookie, err := r.Cookie("jwt_token")
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte("not authenticated"))
			return
		}

		_, err = jwt.Parse(jwtTokenCookie.Value, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte("token not valid"))
			return
		}

		next(w, r)
	}
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	var users []User
	result := DB.Table("users").Preload("Posts").Find(&users)
	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}

	usersByte, err := json.Marshal(users)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(usersByte)
}
