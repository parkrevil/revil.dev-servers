package main

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"
	pb "revil.dev-servers/libs/services/article"
)

type ArticleService struct {
	pb.UnimplementedArticleServiceServer
}

func (s *ArticleService) CreateArticle(ctx context.Context, in *pb.CreateArticleParam) (*pb.ArticleId, error) {
	return &pb.ArticleId{Value: "test"}, nil
}

func (s *ArticleService) DeleteArticle(ctx context.Context, in *pb.ArticleId) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
