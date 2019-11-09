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

type PayloadFactory struct{}

func (p PayloadFactory) FromRequest(request *http.Request) (*mock.Mock, error) {
	decoder := json.NewDecoder(request.Body)

	var decodedJson map[string]interface{}

	err := decoder.Decode(&decodedJson)

	if err != nil {
		return nil, errors.New("unable to deserialize json")
	}

	vars := mux.Vars(request)

	rule, err := p.transformMatchRule(decodedJson["matchRule"].(map[string]interface{}))

	if err != nil {
		return nil, err
	}

	return mock.NewMock(vars["mock"], rule, p.createResponse(decodedJson["response"].(map[string]interface{}))), nil
}

func (p PayloadFactory) transformMatchRule(matchRule map[string]interface{}) (rules.Rule, error) {
	switch matchRule["type"] {
	case "allOf":
		children, err := p.transformChildRules(matchRule["rules"].(map[string]interface{}))

		if err != nil {
			return nil, err
		}

		return rules.NewAllOf(children), nil
	case "anyOf":
		children, err := p.transformChildRules(matchRule["rules"].(map[string]interface{}))

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

func (p PayloadFactory) transformChildRules(childRules map[string]interface{}) ([]rules.Rule, error) {
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

func (p PayloadFactory) createResponse(data map[string]interface{}) *mock.Response {
	return mock.NewResponse(
		data["statusCode"].(int),
		data["headers"].(map[string]string),
		data["content"].(string),
	)
}
