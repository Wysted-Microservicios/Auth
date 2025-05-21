package utils

import (
	"errors"
	"time"
)

func VerifyNotExpiredAt(expiration time.Time, clockType string, err error) error {
	var now time.Time

	switch clockType {
	case "utc":
		now = time.Now().UTC()
	case "local":
		now = time.Now()
	default:
		return errors.New("invalid clock type: must be 'utc' or 'local'")
	}

	if now.After(expiration) {
		return err
	}

	return nil
}
