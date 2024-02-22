package postgres

import (
	pb "registration/post-service/genproto/post-service"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type postRepo struct {
	db *sqlx.DB
}

// NewPostRepo ...
func NewPostRepo(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}

func (r *postRepo) Create(post *pb.Post) (*pb.Post, error) {
	if post.Id == "" {
		id, err := uuid.NewUUID()
		if err != nil {
			return nil, err
		}
		post.Id = id.String()
	}

	var res pb.Post
	query := `INSERT INTO posts(id, title, image_url, owner_id) VALUES ($1, $2, $3, $4) RETURNING id, title, image_url,owner_id`
	err := r.db.QueryRow(query, post.Id, post.Title, post.ImageUrl, post.OwnerId).Scan(
		&res.Id,
		&res.Title,
		&res.ImageUrl,
		&res.OwnerId)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *postRepo) GetPost(id *pb.GetRequest) (*pb.Post, error) {
	var post pb.Post
	query := `select id, title, image_url, owner_id from posts where id=$1`
	err := r.db.QueryRow(query, id.PostId).Scan(
		&post.Id,
		&post.Title,
		&post.ImageUrl,
		&post.OwnerId)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *postRepo) GetUserPosts(post *pb.GetUserRequest) ([]*pb.Post, error) {

	query := `SELECT id, title, image_url
	FROM posts 
	WHERE owner_id = $1;`
	rows, err := r.db.Query(query, post.UserId)
	if err != nil {
		return nil, err
	}
	var posts []*pb.Post
	for rows.Next() {
		var post pb.Post
		err := rows.Scan(&post.Id, &post.Title, &post.ImageUrl)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}
