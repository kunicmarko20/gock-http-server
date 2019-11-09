package rules

import (
	"net/http"
	"strings"
)

type PathPrefixEquals struct {
	path string
}

func NewPathPrefixEquals(path string) *PathPrefixEquals {
	return &PathPrefixEquals{path}
}

func (r *PathPrefixEquals) Matches(request *http.Request) bool {
	return strings.HasPrefix(request.URL.Path, r.path)
}
