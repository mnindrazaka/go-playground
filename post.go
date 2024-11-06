package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	Id      int
	Title   string
	Content string
	UserId  int
	User    User
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// body -> []byte
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// []byte -> Post
	var post Post
	err = json.Unmarshal(bodyByte, &post)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	result := db.Table("posts").Where("id = ?", id).Updates(&post)
	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}

	w.Write([]byte("Success"))
}
