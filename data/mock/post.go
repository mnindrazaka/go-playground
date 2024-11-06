package mock

import (
	"golang-playground/domain/post"
)

type ProductRepository struct{}

func NewProductRepository() post.ProductRepository {
	return ProductRepository{}
}

func (repository ProductRepository) CreatePost(post post.Post) error {
	return nil
}

func (repository ProductRepository) DeletePost(id string) error {
	return nil
}

func (repository ProductRepository) GetPostList() ([]post.Post, error) {
	return []post.Post{
		{Id: 1, Title: "Contoh 1", Content: "testing", UserId: 1},
		{Id: 2, Title: "Contoh 2", Content: "testing", UserId: 1},
		{Id: 3, Title: "Contoh 3", Content: "testing", UserId: 1},
		{Id: 4, Title: "Contoh 4", Content: "testing", UserId: 1},
		{Id: 5, Title: "Contoh 5", Content: "testing", UserId: 1},
	}, nil
}
