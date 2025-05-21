package middleware

import (
	"net/http"

	"github.com/CPU-commits/Template_Go-EventDriven/src/cmd/http/utils"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func RolesMiddleware(roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := utils.NewClaimsFromContext(ctx)
		userRoles := claims.Roles
		for _, allowRole := range roles {
			for _, userRole := range userRoles {
				if allowRole == userRole {
					ctx.Next()
					return
				}

			}
		}
		localizer := utils.GetI18nLocalizer(ctx)

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, &utils.ProblemDetails{
			Title: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "roles.unauthorized",
			}),
		})
	}
}
