package post_test

import (
	"golang-playground/data/mock"
	"golang-playground/domain/post"
	"testing"
)

func TestGetPostList(t *testing.T) {
	productRepository := mock.NewProductRepository()
	productUsecase := post.ProductUsecase{Repository: productRepository}

	posts, _ := productUsecase.GetPostList()

	if posts[0].Title != "Contoh 1" {
		t.Error("logic usecase ada yang salah")
	}
}
