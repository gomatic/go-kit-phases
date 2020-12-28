package create

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/gomatic/go-kit-phases/api/moody"
)

//
type Endpoint endpoint.Endpoint

//
func (c Endpoint) ToEndpoint() endpoint.Endpoint {
	return endpoint.Endpoint(c)
}

//
func New(s moody.SelfServer) Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return &moody.Overall{
			Date:    0,
			Average: 0,
		}, nil
	}
}
