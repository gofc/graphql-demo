package service

import "github.com/gofc/graphql-demo/internal/repository"

type ArticleService struct {
	repo *repository.Repository
}

func NewArticleService(repo *repository.Repository) *ArticleService {
	return &ArticleService{repo: repo}
}
