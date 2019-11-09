package http

import (
	"../mock"
	"../mock/matching/rules"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type PayloadToMockTransformer struct{}

func (p PayloadToMockTransformer) FromRequest(request *http.Request) (*mock.Mock, error) {
	decoder := json.NewDecoder(request.Body)

	var payload Payload

	err := decoder.Decode(&payload)

	if err != nil {
		return nil, errors.New("unable to deserialize json")
	}

	vars := mux.Vars(request)

	rule, err := p.transformMatchRule(payload.MatchRule)

	if err != nil {
		return nil, err
	}

	return mock.NewMock(vars["mock"], rule, payload.Response), nil
}

func (p PayloadToMockTransformer) transformMatchRule(matchRule map[string]interface{}) (rules.Rule, error) {
	switch matchRule["type"] {
	case "allOf":
		children, err := p.transformChildRules(matchRule["rules"].([]interface{}))

		if err != nil {
			return nil, err
		}

		return rules.NewAllOf(children), nil
	case "anyOf":
		children, err := p.transformChildRules(matchRule["rules"].([]interface{}))

		if err != nil {
			return nil, err
		}

		return rules.NewAnyOf(children), nil
	case "methodEquals":
		return rules.NewMethodEquals(matchRule["value"].(string)), nil
	case "pathEquals":
		return rules.NewPathEquals(matchRule["value"].(string)), nil
	case "pathPrefixEquals":
		return rules.NewPathPrefixEquals(matchRule["value"].(string)), nil
	default:
		return nil, errors.New(fmt.Sprintf("Unknown type: \"%v\"", matchRule["type"]))
	}
}

func (p PayloadToMockTransformer) transformChildRules(childRules []interface{}) ([]rules.Rule, error) {
	var children []rules.Rule

	for _, rule := range childRules {
		transformedRule, err := p.transformMatchRule(rule.(map[string]interface{}))

		if err != nil {
			return nil, err
		}

		children = append(children, transformedRule)
	}

	return children, nil
}