package recovery_codes_repository

import (
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"
)

type RecoveryCriteria struct {
	ID       int64
	Code     string
	IDUser   int64
	IsActive *bool
}

type RecoveryDataUpdate struct {
	IsActive *bool
}

type RecoveryRepository interface {
	InsertOne(recovery model.Recovery) (*model.Recovery, error)
	Exists(criteria *RecoveryCriteria) (bool, error)
	FindOne(criteria *RecoveryCriteria) (*model.Recovery, error)
	UpdateOne(id int64, data RecoveryDataUpdate) error
	Find(criteria *RecoveryCriteria) ([]model.Recovery, error)
}
