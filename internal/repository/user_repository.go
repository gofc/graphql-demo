package repository

import (
	"context"
	"fmt"
	"github.com/gofc/graphql-demo/internal/model"
	"github.com/pkg/errors"
	"strconv"
)

type UserRepository struct {
	id    int
	items []*model.User
}

func (r *UserRepository) Create(_ context.Context, m *model.User) (*model.User, error) {
	m.ID = fmt.Sprintf("%d", r.id)
	r.id++
	return m, nil
}

func (r *UserRepository) FindAll() ([]*model.User, error) {
	return r.items, nil
}

func (r *UserRepository) FindOne(id string) (*model.User, error) {
	for _, v := range r.items {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, errors.New("user not found")
}

func NewUserRepository() *UserRepository {
	items := make([]*model.User, 0, 10)
	id := 1
	for ; id < 11; id++ {
		items = append(items, &model.User{
			ID:   strconv.Itoa(id),
			Name: fmt.Sprintf("Name%d", id),
		})
	}
	return &UserRepository{id: id, items: items}
}
