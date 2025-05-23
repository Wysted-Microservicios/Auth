package mock_struct

import "github.com/CPU-commits/Template_Go-EventDriven/src/auth/dto"

type NewSessionParams struct {
	SessionDto dto.SessionDto
	AuthID     int64
	UserID     int64
}
type NewSessionRepository struct {
	NewSessionToken string
	InsertOne       int64
}
type NewSessionExpectedReturns struct {
	RepoError_1    error
	RepoError_2    error
	ExpectedReturn string
	ExpectedError  error
}
