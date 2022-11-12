package self

import (
	"context"

	"github.com/gomatic/go-kit-phases/api/moody"
	service "github.com/gomatic/go-kit-phases/internal/api/moody/service/self"
)

//
func (s *serviceInstance) Create(ctx context.Context, request *moody.Feeling) (*moody.Overall, error) {
	return service.Self{}.Create(ctx, request)
}

//
func (s *serviceInstance) Retrieve(ctx context.Context, request *moody.Query) (*moody.Feeling, error) {
	return service.Self{}.Retrieve(ctx, request)
}

//
func (s *serviceInstance) Update(ctx context.Context, request *moody.Feeling) (*moody.Overall, error) {
	return service.Self{}.Update(ctx, request)
}

//
func (s *serviceInstance) Delete(ctx context.Context, request *moody.Feeling) (*moody.Feeling, error) {
	return service.Self{}.Delete(ctx, request)
}

//
func (s *serviceInstance) List(ctx context.Context, request *moody.Query) (*moody.Complicated, error) {
	return service.Self{}.List(ctx, request)
}
