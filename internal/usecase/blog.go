package usecase

type BlogUseCase struct {
	ArticleRepo
	Advertise
}

// New -.
func New(r ArticleRepo, w Advertise) *BlogUseCase {
	return &BlogUseCase{
		ArticleRepo: r,
		Advertise:   w,
	}
}
