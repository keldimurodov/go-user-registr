package services

import (
	"fmt"

	"api-gateway/config"
	pbp "api-gateway/genproto/post"
	pbu "api-gateway/genproto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	UserService() pbu.UserServiceClient
	PostService() pbp.PostServiceClient
}

type serviceManager struct {
	userService pbu.UserServiceClient
	postService pbp.PostServiceClient
}

func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}

func (s *serviceManager) PostService() pbp.PostServiceClient {
	return s.postService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		userService: pbu.NewUserServiceClient(connUser),
		postService: pbp.NewPostServiceClient(connPost),
	}

	return serviceManager, nil
}
