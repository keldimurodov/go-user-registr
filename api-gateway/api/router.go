package api

import (
	"github.com/casbin/casbin/v2"
	_ "go-user-registr/api-gateway/api/docs" // swag
	v1 "go-user-registr/api-gateway/api/handlers/v1"
	t "go-user-registr/api-gateway/api/tokens"
	"go-user-registr/api-gateway/config"
	"go-user-registr/api-gateway/pkg/logger"
	"go-user-registr/api-gateway/services"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	jwtHandler     t.JWTHandler
	Enforcer       *casbin.Enforcer
	ServiceManager services.IServiceManager
}

// @Title Welcome to RegisterPAGE
// @Version 1.0
// @Description This is a example of Social Network
// @Host localhost:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Enforcer:       *option.Enforcer,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	//handler := ccasbin.CasbinHandler{}
	//api.Use(handler.CheckCasbinPermission(option.Enforcer, option.Conf))
	// users
	api.POST("/users", handlerV1.CreateUser)
	api.GET("/users/:id", handlerV1.GetUser)
	api.GET("/users", handlerV1.GetAll)
	api.PUT("/users/:id", handlerV1.UpdateUser)
	api.DELETE("/users/:id", handlerV1.DeleteUser)

	// register
	api.POST("/register", handlerV1.SignUp)
	api.GET("/login", handlerV1.LogIn)
	api.GET("/verification", handlerV1.Verification)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
