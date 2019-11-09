package rules

import "net/http"

type MethodEquals struct {
	method string
}

func (r *MethodEquals) Matches(request *http.Request) bool {
	return r.method == request.Method
}
