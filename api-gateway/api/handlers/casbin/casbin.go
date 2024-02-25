package casbin

import (
	"errors"
	tokens "go-user-registr/api-gateway/api/tokens"
	"go-user-registr/api-gateway/config"
	"net/http"
	"strings"
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

var t string

type CasbinHandler struct {
	config   config.Config
	enforcer *casbin.Enforcer
}

func (casb *CasbinHandler) CheckCasbinPermission(ca *casbin.Enforcer, conf config.Config) gin.HandlerFunc {

	var CasbinHandler = &CasbinHandler{
		config:   conf,
		enforcer: ca,
	}
	return func(ctx *gin.Context) {
		allowed, err := CasbinHandler.CheckPermission(ctx.Request)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
		}

		if !allowed {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "permission denied",
			})
		}
	}

}

func NewAuthorizer() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("Authorization")
		if token == "" {
			return
		}

		claims, err := tokens.ExtractClaim(token, []byte(config.Load().SignInKey))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			return
		}
		sub := claims["role"]
		obj := ctx.Request.URL.Path
		etc := ctx.Request.Method
		fmt.Println(sub, obj, etc)
		e, _ := casbin.NewEnforcer("auth.conf", "auth.csv")
		t, _ := e.Enforce(sub, obj, etc)

		fmt.Println(t)
		if t {
			ctx.Next()
			return
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "net dostupo",
		})
	}

}

func (casb *CasbinHandler) GetRole(ctx *http.Request) (string, int) {

	token := ctx.Header.Get("Authorization")

	if token == "" {
		return "Unauthorized", http.StatusUnauthorized
	} else if strings.Contains(token, "Bearer") {
		t = strings.Trim(token, "Bearer ")
	} else {
		t = token
	}

	clams, err := tokens.ExtractClaim(token, []byte(config.Load().SignInKey))
	if err != nil {
		return "Unauthorized, token is invalid", http.StatusUnauthorized
	}

	return cast.ToString(clams["role"]), 0
}

func (casb *CasbinHandler) CheckPermission(r *http.Request) (bool, error) {

	role, status := casb.GetRole(r)
	if role == "Unauthorized" {
		return true, nil
	}
	if status != 0 {
		return false, errors.New(role)
	}

	method := r.Method
	action := r.URL.Path

	c, err := casb.enforcer.Enforce(role, action, method)

	if err != nil {
		return false, nil
	}

	return c, nil
}
