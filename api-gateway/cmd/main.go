package main

import (
	"fmt"
	"go-user-registr/api-gateway/api"
	"go-user-registr/api-gateway/config"
	"go-user-registr/api-gateway/pkg/logger"
	"go-user-registr/api-gateway/services"

	"github.com/casbin/casbin/v2"

	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	psqlString := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`, "localhost", 5432, "postgres", "123", "users")

	db, err := gormadapter.NewAdapter("postgres", psqlString, true)
	if err != nil {
		log.Error("gormadapter error", logger.Error(err))
	}

	enforcer, err := casbin.NewEnforcer("auth.conf", db)
	if err != nil {
		log.Error("NewEnforcer error", logger.Error(err))
		return
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		Enforcer:       enforcer,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
