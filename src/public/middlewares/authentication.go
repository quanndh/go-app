package middlewares

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	TokenTypeBasicAuthen = "Basic"
	AuthorizationHeader  = "Authorization"
)

func MiddlewareAuthentication(authInfo string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get(AuthorizationHeader)
		token := strings.SplitN(auth, " ", 2)

		if len(token) != 2 {
			// TODO Build Error Response
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		switch token[0] {
		case TokenTypeBasicAuthen:
			var auth map[string]string
			err := json.Unmarshal([]byte(authInfo), &auth)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, "unmarshal basic data error")
				return
			}

			tokens := genBasicAuthenToken(auth)
			if !checkStringInSlice(token[1], tokens) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
				return
			}
		default:
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}
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
