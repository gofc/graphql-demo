package resolver

import "github.com/gofc/graphql-demo/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todoService *service.TodoService
}

func NewResolver(todoService *service.TodoService) *Resolver {
	return &Resolver{todoService: todoService}
}
