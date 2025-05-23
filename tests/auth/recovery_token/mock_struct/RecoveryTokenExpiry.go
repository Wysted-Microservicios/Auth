package mock_struct

import "github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"

type RecoveryTokenExpiryRepository struct {
	Find []model.RecoveryToken
}
type RecoveryTokenExpiryExpectedReturns struct {
	RepoError_1   error
	RepoError_2   []error
	ExpectedError error
}
