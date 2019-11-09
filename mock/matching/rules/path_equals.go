package rules

import "net/http"

type PathEquals struct {
	path string
}

func NewPathEquals(path string) *PathEquals {
	return &PathEquals{path}
}

func (r *PathEquals) Matches(request *http.Request) bool {
	return r.path == request.URL.Path
}
