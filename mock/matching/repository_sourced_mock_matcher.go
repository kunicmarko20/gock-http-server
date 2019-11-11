package matching

import (
	"errors"
	"github.com/kunicmarko20/gock-http-server/mock"
	"github.com/kunicmarko20/gock-http-server/mock/repository"
	"net/http"
)

type RepositorySourcedMockMatcher struct {
	mockRepository repository.MockRepository
}

func NewRepositorySourcedMockMatcher(mockRepository repository.MockRepository) *RepositorySourcedMockMatcher {
	return &RepositorySourcedMockMatcher{mockRepository}
}

func (m *RepositorySourcedMockMatcher) Match(request *http.Request) (*mock.Mock, error) {
	for _, value := range m.mockRepository.All() {
		if value.Matches(request) {
			return value, nil
		}
	}

	return nil, errors.New("mock not matched")
}
