package service

import (
	"context"

	"github.com/gomatic/go-kit-phases/api/moody"
)

//
type Self struct {
	moody.UnimplementedSelfServer
}

//
func New() moody.SelfServer {
	var svc moody.SelfServer
	svc = NewSelfService()
	return svc
}

//
func NewSelfService() moody.SelfServer {
	return Self{}
}

//
func (s Self) Create(_ context.Context, _ *moody.Feeling) (*moody.Overall, error) {
	panic("implement me")
}

//
func (s Self) Retrieve(_ context.Context, _ *moody.Query) (*moody.Feeling, error) {
	panic("implement me")
}

//
func (s Self) Update(_ context.Context, _ *moody.Feeling) (*moody.Overall, error) {
	panic("implement me")
}

//
func (s Self) Delete(_ context.Context, _ *moody.Feeling) (*moody.Feeling, error) {
	panic("implement me")
}

//
func (s Self) List(_ context.Context, _ *moody.Query) (*moody.Feelings, error) {
	panic("implement me")
}
