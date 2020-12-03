package dataloader

import (
	"context"
	"github.com/gofc/graphql-demo/internal/model"
	"github.com/gofc/graphql-demo/internal/repository"
	"github.com/gofc/graphql-demo/pkg/logger"
	"go.uber.org/zap"
	"time"
)

func DefaultUserLoader(repo *repository.Repository) *UserLoader {
	return &UserLoader{
		wait:     1 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*model.User, []error) {
			logger.Info(context.Background(), "UserLoader called", zap.Strings("UserIDs", keys))
			users, err := repo.User.FindByIDs(keys)
			if err != nil {
				return nil, []error{err}
			}
			return users, nil
		},
	}
}
