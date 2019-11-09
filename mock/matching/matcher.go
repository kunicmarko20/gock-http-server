package matching

import (
	"../../mock"
	"../repository"
	"errors"
	"net/http"
)

type MockMatcher struct {
	MockRepository repository.MockRepository
}

func (m *MockMatcher) Match(request *http.Request) (*mock.Mock, error) {
	for _, value := range m.MockRepository.All() {
		if value.Matches(request) {
			return value, nil
		}
	}

	return nil, errors.New("mock not matched")
}
