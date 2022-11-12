package complicated

import (
	"encoding/json"
	"fmt"

	t "github.com/gomatic/go-kit-phases/api/moody"
	"github.com/gomatic/go-kit-phases/internal/api/moody/transform"
)

func To(q interface{}) (*t.Complicated, error) {
	r, ok := q.(*t.Complicated)
	if !ok {
		return nil, fmt.Errorf("got %T. expected %T", q, &t.Complicated{})
	}
	return r, nil
}

func FromJSON(q json.RawMessage) (interface{}, error) {
	var value t.Complicated
	return &value, json.Unmarshal(q, &value)
}

func ToJSON(p interface{}) (json.RawMessage, error) { return transform.Marshaller(To(p)) }
