package service

import (
	"context"
	"github.com/gofc/graphql-demo/internal/model"
	"github.com/gofc/graphql-demo/internal/repository"
	"github.com/gofc/graphql-demo/pkg/logger"
	"go.uber.org/zap"
)

type TodoService struct {
	repo *repository.Repository
}

func (s *TodoService) CreateTodo(ctx context.Context, input model.NewTodo) (res *model.Todo, err error) {
	logger.Info(ctx, "create new todo",
		zap.String("text", input.Text),
		zap.String("userID", input.UserID))
	res, err = s.repo.Todo.Create(ctx, &model.Todo{
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	})
	if err != nil {
		logger.Error(ctx, "failed to call todoRepository.Create", zap.Error(err))
	}
	return
}

func (s *TodoService) ListTodos(ctx context.Context) (res []*model.Todo, err error) {
	logger.Info(ctx, "ListTodos called")
	res, err = s.repo.Todo.FindAll()
	if err != nil {
		logger.Error(ctx, "failed to call todoRepository.FindAll", zap.Error(err))
	}
	return
}

func (s *TodoService) GetTodoOwner(ctx context.Context, input *model.Todo) (res *model.User, err error) {
	logger.Info(ctx, "GetTodoOwner called", zap.String("UserID", input.UserID))
	res, err = s.repo.User.FindOne(input.UserID)
	if err != nil {
		logger.Error(ctx, "failed to call repo.User.FindOne", zap.Error(err))
	}
	return
}

func NewTodoService(repo *repository.Repository) *TodoService {
	return &TodoService{repo: repo}
}
