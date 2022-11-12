package caprequest

import (
	"encoding/json"
	"fmt"

	t "github.com/gomatic/go-kit-phases/api/moody"
	"github.com/gomatic/go-kit-phases/internal/api/moody/transform"
)

func To(p interface{}) (*t.Feeling, error) {
	r, ok := p.(*t.Feeling)
	if !ok {
		return nil, fmt.Errorf("got %T. expected %T", p, &t.Feeling{})
	}
	return r, nil
}

func FromJSON(q json.RawMessage) (interface{}, error) {
	var value t.Feeling
	return &value, json.Unmarshal(q, &value)
}

func ToJSON(p interface{}) (json.RawMessage, error) { return transform.Marshaller(To(p)) }
