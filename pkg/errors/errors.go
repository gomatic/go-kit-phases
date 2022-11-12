// Copyright © 2021 nicerobot
//go:build linux || darwin || freebsd
// +build linux darwin freebsd

package errors

import (
	"context"
	"encoding/json"
	"net/http"
)

type Sentinel string

func (e Sentinel) Error() string { return string(e) }

//
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(ErrorWrapper{Error: err.Error()})
}

//
type ErrorWrapper struct {
	Error string `json:"error"`
}
