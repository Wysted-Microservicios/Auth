package service

import "errors"

var ErrUserLoginNotFound = errors.New("err: user login not found")
var ErrInvalidCredentials = errors.New("err: invalid credentials")
var ErrSessionNotExists = errors.New("err: session not exists")
var ErrUserNotFound = errors.New("err: user not found")
var ErrTokenRevokedOrNotExists = errors.New("err: token revoked")
var ErrExistsEmailOrUsername = errors.New("err: exists email or username")
var ErrUsernameNotExists = errors.New("err: username not exists")
var ErrCodeNotValid = errors.New("err: code not valid")
var ErrTokenNotValid = errors.New("err: token not valid")
