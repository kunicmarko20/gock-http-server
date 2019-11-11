package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kunicmarko20/gock-http-server/mock"
	"github.com/kunicmarko20/gock-http-server/mock/matching/rules"
	"github.com/kunicmarko20/gock-http-server/propertyaccess"
	"net/http"
)

type PayloadToMockTransformer struct {
	propertyAccessor *propertyaccess.PropertyAccessor
}

func NewPayloadToMockTransformer(propertyAccessor *propertyaccess.PropertyAccessor) *PayloadToMockTransformer {
	return &PayloadToMockTransformer{propertyAccessor}
}

func (p PayloadToMockTransformer) FromRequest(request *http.Request) (*mock.Mock, error) {
	decoder := json.NewDecoder(request.Body)

	var payload Payload

	err := decoder.Decode(&payload)

	if err != nil {
		return nil, fmt.Errorf("unable to deserialize json, with error: %s", err)
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
	case "jsonRequestPayloadPropertyEquals":
		return rules.NewJsonRequestPayloadPropertyEquals(matchRule["property"].(string), matchRule["value"], p.propertyAccessor), nil
	default:
		return nil, fmt.Errorf("unknown type: \"%v\"", matchRule["type"])
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
