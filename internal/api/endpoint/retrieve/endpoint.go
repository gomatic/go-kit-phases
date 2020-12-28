package retrieve

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
		return &moody.Feeling{
			Enjoyment: &moody.Enjoyment{
				Feeling: 0,
				Level:   0,
			},
			Anger: &moody.Anger{
				Feeling: 0,
				Level:   0,
			},
			Fear: &moody.Fear{
				Feeling: 0,
				Level:   0,
			},
			Sadness: &moody.Sadness{
				Feeling: 0,
				Level:   0,
			},
			Disgust: &moody.Disgust{
				Feeling: 0,
				Level:   0,
			},
			Date: 0,
		}, nil
	}
}
