package service

import (
	"github.com/gomatic/go-kit-phases/api/moody"
)

//
type Self struct {
	moody.UnimplementedSelfServer
}

//
func New() moody.SelfServer {
	return Self{}
}
