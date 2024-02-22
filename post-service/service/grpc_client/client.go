package grpcClient

import (
	"fmt"
	"registration/post-service/config"
	pbu "registration/post-service/genproto/user-service"

	"google.golang.org/grpc"
)

type IServiceManager interface {
	TemplateService() pbu.UserServiceClient
}
type serviceManager struct {
	cfg             config.Config
	templateService pbu.UserServiceClient
}

func (s *serviceManager) TemplateService() pbu.UserServiceClient {
	return s.templateService

}

func New(cfg config.Config) (IServiceManager, error) {
	templateConnection, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &serviceManager{
		cfg:             cfg,
		templateService: pbu.NewUserServiceClient(templateConnection)}, nil

}
