package mock_struct

import "github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"

type CheckTokenRepository struct {
	Exists  bool
	FindOne *model.RecoveryToken
}
type CheckTokenExpectedReturns struct {
	RepoError_1    error
	RepoError_2    error
	ExpectedReturn model.RecoveryToken
	ExpectedError  error
}
