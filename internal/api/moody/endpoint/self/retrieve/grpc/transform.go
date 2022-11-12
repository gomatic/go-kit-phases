package grpc

import (
	"context"

	out "github.com/gomatic/go-kit-phases/internal/api/moody/transform/feeling"
	in "github.com/gomatic/go-kit-phases/internal/api/moody/transform/query"
)

func Request(_ context.Context, q interface{}) (interface{}, error)  { return in.To(q) }
func Response(_ context.Context, p interface{}) (interface{}, error) { return out.To(p) }
