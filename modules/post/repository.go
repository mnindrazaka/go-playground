package post

import "gorm.io/gorm"

type ProductRepository struct {
	DB *gorm.DB
}

func (repository ProductRepository) CreatePost(post Post) error {
	result := repository.DB.Table("posts").Create(&post)
	return result.Error
}

func (repository ProductRepository) DeletePost(id string) error {
	result := repository.DB.Table("posts").Where("id = ?", id).Delete(Post{})
	return result.Error
}

func (repository ProductRepository) GetPostList() ([]Post, error) {
	var posts []Post
	result := repository.DB.Table("posts").Joins("User").Find(&posts)
	return posts, result.Error
}
