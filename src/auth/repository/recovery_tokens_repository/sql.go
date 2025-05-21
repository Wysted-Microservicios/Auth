package recovery_tokens_repository

import (
	"context"
	"database/sql"

	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"
	"github.com/CPU-commits/Template_Go-EventDriven/src/package/db/models"
	"github.com/CPU-commits/Template_Go-EventDriven/src/utils"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type sqlRecoveryTokenRepository struct {
	db *sql.DB
}

func (sqlRecoveryTokenRepository) sqlRecoveryTokenToRecoveryToken(sqlRecoveryToken *models.RecoveryToken) *model.RecoveryToken {
	return &model.RecoveryToken{
		ID:         sqlRecoveryToken.ID,
		IDUser:     sqlRecoveryToken.IDUser,
		Token:      sqlRecoveryToken.Token,
		IsUsed:     sqlRecoveryToken.IsUsed,
		Expires_at: sqlRecoveryToken.ExpiresAt,
		CreatedAt:  sqlRecoveryToken.CreatedAt,
	}
}

func (sqlRecoveryTokenRepository) criteriaToWhere(criteria *RecoveryTokenCriteria) []QueryMod {
	var mod []QueryMod
	if criteria == nil {
		return nil
	}
	if criteria.Token != "" {
		mod = append(mod, Where("token = ?", criteria.Token))
	}
	if criteria.ID != 0 {
		mod = append(mod, Where("id = ?", criteria.ID))
	}
	if criteria.IsUsed != nil {
		mod = append(mod, Where("is_used = ?", criteria.IsUsed))
	}
	return mod
}

func (sqlRTR sqlRecoveryTokenRepository) InsertOne(recoveryToken model.RecoveryToken) (*model.RecoveryToken, error) {

	recoveryTokenSQL := models.RecoveryToken{
		Token:     recoveryToken.Token,
		IDUser:    recoveryToken.IDUser,
		ExpiresAt: recoveryToken.Expires_at,
	}
	err := recoveryTokenSQL.Insert(context.Background(), sqlRTR.db, boil.Infer())
	if err != nil {
		return nil, utils.ErrRepositoryFailed
	}
	return sqlRTR.sqlRecoveryTokenToRecoveryToken(&recoveryTokenSQL), nil
}

func (sqlRTR sqlRecoveryTokenRepository) FindOne(criteria *RecoveryTokenCriteria) (*model.RecoveryToken, error) {
	where := sqlRTR.criteriaToWhere(criteria)

	recoveryToken, err := models.RecoveryTokens(where...).One(context.Background(), sqlRTR.db)
	if err != nil {
		return nil, utils.ErrRepositoryFailed
	}

	return sqlRTR.sqlRecoveryTokenToRecoveryToken(recoveryToken), nil
}
func (sqlRTR sqlRecoveryTokenRepository) Exists(criteria *RecoveryTokenCriteria) (bool, error) {
	where := sqlRTR.criteriaToWhere(criteria)

	exists, err := models.RecoveryTokens(where...).Exists(context.Background(), sqlRTR.db)
	if err != nil {
		return false, utils.ErrRepositoryFailed
	}

	return exists, nil
}

func (sqlRTR sqlRecoveryTokenRepository) UpdateOne(id int64, dataUpdate RecoveryTokenUpdate) error {
	recoveryToken, err := models.FindRecoveryToken(context.Background(), sqlRTR.db, id)
	if err != nil {
		return utils.ErrNotFoundRow
	}
	recoveryToken.IsUsed = *dataUpdate.IsUsed

	_, err = recoveryToken.Update(context.Background(), sqlRTR.db, boil.Infer())
	if err != nil {
		return utils.ErrRepositoryFailed
	}
	return nil
}
func (sqlRTR sqlRecoveryTokenRepository) Find(criteria *RecoveryTokenCriteria) ([]model.RecoveryToken, error) {
	where := sqlRTR.criteriaToWhere(criteria)

	recoveryTokenSQL, err := models.RecoveryTokens(where...).All(context.Background(), sqlRTR.db)
	if err != nil {
		return nil, utils.ErrRepositoryFailed
	}

	return utils.MapNoError(recoveryTokenSQL, func(recoveryToken *models.RecoveryToken) model.RecoveryToken {
		return *sqlRTR.sqlRecoveryTokenToRecoveryToken(recoveryToken)
	}), nil
}

func NewSQLRecoveryTokenRepository(db *sql.DB) sqlRecoveryTokenRepository {
	return sqlRecoveryTokenRepository{
		db: db,
	}
}
