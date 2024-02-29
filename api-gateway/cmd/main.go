package main

import (
	"go-user-registr/api-gateway/api"
	"go-user-registr/api-gateway/config"
	"go-user-registr/api-gateway/pkg/logger"
	"go-user-registr/api-gateway/services"
	"go-user-registr/api-gateway/queue/kafka/producer"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	writer, err := producer.NewKafkaProducerInit([]string{"localhost:9092"})
	if err != nil {
		log.Error("NewKafkaProducerInit: %v", logger.Error(err))
	}

	err = writer.ProduceMessage("test-topic", []byte("message"))
	if err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
	}

	defer writer.Close()

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
