package matching

import (
	"github.com/kunicmarko20/gock-http-server/mock"
	"net/http"
)

type MockMatcher interface {
	Match(request *http.Request) (*mock.Mock, error)
}
