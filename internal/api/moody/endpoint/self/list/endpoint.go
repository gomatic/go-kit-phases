package list

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
		return &api.Complicated{
			Feeling: []*api.Feeling{
				{
					Enjoyment: &api.Enjoyment{
						Feeling: 1,
						Level:   2,
					},
					Anger: &api.Anger{
						Feeling: 3,
						Level:   4,
					},
					Fear: &api.Fear{
						Feeling: 5,
						Level:   6,
					},
					Sadness: &api.Sadness{
						Feeling: 7,
						Level:   8,
					},
					Disgust: &api.Disgust{
						Feeling: 9,
						Level:   10,
					},
					Date: 0,
				},
			},
		}, nil
	}
}
