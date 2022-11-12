// Copyright © 2021 nicerobot
//go:build linux || darwin || freebsd
// +build linux darwin freebsd

package instance

import (
	"github.com/gomatic/go-kit-phases/pkg/errors"
)

const (
	ErrorNoListener errors.Sentinel = "no listener"
)
