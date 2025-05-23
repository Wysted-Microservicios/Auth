package mock_struct

type RefreshSessionParams struct {
	SessionToken string
	UserID       int64
}
type RefreshSessionRepository struct {
	NewSessionToken string
}
type RefreshSessionExpectedReturns struct {
	RepoError_1    error
	RepoError_2    error
	ExpectedReturn string
	ExpectedError  error
}
