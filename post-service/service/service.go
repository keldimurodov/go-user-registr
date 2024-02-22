package service

import (
	"context"
	pb "registration/post-service/genproto/post-service"
	pbu "registration/post-service/genproto/user-service"
	l "registration/post-service/pkg/logger"
	grpcClient "registration/post-service/service/grpc_client"
	"registration/post-service/storage"

	"github.com/jmoiron/sqlx"
)

// PostService ...
type PostService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.IServiceManager
}

// NewPostService ...
func NewPostService(db *sqlx.DB, log l.Logger, client grpcClient.IServiceManager) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (s *PostService) Create(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	user, err := s.storage.Post().Create(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *PostService) GetPost(ctx context.Context, req *pb.GetRequest) (*pb.PostResponse, error) {
	post, err := s.storage.Post().GetPost(req)
	if err != nil {
		s.logger.Error(err.Error())
	}

	user, err := s.client.TemplateService().GetUser(ctx, &pbu.GetRequest{
		UserId: post.OwnerId,
	})

	if err != nil {
		return nil, err
	}

	return &pb.PostResponse{
		Id:       post.Id,
		Title:    post.Title,
		ImageUrl: post.ImageUrl,
		Owner: &pb.Owner{
			Id:       user.Id,
			Name:     user.Name,
			LastName: user.LastName,
		},
	}, nil
}

func (s *PostService) GetUserPosts(ctx context.Context, req *pb.GetUserRequest) (*pb.GetAllResponse, error) {
	posts, err := s.storage.Post().GetUserPosts(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return &pb.GetAllResponse{
		Posts: posts,
	}, nil
}
