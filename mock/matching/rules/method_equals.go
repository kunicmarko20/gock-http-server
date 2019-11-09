package rules

import "net/http"

type MethodEquals struct {
	method string
}

func NewMethodEquals(method string) *MethodEquals {
	return &MethodEquals{method}
}

func (r *MethodEquals) Matches(request *http.Request) bool {
	return r.method == request.Method
}
