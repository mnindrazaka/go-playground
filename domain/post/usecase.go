package post

type ProductUsecase struct {
	Repository ProductRepository
}

func (usecase ProductUsecase) CreatePost(post Post) error {
	err := usecase.Repository.CreatePost(post)
	return err
}

func (usecase ProductUsecase) DeletePost(id string) error {
	err := usecase.Repository.DeletePost(id)
	return err
}

func (usecase ProductUsecase) GetPostList() ([]Post, error) {
	posts, err := usecase.Repository.GetPostList()
	return posts, err
}
