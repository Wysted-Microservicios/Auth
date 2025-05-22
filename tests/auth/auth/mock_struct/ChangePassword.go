package mock_struct

import "github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"

type ChangePasswordRepository struct {
	RecoveryTokenExists    bool
	RecoveryTokenFindOne   *model.RecoveryToken
	AuthUpdateOne          error
	RecoveryTokenUpdateOne error
}
type ChangePasswordExpectedReturns struct {
	RepoError_1   error
	RepoError_2   error
	RepoError_3   error
	RepoError_4   error
	ExpectedError error
}
