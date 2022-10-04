package common

import "errors"

var (
	ErrInvalidParameter = errors.New("invalid parameter")
	ErrExit             = errors.New("exit")
)
