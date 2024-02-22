package casbin

import (
	tokens "api-gateway/api"
	"api-gateway/config"
	"net/http"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type CasbinHandler struct {
	config   config.Config
	enforcer casbin.Enforcer
}

func (casb *CasbinHandler) CheckCasbinPermission(casbin *casbin.Enforcer, conf config.Config) gin.HandlerFunc {

	CasbinHandler := &CasbinHandler{
		config:   conf,
		enforcer: casbin,
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

func (casb *CasbinHandler) GetRole(ctx *gin.Context) (string, int) {

	var t string

	token := ctx.GetHeader("Authorization")

	if token == "" {
		return "Unauthorized", http.StatusUnauthorized
	} else if strings.Contains(token, "Bearer") {
		t = strings.Trim(token, "Bearer ")
	} else {
		t = token
	}

	clams, err := tokens.ExtractClaim(token, []byte(conf.SignInKey))
	if err != nil {
		return "Unauthorized, token is invalid", http.StatusUnauthorized
	}

	return cast.ToString(clams["role"]), 0
}

func (casb *CasbinHandler) CheckPermission(r *http.Request) (bool, error) {
	role, status := GetRole(r, cfg)
	if status != 0 {
		return false, erors.New(role)
	}

	method := r.Method
	action := r.URL.Path

	c, err := casb.enforcer.Enforce(role, action, method)

}
