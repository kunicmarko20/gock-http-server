package mock

import "errors"

type Repository struct {
	mocks map[string]*Mock
}

func (r *Repository) All() map[string]*Mock {
	return r.mocks
}

func (r *Repository) Reset() {
	r.mocks = make(map[string]*Mock)
}

func (r *Repository) Set(mock *Mock) {
	r.mocks[mock.name] = mock
}

func (r *Repository) Get(mockName string) (*Mock, error) {
	if mock, ok := r.mocks[mockName]; ok {
		return mock, nil
	}

	return nil, errors.New("mock not found")
}
