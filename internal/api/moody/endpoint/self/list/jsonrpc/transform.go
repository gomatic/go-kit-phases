package jsonrpc

import (
	"context"
	"encoding/json"

	out "github.com/gomatic/go-kit-phases/internal/api/moody/transform/complicated"
	in "github.com/gomatic/go-kit-phases/internal/api/moody/transform/query"
)

func Request(_ context.Context, q json.RawMessage) (interface{}, error)  { return in.FromJSON(q) }
func Response(_ context.Context, p interface{}) (json.RawMessage, error) { return out.ToJSON(p) }
