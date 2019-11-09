package rules

import "net/http"

type Rule interface {
	Matches(request *http.Request) bool
}
