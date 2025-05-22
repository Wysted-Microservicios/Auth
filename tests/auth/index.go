package auth

import (
	"os"
	"testing"

	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/service"
	"github.com/CPU-commits/Template_Go-EventDriven/tests/mocks"
)

var UserS *service.UserService

var mockUserRepo *mocks.MockUserRepository
var mockTokenGeneratorRepo *mocks.MockTokenGenerator
var mockSessionRepo *mocks.MockSessionRepository
var mockRecoveryTokenRepo *mocks.MockRecoveryTokenRepository
var mockRecoveryRepo *mocks.MockRecoveryRepository
var mockLoggerRepo *mocks.MockLogger
var mockBusRepo *mocks.MockBus
var mockAuthRepo *mocks.MockAuthRepository
var mockAccessRepo *mocks.MockAccessRepository

func TestMain(m *testing.M) {
	mockUserRepo = &mocks.MockUserRepository{}
	mockTokenGeneratorRepo = &mocks.MockTokenGenerator{}
	mockSessionRepo = &mocks.MockSessionRepository{}
	mockRecoveryTokenRepo = &mocks.MockRecoveryTokenRepository{}
	mockRecoveryRepo = &mocks.MockRecoveryRepository{}
	mockLoggerRepo = &mocks.MockLogger{}
	mockBusRepo = &mocks.MockBus{}
	mockAuthRepo = &mocks.MockAuthRepository{}
	mockAccessRepo = &mocks.MockAccessRepository{}

	UserS = service.NewUserService(mockUserRepo)

	code := m.Run()
	os.Exit(code)
}
