package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/gomatic/go-kit-phases/api/moody"
)

//
type Set struct {
	CreateEndpoint   endpoint.Endpoint
	RetrieveEndpoint endpoint.Endpoint
	UpdateEndpoint   endpoint.Endpoint
	DeleteEndpoint   endpoint.Endpoint
	ListEndpoint     endpoint.Endpoint
}

//
func New(svc moody.SelfServer) Set {
	return Set{
		CreateEndpoint:   NewCreateEndpoint(svc),
		RetrieveEndpoint: NewRetrieveEndpoint(svc),
		UpdateEndpoint:   NewUpdateEndpoint(svc),
		DeleteEndpoint:   NewDeleteEndpoint(svc),
		ListEndpoint:     NewListEndpoint(svc),
	}
}

//
func NewCreateEndpoint(s moody.SelfServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return &moody.Overall{
			Date:    0,
			Average: 0,
		}, nil
	}
}

//
func NewRetrieveEndpoint(s moody.SelfServer) endpoint.Endpoint {
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

//
func NewUpdateEndpoint(s moody.SelfServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return &moody.Overall{
			Date:    0,
			Average: 0,
		}, nil
	}
}

//
func NewDeleteEndpoint(s moody.SelfServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return &moody.Overall{
			Date:    0,
			Average: 0,
		}, nil
	}
}

//
func NewListEndpoint(s moody.SelfServer) endpoint.Endpoint {
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
