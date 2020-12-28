package server

import (
	"fmt"

	"github.com/gomatic/go-kit-phases/api/moody"
)

//
func FromInterface(request interface{}) (*moody.Query, error) {
	feeling, ok := request.(*moody.Query)
	if !ok {
		return nil, fmt.Errorf("got %T. expected %T", request, &moody.Query{})
	}
	return feeling, nil
}
