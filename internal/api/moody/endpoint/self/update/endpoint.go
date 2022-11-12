package update

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	api "github.com/gomatic/go-kit-phases/api/moody"
	"github.com/gomatic/go-kit-phases/pkg/api/middleware"
)

type Endpoint endpoint.Endpoint

func (c Endpoint) ToEndpoint() endpoint.Endpoint { return middleware.JWT(endpoint.Endpoint(c)) }

func New(s api.SelfServer) Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// TODO implement me
		return &api.Overall{
			Date:    2,
			Average: 3,
		}, nil
	}
}
