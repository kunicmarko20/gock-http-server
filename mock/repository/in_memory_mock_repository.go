package repository

import (
	"../../mock"
	"errors"
)

type InMemoryMockRepository struct {
	mocks map[string]*mock.Mock
}

func NewInMemoryMockRepository() *InMemoryMockRepository {
	return &InMemoryMockRepository{make(map[string]*mock.Mock)}
}

func (r *InMemoryMockRepository) All() map[string]*mock.Mock {
	return r.mocks
}

func (r *InMemoryMockRepository) Reset() {
	r.mocks = make(map[string]*mock.Mock)
}

func (r *InMemoryMockRepository) Set(mock *mock.Mock) {
	r.mocks[mock.Name()] = mock
}

func (r *InMemoryMockRepository) Get(mockName string) (*mock.Mock, error) {
	if value, ok := r.mocks[mockName]; ok {
		return value, nil
	}

	return nil, errors.New("mock not found")
}
