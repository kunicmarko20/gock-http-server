package http

import "github.com/kunicmarko20/gock-http-server/mock"

type Payload struct {
	MatchRule map[string]interface{}
	Response  *mock.Response
}
