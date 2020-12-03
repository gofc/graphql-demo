package service

import (
	"context"
	"github.com/gofc/graphql-demo/internal/model"
	"github.com/gofc/graphql-demo/internal/repository"
	"github.com/gofc/graphql-demo/pkg/logger"
	"go.uber.org/zap"
)

type UserService struct {
	repo *repository.Repository
}

func (s *UserService) ListUserByIDs(ctx context.Context, ids []string) (res []*model.User, err error) {
	logger.Info(ctx, "ListUserByIDs called", zap.Strings("UserIDs", ids))
	res, err = s.repo.User.FindByIDs(ids)
	if err != nil {
		logger.Error(ctx, "failed to call s.repo.User.FindByIDs", zap.Error(err))
	}
	return
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}
