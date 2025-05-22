package mock_struct

import "github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"

type RecoveryCodeExpiryRepository struct {
	Find []model.Recovery
}
type RecoveryCodeExpiryExpectedReturns struct {
	RepoError_1   error
	RepoError_2   []error
	ExpectedError error
}
