package mapStorage

import (
	"errors"
	"fmt"
)

type Storage struct {
	storage map[int]interface{}
}

func New() *Storage {
	return &Storage{storage: map[int]interface{}{}}
}

func (s *Storage) Add(object interface{}, key int) error {
	if _, exist := s.storage[key]; exist {
		return errors.New("article already exist")
	}
	s.storage[key] = object
	return nil
}

func (s *Storage) Delete(key int) {
	delete(s.storage, key)
}

func (s *Storage) Update(key int, object interface{}) {
	s.storage[key] = object
}

func (s *Storage) Retrieve(key int) (interface{}, error) {
	article, exist := s.storage[key]
	if exist {
		return article, nil
	}
	return nil, fmt.Errorf("article %d doesn't exist", key)
}
