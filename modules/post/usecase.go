package post

type ProductUsecase struct {
	Repository ProductRepository
}

func (usecase ProductUsecase) CreatePostUsecase(post Post) error {
	err := usecase.Repository.CreatePost(post)
	return err
}

func (usecase ProductUsecase) DeletePostUsecase(id string) error {
	err := usecase.Repository.DeletePost(id)
	return err
}
