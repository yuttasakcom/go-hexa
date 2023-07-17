package store

import (
	"context"
)

type IStore interface {
	Create(c context.Context, entity interface{}) error
	Update(c context.Context, entity interface{}) error
	Delete(c context.Context, entity interface{}) error
	GetByID(c context.Context, entity interface{}, id int) (interface{}, error)
	Get(c context.Context, entity interface{}) ([]interface{}, error)
}

type Store struct {
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Create(c context.Context, entity interface{}) error {
	return nil
}

func (s *Store) Update(c context.Context, entity interface{}) error {
	return nil
}

func (s *Store) Delete(c context.Context, entity interface{}) error {
	return nil
}

func (s *Store) GetByID(c context.Context, entity interface{}, id int) (interface{}, error) {
	return nil, nil
}

func (s *Store) Get(c context.Context, entity interface{}) ([]interface{}, error) {
	return nil, nil
}
