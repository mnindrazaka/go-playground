package main

import (
	"fmt"
	"golang-playground/data/db"
	"golang-playground/domain/post"
	"golang-playground/presentation/api"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func main() {
	DB, err = gorm.Open(mysql.Open("root:((root))@tcp(127.0.0.1:3306)/news?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err.Error())
	}

	productRepository := db.NewProductRepository(DB)
	productUsecase := post.ProductUsecase{Repository: productRepository}
	productHandler := api.ProductHandler{Usecase: productUsecase}

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
