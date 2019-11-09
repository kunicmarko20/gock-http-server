package http

import "../mock"

type Payload struct {
	MatchRule map[string]interface{}
	Response  *mock.Response
}
