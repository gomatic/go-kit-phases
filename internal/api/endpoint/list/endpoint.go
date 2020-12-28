package list

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
		return &moody.Feelings{
			Feeling: []*moody.Feeling{
				{
					Enjoyment: &moody.Enjoyment{
						Feeling: 1,
						Level:   1,
					},
					Anger: &moody.Anger{
						Feeling: 1,
						Level:   1,
					},
					Fear: &moody.Fear{
						Feeling: 1,
						Level:   1,
					},
					Sadness: &moody.Sadness{
						Feeling: 1,
						Level:   1,
					},
					Disgust: &moody.Disgust{
						Feeling: 1,
						Level:   1,
					},
					Date: 1,
				},
				{
					Enjoyment: &moody.Enjoyment{
						Feeling: 2,
						Level:   2,
					},
					Anger: &moody.Anger{
						Feeling: 2,
						Level:   2,
					},
					Fear: &moody.Fear{
						Feeling: 2,
						Level:   2,
					},
					Sadness: &moody.Sadness{
						Feeling: 2,
						Level:   2,
					},
					Disgust: &moody.Disgust{
						Feeling: 2,
						Level:   2,
					},
					Date: 2,
				},
			},
		}, nil
	}
}
