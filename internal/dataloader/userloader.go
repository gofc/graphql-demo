package dataloader

import (
	"github.com/gofc/graphql-demo/internal/model"
	"github.com/gofc/graphql-demo/internal/repository"
	"time"
)

func DefaultUserLoader(repo *repository.Repository) *UserLoader {
	return &UserLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*model.User, []error) {
			users, err := repo.User.FindByIDs(keys)
			if err != nil {
				panic(err) //todo need improve
			}
			return users, nil
		},
	}
}
