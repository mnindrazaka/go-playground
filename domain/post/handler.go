package post

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	Usecase ProductUsecase
}

func (handler ProductHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var post Post
	err = json.Unmarshal(bodyByte, &post)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = handler.Usecase.CreatePost(post)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Success"))
}

func (handler ProductHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := handler.Usecase.DeletePost(id)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Success"))
}

func (handler ProductHandler) GetPostList(w http.ResponseWriter, r *http.Request) {
	posts, err := handler.Usecase.GetPostList()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	postsByte, err := json.Marshal(posts)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(postsByte)
}
