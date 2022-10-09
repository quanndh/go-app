package middlewares

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/quanndh/go-app/public/services"
	"net/http"
	"strings"
)

const (
	TokenTypeBasicAuthen = "Basic"
	TokenBearerAuthen    = "Bearer"
	AuthorizationHeader  = "Authorization"
)

func MiddlewareAuthentication(jwtService services.IJwtService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get(AuthorizationHeader)
		token := strings.SplitN(auth, " ", 2)

		if len(token) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		if token[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		claims, err := jwtService.Verify(token[1])

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		ctx.Set("UserId", claims.UserId)
		ctx.Next()

	}
}

func genBasicAuthenToken(m map[string]string) []string {
	var tokens []string
	for k, v := range m {
		token := base64.StdEncoding.EncodeToString([]byte(
			fmt.Sprintf("%s:%s", k, v)))
		tokens = append(tokens, token)
	}
	return tokens
}

// CheckStringInSlice CheckStringInSlice
func checkStringInSlice(s string, l []string) bool {
	for _, item := range l {
		if item == s {
			return true
		}
	}
	return false
}
