package server

import (
	"fmt"

	"github.com/gomatic/go-kit-phases/api/moody"
)

//
func FromInterface(request interface{}) (*moody.Feeling, error) {
	feeling, ok := request.(*moody.Feeling)
	if !ok {
		return nil, fmt.Errorf("got %T. expected %T", request, &moody.Feeling{})
	}
	return feeling, nil
}
