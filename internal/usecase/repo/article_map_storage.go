package repo

import (
	"fiber_blog/internal/entity"
	"fiber_blog/pkg/mapStorage"
)

type ArticleRepo struct {
	articlesCount int
	*mapStorage.Storage
}

func New(ms *mapStorage.Storage) *ArticleRepo {
	return &ArticleRepo{Storage: ms}
}

func (r *ArticleRepo) Add(a entity.Article) (entity.ArticleID, error) {
	r.articlesCount += 1
	a.ID = entity.ArticleID(r.articlesCount)
	return a.ID, r.Storage.Add(a, r.articlesCount)
}

func (r *ArticleRepo) Delete(id entity.ArticleID) error {
	r.Storage.Delete(int(id))
	return nil
}

func (r *ArticleRepo) Update(a entity.Article) error {
	r.Storage.Update(int(a.ID), a)
	return nil
}

func (r *ArticleRepo) Retrieve(aID entity.ArticleID) (entity.Article, error) {
	article, err := r.Storage.Retrieve(int(aID))
	if err != nil {
		return entity.Article{}, err
	}
	return article.(entity.Article), nil
}
