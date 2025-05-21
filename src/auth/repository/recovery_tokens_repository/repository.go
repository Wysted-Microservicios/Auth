package recovery_tokens_repository

import "github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"

type RecoveryTokenCriteria struct {
	Token  string
	ID     int64
	IsUsed *bool
}

type RecoveryTokenUpdate struct {
	IsUsed *bool
}

type RecoveryTokenRepository interface {
	InsertOne(recoveryToken model.RecoveryToken) (*model.RecoveryToken, error)
	FindOne(criteria *RecoveryTokenCriteria) (*model.RecoveryToken, error)
	Find(criteria *RecoveryTokenCriteria) ([]model.RecoveryToken, error)
	Exists(criteria *RecoveryTokenCriteria) (bool, error)
	UpdateOne(id int64, dataUpdate RecoveryTokenUpdate) error
}
