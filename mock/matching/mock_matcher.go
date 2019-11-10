package matching

import (
	"../../mock"
	"net/http"
)

type MockMatcher interface {
	Match(request *http.Request) (*mock.Mock, error)
}
