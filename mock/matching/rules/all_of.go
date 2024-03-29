package rules

import "net/http"

type AllOf struct {
	rules []Rule
}

func NewAllOf(rules []Rule) *AllOf {
	return &AllOf{rules}
}

func (r *AllOf) Matches(request *http.Request) bool {
	for _, rule := range r.rules {
		if !rule.Matches(request) {
			return false
		}
	}

	return true
}
