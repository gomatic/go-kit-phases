package overall

import (
	"encoding/json"
	"fmt"

	t "github.com/gomatic/go-kit-phases/api/moody"
	"github.com/gomatic/go-kit-phases/internal/api/moody/transform"
)

func To(p interface{}) (*t.Overall, error) {
	r, ok := p.(*t.Overall)
	if !ok {
		return nil, fmt.Errorf("got %T. expected %T", p, &t.Overall{})
	}
	return r, nil
}

func FromJSON(q json.RawMessage) (interface{}, error) {
	var value t.Overall
	return &value, json.Unmarshal(q, &value)
}

func ToJSON(p interface{}) (json.RawMessage, error) { return transform.Marshaller(To(p)) }
