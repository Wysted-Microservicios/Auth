package token_generator_repository

import (
	"time"

	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"
)

type TokenGenerator interface {
	NewSessionToken(expiredAt time.Time, idUser int64) (string, error)
	NewAccessToken(expiredAt time.Time, user model.User) (string, error)
	NewFirstTimeToken(IDUser int64) (string, error)
	NewRecoveryCodeToken(expireddAt time.Time, user model.User) (string, error)
}
