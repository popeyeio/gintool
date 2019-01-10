package instance

import (
	"errors"
)

var (
	ErrNoInstance = errors.New("no instance available")
)

type Instance interface {
	GetValue() interface{}
	GetWeight() int
	GetIDC() string
	GetCluster() string
}
