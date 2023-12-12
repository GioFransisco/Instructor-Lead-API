package middleware

import (
	"net/http"
	"strings"
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type JwtMiddleware interface {
	AuthMiddleware(role ...string) gin.HandlerFunc
}

type jwtMiddleware struct {
	token common.JwtToken
}

func (j *jwtMiddleware) AuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		author := ctx.Request.Header.Get("Authorization")
		tokenString := strings.Replace(author, "Bearer ", "", -1)

		payloadToken := model.TokenModel{
			Token: tokenString,
		}

		claims, err := j.token.VerifyToken(payloadToken)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		roleValidate := false

		for _, v := range roles {
			if claims["role"].(string) == v {
				roleValidate = true
			}
		}

		if !roleValidate {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "can't running invalid role"})
			return
		}

		exp := claims["exp"].(float64)

		if time.Now().After(time.Unix(int64(exp), 10)) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is expired"})
			return
		}

		ctx.Set("userEmail", claims["email"].(string))
		ctx.Set("userRole", claims["role"].(string))
		ctx.Set("userId", claims["userId"].(string))
		ctx.Next()
	}
}

func NewMiddlewareAuth(jwt common.JwtToken) JwtMiddleware {
	return &jwtMiddleware{jwt}
}
