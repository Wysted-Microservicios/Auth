package auth_repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"
	"github.com/CPU-commits/Template_Go-EventDriven/src/package/db/models"
	"github.com/CPU-commits/Template_Go-EventDriven/src/utils"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

type sqlAuthRepository struct {
	db *sql.DB
}

func (sqlAuthRepository) sqlAuthToAuth(auth *models.Auth) *model.Auth {
	return &model.Auth{
		ID:       auth.ID,
		Password: auth.Password,
		IDUser:   auth.IDUser,
	}
}

func (sqlAR sqlAuthRepository) FindOneByUsername(username string) (*model.Auth, error) {
	auth, err := models.Auths(
		InnerJoin("users u on u.id = auths.id_user"),
		Where("u.email = ?", username),
	).One(context.Background(), sqlAR.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, utils.ErrRepositoryFailed
	}

	return sqlAR.sqlAuthToAuth(auth), nil
}

func (sqlAR sqlAuthRepository) UpdateOne(id int64, dataUpdate *AuthDataUpdate) error {
	auth, err := models.FindAuth(context.Background(), sqlAR.db, id)
	if err != nil {
		return utils.ErrNotFoundRow
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(dataUpdate.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	auth.Password = string(passwordHashed)

	_, err = auth.Update(context.Background(), sqlAR.db, boil.Infer())
	if err != nil {
		return utils.ErrRepositoryFailed
	}
	return nil
}

func NewSQLAuthRepository(db *sql.DB) AuthRepository {
	return sqlAuthRepository{
		db: db,
	}
}
