package auth_repository

import "github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"

type AuthDataUpdate struct {
	Password string
}

type AuthRepository interface {
	FindOneByUsername(username string) (*model.Auth, error)
	UpdateOne(id int64, dataUpdate *AuthDataUpdate) error
}
