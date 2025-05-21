package controller

import (
	"errors"
	"net/http"

	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/dto"
	"github.com/CPU-commits/Template_Go-EventDriven/src/cmd/http/utils"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type HttpRecoveryController struct{}

// Generate godoc
//
//	@Summary	Generar codigo de recuperacion
//	@Tags		auth
//	@Success	200		{object}	controller.RecoveryCodeResponse
//	@Failure	503		{object}	utils.ProblemDetails	"Error con la base de datos"
//
//	@Failure	403		{object}	utils.ProblemDetails	"Credenciales inválidas"
//
//	@Router		/api/auth/recovery [post]
func (*HttpRecoveryController) GenerateRecoveryCode(c *gin.Context) {
	var GetRecoveryCodeDto *dto.GetRecoveryCode
	if err := c.BindJSON(&GetRecoveryCodeDto); err != nil {
		localizer := utils.GetI18nLocalizer(c)
		errors := utils.ValidatorErrorToErrorProblemDetails(err, localizer)

		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			utils.ProblemDetails{
				Title: localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: utils.FORM_ERR,
				}),
				Errors: errors,
			},
		)
		return
	}
	code, err := recoveryService.RecoveryCode(GetRecoveryCodeDto.Email)
	if err != nil {
		utils.ResFromErr(c, err)
		return
	}

	c.JSON(http.StatusCreated, RecoveryCodeResponse{
		Recovery: code,
	})
}

// Verify code godoc
//
//	@Summary	Verificar codigo de recuperacion
//	@Tags		auth
//	@Success	200		{object}	controller.VerifyRecoveryCodeResponse
//	@Failure	503		{object}	utils.ProblemDetails	"Error con la base de datos"
//	@Param		email	query		int						true	"email"
//	@Param		code			query		string						true	"codigo de recuperacion"
//	@Failure	403		{object}	utils.ProblemDetails	"Credenciales inválidas"
//
//	@Router		/api/auth/recovery/verify [post]
func (*HttpRecoveryController) VerifyRecoveryCode(c *gin.Context) {
	email := c.DefaultQuery("email", "")
	code := c.DefaultQuery("code", "")
	if email == "" {
		err := errors.New("email is required")
		utils.ResWithMessageID(c, utils.FORM_ERR, http.StatusBadRequest, err)
		return
	}
	if code == "" {
		err := errors.New("code is required")
		utils.ResWithMessageID(c, utils.FORM_ERR, http.StatusBadRequest, err)
		return
	}

	exists, token, err := recoveryService.VerifyRecoveryCode(&dto.VerifyRecoveryCode{
		Email: email,
		Code:  code,
	})
	if err != nil {
		utils.ResFromErr(c, err)
		return
	}

	c.JSON(http.StatusOK, VerifyRecoveryCodeResponse{
		Exists: exists,
		Token:  token,
	})
}
