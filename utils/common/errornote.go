package common

import "errors"

var (
	BadRequestError      = errors.New("bad request")
	NotFoundError        = errors.New("not found")
	InternalServerError = errors.New("internal server error")
	NotFoundErrorByID    = errors.New("not found by ID")
)
