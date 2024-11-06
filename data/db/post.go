package db

import (
	"golang-playground/domain/post"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) post.ProductRepository {
	return ProductRepository{DB: DB}
}

func (repository ProductRepository) CreatePost(post post.Post) error {
	result := repository.DB.Table("posts").Create(&post)
	return result.Error
}

func (repository ProductRepository) DeletePost(id string) error {
	result := repository.DB.Table("posts").Where("id = ?", id).Delete(post.Post{})
	return result.Error
}

func (repository ProductRepository) GetPostList() ([]post.Post, error) {
	var posts []post.Post
	result := repository.DB.Table("posts").Joins("User").Find(&posts)
	return posts, result.Error
}
