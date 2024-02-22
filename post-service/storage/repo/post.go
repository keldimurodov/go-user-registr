package repo

import (
	pb "registration/post-service/genproto/post-service"
)

// PostStorageI ...
type PostStorageI interface {
	Create(post *pb.Post) (*pb.Post, error)
	GetPost(id *pb.GetRequest) (*pb.Post, error)
	GetUserPosts(req *pb.GetUserRequest) ([]*pb.Post, error)
	//Delete(user *pb.User) (*pb.User, error)
	//Update(user *pb.User) (*pb.User, error)
}
