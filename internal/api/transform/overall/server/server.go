package server

import (
	"fmt"

	"github.com/gomatic/go-kit-phases/api/moody"
)

//
func FromInterface(request interface{}) (*moody.Overall, error) {
	feeling, ok := request.(*moody.Overall)
	if !ok {
		return nil, fmt.Errorf("got %T. expected %T", request, &moody.Overall{})
	}
	return feeling, nil
}
