package mock

import (
	matching "./matching/rules"
	"net/http"
)

type Mock struct {
	name     string
	rule     matching.Rule
	Response Response
}

func (m *Mock) Matches(request *http.Request) bool {
	return m.rule.Matches(request)
}
