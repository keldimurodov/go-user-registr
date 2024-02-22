package v1

import (
	"api-gateway/api/handlers/models"
	"api-gateway/config"
	"api-gateway/pkg/logger"
	"api-gateway/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
	}
}

func handleBadRequestWithErrorMessage(c *gin.Context, l logger.Logger, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, models.StandardErrorModel{
			Error: models.Error{
				Message: "Incorrect data supplied",
			},
		})
		l.Error(message, logger.Error(err))
		return true
	}
	return false
}
