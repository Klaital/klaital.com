package repositories

import "errors"

var ErrNoData = errors.New("no data")
var ErrNotValid error = errors.New("input data not valid")
var ErrNotImplemented = errors.New("operation not implemented")
