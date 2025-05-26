package tests

import (
	"time"

	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/dto"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"
	"golang.org/x/crypto/bcrypt"
)

// AnythingOfType
var (
	USER_CRITERIA_PTR        string = "*user_repository.Criteria"
	USER_FIND_ONE_OPTIOS_PTR string = "*user_repository.FindOneOptions"

	RECOVERY_TOKEN_CRITERIA_PTR string = "*recovery_tokens_repository.RecoveryTokenCriteria"
	RECOVERY_TOKEN_DATA_UPDATE  string = "recovery_tokens_repository.RecoveryTokenUpdate"

	RECOVERY_CRITERIA_PTR string = "*recovery_codes_repository.RecoveryCriteria"
	RECOVERY_DATA_UPDATE  string = "recovery_codes_repository.RecoveryDataUpdate"

	AUTH_DATA_UPDATE string = "*auth_repository.AuthDataUpdate"

	SESSION_CRITERIA_PTR string = "*session_repository.SessionCriteria"
	SESSION_DATA_UPDATE  string = "session_repository.SessionUpdateData"

	ACCESS_CRITERIA_PTR string = "*access_repository.AccessCriteria"
	ACCESS_DATA_UPDATE  string = "access_repository.AccessUpdateData"

	USER_MODEL_PTR       string = "*model.User"
	USER_MODEL           string = "model.User"
	RECOVERY_MODEL       string = "model.Recovery"
	RECOVERY_MODEL_PTR   string = "*model.Recovery"
	RECOVERY_TOKEN_MODEL string = "model.RecoveryToken"
	SESSION_MODEL        string = "model.Session"
	ACCESS_MODEL         string = "model.Access"

	BUS_EVENT string = "bus.Event"

	STRING_ARRAY string = "[]string"
	STRING       string = "string"
	STRING_PTR   string = "*string"
	INT64        string = "int64"
	TIME         string = "time.Time"
)

// Mock.On
var (
	EXISTS                  string = "Exists"
	FIND                    string = "Find"
	FIND_ONE                string = "FindOne"
	FIND_ONE_BY_USERNAME    string = "FindOneByUsername"
	FIND_ONE_BY_EMAIL       string = "FindOneByEmail"
	FIND_ONE_BY_ID          string = "FindOneByID"
	INSERT                  string = "Insert"
	INSERT_ONE              string = "InsertOne"
	UPDATE                  string = "Update"
	UPDATE_ONE              string = "UpdateOne"
	DELETE                  string = "Delete"
	CHECK_TOKEN             string = "CheckToken"
	NEW_RECOVERY_CODE_TOKEN string = "NewRecoveryCodeToken"
	NEW_SESSION_TOKEN       string = "NewSessionToken"
	NEW_ACCESS_TOKEN        string = "NewAccessToken"
	PUBLISH                 string = "Publish"
)

// Datos ficticios para modelos y dto's
var (
	// ID's
	ID_1 int64 = 1

	//	Names
	Name_1 string = "Wysted"

	// Usernames
	Username_1 string = "WystedTest"

	// Emails
	Email_1 string = "wystedtest@test.cl"

	// Roles
	Role_1 string = "user"

	// Passwords
	Password_1 string = "WystedPassword"
	Password_2 string = "OtherPassword"

	// Tokens
	Token_1 string = "token"

	// Bools
	Bool_1 bool = true
	Bool_2 bool = false

	//Time
	Time_1 time.Time = time.Now().UTC()
	Time_2 time.Time = time.Now().UTC().Add(time.Minute * -5)
	Time_3 time.Time = time.Now().UTC().Add(time.Minute * 5)

	// Code
	Code_1 string = "SFHBDA"
)

func hashed(p string) string {
	h, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return string(h)
}

// Models - DTO's | Para casos de uso
var (
	// Models
	User_1 *model.User = &model.User{
		ID:       ID_1,
		Email:    Email_1,
		Name:     Name_1,
		Username: Username_1,
		Roles:    []model.Role{model.Role(Role_1)},
	}
	Auth_1 *model.Auth = &model.Auth{
		ID:       ID_1,
		Password: hashed(Password_1),
		IDUser:   ID_1,
	}
	Auth_2 *model.Auth = &model.Auth{
		ID:       ID_1,
		Password: hashed(Password_2),
		IDUser:   ID_1,
	}

	RecoveryToken_1 *model.RecoveryToken = &model.RecoveryToken{
		ID:         ID_1,
		Token:      Token_1,
		IDUser:     ID_1,
		IsUsed:     Bool_1,
		Expires_at: Time_3,
		CreatedAt:  Time_1,
	}
	RecoveryToken_2 *model.RecoveryToken = &model.RecoveryToken{
		ID:         ID_1,
		Token:      Token_1,
		IDUser:     ID_1,
		IsUsed:     Bool_1,
		Expires_at: Time_2,
		CreatedAt:  Time_1,
	}

	RecoveryCode_1 *model.Recovery = &model.Recovery{
		ID:         ID_1,
		Code:       Code_1,
		IDUser:     ID_1,
		IsActive:   Bool_2,
		Expires_at: Time_3,
		CreatedAt:  Time_1,
	}
	RecoveryCode_2 *model.Recovery = &model.Recovery{
		ID:         ID_1,
		Code:       Code_1,
		IDUser:     ID_1,
		IsActive:   Bool_2,
		Expires_at: Time_2,
		CreatedAt:  Time_1,
	}

	Session_1 model.Session = model.Session{
		Token:     Token_1,
		IDAuth:    ID_1,
		ExpiresAt: Time_3,
	}

	// Dto's
	RegisterDto_1 *dto.RegisterDto = &dto.RegisterDto{
		Name:     Name_1,
		Username: Username_1,
		Email:    Email_1,
		Password: Password_1,
		Role:     Role_1,
	}
	AuthDto_1 *dto.AuthDto = &dto.AuthDto{
		Username: Name_1,
		Password: Password_1,
	}
	ChangePasswordDto_1 dto.ChangePassword = dto.ChangePassword{
		Token:           Token_1,
		Password:        Password_1,
		ConfirmPassword: Password_1,
	}

	ChangePasswordDto_2 dto.ChangePassword = dto.ChangePassword{
		Token:           Token_1,
		Password:        Password_1,
		ConfirmPassword: Password_2,
	}

	VerifyRecoveryCode_1 *dto.VerifyRecoveryCode = &dto.VerifyRecoveryCode{
		Email: Email_1,
		Code:  Code_1,
	}

	SessionDto dto.SessionDto = dto.SessionDto{}
)
