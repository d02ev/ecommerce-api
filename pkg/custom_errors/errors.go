package custom_errors

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
)
