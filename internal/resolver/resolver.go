package resolver

import (
	"github.com/gofc/graphql-demo/internal/dataloader"
	"github.com/gofc/graphql-demo/internal/repository"
	"github.com/gofc/graphql-demo/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todoService *service.TodoService
	userLoader  *dataloader.UserLoader
}

func NewResolver(repo *repository.Repository) *Resolver {
	return &Resolver{
		todoService: service.NewTodoService(repo),
		userLoader:  dataloader.DefaultUserLoader(repo),
	}
}
