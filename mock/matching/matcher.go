package matching

import (
	"../../mock"
	"errors"
	"net/http"
)

type MockMatcher struct {
	MockRepository *mock.Repository
}

func (m *MockMatcher) Match (request *http.Request) (*mock.Mock, error) {
	for _, value := range m.MockRepository.All() {
		if value.Matches(request) {
			return value, nil
		}
	}

	return nil, errors.New("mock not matched")
}
