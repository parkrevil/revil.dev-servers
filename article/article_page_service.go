package main

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"
	pb "revil.dev-servers/libs/services/article"
)

type ArticlePageService struct {
	pb.UnimplementedArticlePageServiceServer
}

func (s *ArticleService) AddPage(ctx context.Context, in *pb.AddPageParam) (*pb.PageId, error) {
	return &pb.PageId{Value: "page_test"}, nil
}

func (s *ArticleService) DeletePage(ctx context.Context, in *pb.PageId) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
