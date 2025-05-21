package controller

import (
	"net/http"
	"strings"

	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/dto"
	"github.com/CPU-commits/Template_Go-EventDriven/src/cmd/http/utils"
	"github.com/gin-gonic/gin"
)

type HttpAuthController struct{}

// Register godoc
//
//	@Summary	Registrase
//	@Tags		auth
//	@Success	200		{object}	controller.LoginResponse
//	@Param		authDto	body		dto.RegisterDto			true	"name, username, email, password, role"
//	@Failure	503		{object}	utils.ProblemDetails	"Error con la base de datos"
//
//	@Failure	403		{object}	utils.ProblemDetails	"Credenciales inválidas"
//	@Failure	409		{object}	utils.ProblemDetails	"La sesión no existe. Probablemente porque la eliminaron"
//
//	@Router		/api/auth/register [post]
func (*HttpAuthController) Register(c *gin.Context) {
	var registerDto *dto.RegisterDto

	if err := c.BindJSON(&registerDto); err != nil {
		utils.ResErrValidators(c, err)
		return
	}
	if err := authService.Register(registerDto); err != nil {
		utils.ResFromErr(c, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// Login godoc
//
//	@Summary	Loggearse
//	@Tags		auth
//	@Success	200		{object}	controller.LoginResponse
//	@Param		authDto	body		dto.AuthDto				true	"Password y username"
//	@Failure	503		{object}	utils.ProblemDetails	"Error con la base de datos"
//
//	@Failure	403		{object}	utils.ProblemDetails	"Credenciales inválidas"
//	@Failure	409		{object}	utils.ProblemDetails	"La sesión no existe. Probablemente porque la eliminaron"
//
//	@Router		/api/auth/login [post]
func (*HttpAuthController) Login(c *gin.Context) {
	var authDto *dto.AuthDto

	if err := c.BindJSON(&authDto); err != nil {
		utils.ResErrValidators(c, err)
		return
	}
	user, idAuth, err := authService.Login(*authDto)
	if err != nil {
		utils.ResFromErr(c, err)
		return
	}
	// Generate token and session
	sessionDto := dto.SessionDto{}
	tokenSession, err := sessionService.NewSession(sessionDto, idAuth, user.ID)
	if err != nil {
		utils.ResFromErr(c, err)
		return
	}
	tokenAccess, err := sessionService.GenerateAccess(tokenSession, user)
	if err != nil {
		utils.ResFromErr(c, err)
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		AccessToken:  tokenAccess,
		RefreshToken: tokenSession,
		User:         user,
	})
}

// Refresh godoc
//
//	@Summary	Refrescar sesión
//	@Tags		auth
//	@Success	200			{object}	controller.LoginResponse
//	@Param		X-Refresh	header		string					true	"Token de refresco, es decir, de sesión"
//	@Failure	403			{object}	utils.ProblemDetails	"No está el token de refresco en el header X-Refresh"
//	@Failure	400			{object}	utils.ProblemDetails	"No es un token válido JWT"
//	@Failure	404			{object}	utils.ProblemDetails	"El token no tiene un usuario registrado en la BD"
//	@Failure	409			{object}	utils.ProblemDetails	"La sesión no existe. Probablemente porque la eliminaron"
//	@Router		/api/auth/refresh [post]
func (*HttpAuthController) Refresh(c *gin.Context) {
	refreshToken := c.GetHeader("X-Refresh")
	if refreshToken == "" {
		utils.ResWithMessageID(c, "refresh_not_found", http.StatusForbidden)
		return
	}
	// Manage token
	token, err := utils.VerifyToken(refreshToken)
	if err != nil {
		utils.ResWithMessageID(c, "unauthorized", http.StatusBadRequest)
		return
	}
	metadata, err := utils.ExtractRefreshTokeMetadata(token)
	if err != nil {
		utils.ResWithMessageID(c, "unauthorized", http.StatusBadRequest)
		return
	}
	// Refresh session with user
	user, err := userService.GetUserById(int64(metadata.UID))
	if err != nil {
		utils.ResFromErr(c, err)
		return
	}
	oldSessionToken := strings.ReplaceAll(refreshToken, "Bearer ", "")

	tokenSession, err := sessionService.RefreshSession(
		oldSessionToken,
		user.ID,
	)
	if err != nil {
		utils.ResFromErr(c, err)
		return
	}
	tokenAccess, err := sessionService.GenerateAccess(
		tokenSession,
		user,
	)
	if err != nil {
		utils.ResFromErr(c, err)
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		AccessToken:  tokenAccess,
		RefreshToken: tokenSession,
		User:         user,
	})
}

// ChangePassword godoc
//
//	@Summary	Cambiar la contraseña
//	@Tags		auth
//	@Success	200		{object}	controller.LoginResponse
//	@Param		authDto	body		dto.AuthDto				true	"Password y username"
//	@Failure	503		{object}	utils.ProblemDetails	"Error con la base de datos"
//
//	@Failure	403		{object}	utils.ProblemDetails	"Credenciales inválidas"
//	@Failure	409		{object}	utils.ProblemDetails	"La sesión no existe. Probablemente porque la eliminaron"
//
//	@Router		/api/auth/password [patch]
func (*HttpAuthController) ChangePassword(c *gin.Context) {
	var ChangePassword *dto.ChangePassword

	if err := c.BindJSON(&ChangePassword); err != nil {
		utils.ResErrValidators(c, err)
		return
	}
	err := authService.ChangePassword(*ChangePassword)
	if err != nil {
		utils.ResFromErr(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
