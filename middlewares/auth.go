package middlewares

import (
	"api-rs/utility"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			ctx.Abort()
			utility.ApiResponse(ctx, http.StatusUnauthorized, "unauthorized", nil, nil)
			return
		}

		token := strings.Replace(authHeader, "Bearer ", "", -1)
		claims, err := utility.VerifyToken(token)
		if err != nil {
			ctx.Abort()
			utility.ApiResponse(ctx, http.StatusUnauthorized, err.Error(), nil, nil)
			return
		}

		ctx.Set("user", claims)
	}
}
