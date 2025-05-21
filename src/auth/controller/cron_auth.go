package controller

type cronAuthController struct{}

func NewCronAuth() *cronAuthController {
	return &cronAuthController{}
}

func (*cronAuthController) DeleteTokens() {
	sessionService.DeleteRevokedTokens()
}

func (*cronAuthController) DeleteExpiredSessions() {
	sessionService.DeleteExpiredSessions()
}
