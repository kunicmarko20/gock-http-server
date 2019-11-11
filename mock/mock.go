package mock

import (
	matching "github.com/kunicmarko20/gock-http-server/mock/matching/rules"
	"net/http"
)

type Mock struct {
	name     string
	rule     matching.Rule
	response *Response
}

func NewMock(name string, rule matching.Rule, response *Response) *Mock {
	return &Mock{
		name,
		rule,
		response,
	}
}

func (m *Mock) Matches(request *http.Request) bool {
	return m.rule.Matches(request)
}

func (m *Mock) Name() string {
	return m.name
}

func (m *Mock) Response() *Response {
	return m.response
}
