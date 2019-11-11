package repository

import (
	"github.com/kunicmarko20/gock-http-server/mock"
)

type MockRepository interface {
	All() map[string]*mock.Mock
	Reset()
	Set(mock *mock.Mock)
	Get(mockName string) (*mock.Mock, error)
}
