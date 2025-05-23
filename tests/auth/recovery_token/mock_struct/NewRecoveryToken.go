package mock_struct

import "github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"

type NewRecoveryTokenRepository struct {
	Exists            bool
	RecoveryCodeToken string
	InsertOne         *model.RecoveryToken
}
type NewRecoveryTokenExpectedReturns struct {
	RepoError_1    error
	RepoError_2    error
	RepoError_3    error
	ExpectedReturn string
	ExpectedError  error
}
