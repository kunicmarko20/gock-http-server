package rules

import (
	"net/http"
	"strings"
)

type PathEquals struct {
	path string
}

func NewPathEquals(path string) *PathEquals {
	return &PathEquals{path}
}

func (r *PathEquals) Matches(request *http.Request) bool {
	return r.path == strings.TrimPrefix(request.URL.Path, "/mock/")
}
