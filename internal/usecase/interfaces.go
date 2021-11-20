// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"fiber_blog/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// ArticleRepo -.
	ArticleRepo interface {
		Add(entity.Article) (entity.ArticleID, error)
		Delete(entity.ArticleID) error
		Update(entity.Article) error
		Retrieve(entity.ArticleID) (entity.Article, error)
	}
)
