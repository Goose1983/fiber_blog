package webapi

import "fiber_blog/internal/entity"

type AdWebAPI struct {
	url string
}

func New(url string) *AdWebAPI {
	return &AdWebAPI{url: url}
}

func (a *AdWebAPI) GetAd() (entity.Ad, error) {
	// Do Some GET Request to a.url and return retrieved ad or error
	return entity.Ad{}, nil
}
