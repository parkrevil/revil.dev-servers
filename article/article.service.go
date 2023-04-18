package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	empty "github.com/golang/protobuf/ptypes/empty"
	pb "revil.dev-servers/libs/services/article"
)

type Article struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ArticleService struct {
	articles []Article
	pb.UnimplementedArticleServiceServer
}

func NewArticleService() *ArticleService {
	var articles []Article

	content, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = json.Unmarshal(content, &articles)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return &ArticleService{articles: articles}
}

func (as *ArticleService) GetArticles(ctx context.Context, in *empty.Empty) (*pb.GetArticlesResponse, error) {
	var articles []*pb.Article

	for _, article := range as.articles {
		articles = append(articles, &pb.Article{
			Id:          article.Id,
			Title:       article.Title,
			Description: article.Description,
		})
	}

	return &pb.GetArticlesResponse{Articles: articles}, nil
}

func (as *ArticleService) CreateArticle(ctx context.Context, in *pb.CreateArticleParam) (*pb.ArticleId, error) {
	return &pb.ArticleId{Value: "test"}, nil
}

func (as *ArticleService) DeleteArticle(ctx context.Context, in *pb.ArticleId) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
