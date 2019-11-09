package rules

import "net/http"

type AnyOf struct {
	rules []Rule
}

func (r *AnyOf) Matches(request *http.Request) bool {
	for _, rule := range r.rules {
		if rule.Matches(request) {
			return true
		}
	}

	return false
}
