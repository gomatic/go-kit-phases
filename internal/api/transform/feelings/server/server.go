package server

import (
	"fmt"

	"github.com/gomatic/go-kit-phases/api/moody"
)

//
func FromInterface(request interface{}) (*moody.Feelings, error) {
	feeling, ok := request.(*moody.Feelings)
	if !ok {
		return nil, fmt.Errorf("got %T. expected %T", request, &moody.Feelings{})
	}
	return feeling, nil
}
