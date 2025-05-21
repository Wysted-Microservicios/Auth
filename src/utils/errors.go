package utils

import "errors"

var ErrRepositoryFailed = errors.New("err: repository failed")
var ErrNotFoundRow = errors.New("err: not found row")
var ErrTokenNotValid = errors.New("err: token not valid")
