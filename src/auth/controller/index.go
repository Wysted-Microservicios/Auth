package controller

import (
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/repository/access_repository"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/repository/auth_repository"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/repository/recovery_codes_repository"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/repository/recovery_tokens_repository"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/repository/session_repository"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/repository/user_repository"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/service"
	"github.com/CPU-commits/Template_Go-EventDriven/src/cmd/http/utils"
	"github.com/CPU-commits/Template_Go-EventDriven/src/package/bus"
	"github.com/CPU-commits/Template_Go-EventDriven/src/package/db"
)

// Generator
var generator = utils.NewGeneratorToken()

// Repositories
var (
	sqlAuthRepository          = auth_repository.NewSQLAuthRepository(db.DB)
	sqlSessionRepository       = session_repository.NewSQLSessionRepository(db.DB)
	sqlUserRepository          = user_repository.NewSQLUserRepository(db.DB)
	sqlAccessRepository        = access_repository.NewSQLAccessRepository(db.DB)
	sqlRecoveryRepository      = recovery_codes_repository.NewSQLRecoveryRepository(db.DB)
	sqlRecoveryTokenRepository = recovery_tokens_repository.NewSQLRecoveryTokenRepository(db.DB)
)

// Events
const (
	INSERTED_USER bus.EventName = "user.token_created"
)

// Services
var (
	recoveryTokenService = service.NewRecoveryTokenService(
		sqlRecoveryTokenRepository,
		generator,
		sqlUserRepository,
	)
	recoveryService = service.NewRecoveryService(
		sqlRecoveryRepository,
		sqlUserRepository,
		*recoveryTokenService,
	)
	authService = service.NewAuthService(
		sqlAuthRepository,
		sqlUserRepository,
		*recoveryTokenService,
		sqlRecoveryTokenRepository,
	)
	sessionService = service.NewSessionService(
		sqlSessionRepository,
		sqlAccessRepository,
		generator,
	)
	userService = service.NewUserService(
		sqlUserRepository,
	)
)
