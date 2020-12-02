package repository

import (
	"context"
	"fmt"
	"github.com/gofc/graphql-demo/internal/model"
	"strconv"
)

type TodoRepository struct {
	id    int
	items []*model.Todo
}

func (r *TodoRepository) Create(_ context.Context, m *model.Todo) (*model.Todo, error) {
	m.ID = fmt.Sprintf("%d", r.id)
	r.id++
	return m, nil
}

func (r *TodoRepository) FindAll() ([]*model.Todo, error) {
	return r.items, nil
}

func NewTodoRepository() *TodoRepository {
	items := make([]*model.Todo, 0, 10)
	id := 1
	for ; id < 11; id++ {
		items = append(items, &model.Todo{
			ID:     strconv.Itoa(id),
			Text:   fmt.Sprintf("TodoText %d", id),
			Done:   false,
			UserID: "1",
		})
	}
	return &TodoRepository{id: id, items: items}
}
