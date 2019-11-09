package repository

import (
	"../../mock"
)

type MockRepository interface {
	All() map[string]*mock.Mock
	Reset()
	Set(mock *mock.Mock)
	Get(mockName string) (*mock.Mock, error)
}
