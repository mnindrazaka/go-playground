package main

import (
	"fmt"
	"golang-playground/modules/post"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open(mysql.Open("root:((root))@tcp(127.0.0.1:3306)/news?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err.Error())
	}

	productRepository := post.ProductRepository{DB: db}
	productUsecase := post.ProductUsecase{Repository: productRepository}
	productHandler := post.ProductHandler{Usecase: productUsecase}

	router := mux.NewRouter()

	router.HandleFunc("/posts/{id}", CheckLogin(UpdatePost)).Methods("PUT")

	router.HandleFunc("/posts/{id}", CheckLogin(productHandler.CreatePost)).Methods("DELETE")
	router.HandleFunc("/posts", CheckLogin(productHandler.DeletePost)).Methods("POST")
	router.HandleFunc("/posts", CheckLogin(productHandler.GetPostList)).Methods("GET")
	router.HandleFunc("/users", CheckLogin(GetUserList)).Methods("GET")
	router.HandleFunc("/login", Login).Methods("POST")

	err = http.ListenAndServe(":3000", router)
	fmt.Println(err)
}
