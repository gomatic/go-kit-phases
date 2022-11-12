package jsonrpc

import (
	"context"
	"encoding/json"

	in "github.com/gomatic/go-kit-phases/internal/api/moody/transform/feeling"
	out "github.com/gomatic/go-kit-phases/internal/api/moody/transform/overall"
)

func Request(_ context.Context, q json.RawMessage) (interface{}, error)  { return in.FromJSON(q) }
func Response(_ context.Context, p interface{}) (json.RawMessage, error) { return out.ToJSON(p) }
