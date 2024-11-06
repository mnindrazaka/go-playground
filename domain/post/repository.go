package post

type ProductRepository interface {
	CreatePost(post Post) error
	DeletePost(id string) error
	GetPostList() ([]Post, error)
}
