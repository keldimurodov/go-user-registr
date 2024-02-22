package main

import (
	"net"
	"registration/post-service/config"
	pb "registration/post-service/genproto/post-service"
	"registration/post-service/pkg/db"
	"registration/post-service/pkg/logger"
	"registration/post-service/service"
	grpcClient "registration/post-service/service/grpc_client"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "template-post-	service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatasbase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	grpcClien, err := grpcClient.New(cfg)

	if err != nil {
		log.Fatal("grpc client dial error", logger.Error(err))
	}

	postService := service.NewPostService(connDB, log, grpcClien)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, postService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
