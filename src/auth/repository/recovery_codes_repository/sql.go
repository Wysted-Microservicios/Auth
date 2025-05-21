package recovery_codes_repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"
	"github.com/CPU-commits/Template_Go-EventDriven/src/package/db/models"
	"github.com/CPU-commits/Template_Go-EventDriven/src/utils"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type sqlRecoveryRepository struct {
	db *sql.DB
}

func (sqlRecoveryRepository) sqlRecoveryToRecovery(sqlRecovery *models.RecoveryCode) *model.Recovery {
	return &model.Recovery{
		ID:         sqlRecovery.ID,
		IDUser:     sqlRecovery.IDUser,
		Code:       sqlRecovery.Code,
		IsActive:   sqlRecovery.IsActive,
		Expires_at: sqlRecovery.ExpiresAt,
		CreatedAt:  sqlRecovery.CreatedAt,
	}
}
func (sqlRecoveryRepository) criteriaToWhere(criteria *RecoveryCriteria) []QueryMod {
	var mod []QueryMod
	if criteria == nil {
		return nil
	}
	if criteria.Code != "" {
		mod = append(mod, Where("code = ?", criteria.Code))
	}
	if criteria.IDUser != 0 {
		mod = append(mod, Where("id_user = ?", criteria.IDUser))
	}
	if criteria.IsActive != nil {
		mod = append(mod, Where("is_active = ?", criteria.IsActive))
	}

	return mod
}

func (sqlRCY sqlRecoveryRepository) InsertOne(recovery model.Recovery) (*model.Recovery, error) {
	utcNow := time.Now().UTC()
	expiresAt := utcNow.Add(2 * time.Minute)
	recoverySQL := models.RecoveryCode{
		Code:      recovery.Code,
		IDUser:    recovery.IDUser,
		ExpiresAt: expiresAt,
	}

	err := recoverySQL.Insert(context.Background(), sqlRCY.db, boil.Infer())
	if err != nil {
		return nil, utils.ErrRepositoryFailed
	}
	return sqlRCY.sqlRecoveryToRecovery(&recoverySQL), nil
}

func (sqlRCY sqlRecoveryRepository) Exists(criteria *RecoveryCriteria) (bool, error) {
	where := sqlRCY.criteriaToWhere(criteria)

	exists, err := models.RecoveryCodes(where...).Exists(context.Background(), sqlRCY.db)
	if err != nil {
		return false, utils.ErrRepositoryFailed
	}

	return exists, nil
}

func (sqlRCY sqlRecoveryRepository) FindOne(criteria *RecoveryCriteria) (*model.Recovery, error) {
	where := sqlRCY.criteriaToWhere(criteria)

	recoverySQL, err := models.RecoveryCodes(where...).One(context.Background(), sqlRCY.db)
	if err != nil {
		return nil, utils.ErrRepositoryFailed
	}

	return sqlRCY.sqlRecoveryToRecovery(recoverySQL), nil
}

func (sqlRCY sqlRecoveryRepository) UpdateOne(id int64, data RecoveryDataUpdate) error {

	recoveryCode, err := models.FindRecoveryCode(context.Background(), sqlRCY.db, id)
	if err != nil {
		return utils.ErrNotFoundRow
	}

	if !recoveryCode.IsActive {
		return nil
	}
	recoveryCode.IsActive = *data.IsActive

	_, err = recoveryCode.Update(context.Background(), sqlRCY.db, boil.Infer())
	if err != nil {
		return utils.ErrRepositoryFailed
	}

	return nil
}

func (sqlRCY sqlRecoveryRepository) Find(criteria *RecoveryCriteria) ([]model.Recovery, error) {
	where := sqlRCY.criteriaToWhere(criteria)

	recoveryCodeSQL, err := models.RecoveryCodes(where...).All(context.Background(), sqlRCY.db)
	if err != nil {
		return nil, utils.ErrRepositoryFailed
	}

	return utils.MapNoError(recoveryCodeSQL, func(recoveryCode *models.RecoveryCode) model.Recovery {
		return *sqlRCY.sqlRecoveryToRecovery(recoveryCode)
	}), nil
}
func NewSQLRecoveryRepository(db *sql.DB) RecoveryRepository {
	return sqlRecoveryRepository{
		db: db,
	}
}
