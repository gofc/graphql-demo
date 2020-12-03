package repository

import (
	"context"
	"fmt"
	"github.com/gofc/graphql-demo/internal/model"
	"strconv"
)

type ArticleRepository struct {
	id    int
	items []*model.Article
}

func (r *ArticleRepository) Create(_ context.Context, m *model.Article) (*model.Article, error) {
	m.ID = fmt.Sprintf("%d", r.id)
	r.id++
	return m, nil
}

func (r *ArticleRepository) FindAll() ([]*model.Article, error) {
	return r.items, nil
}

func NewArticleRepository() *ArticleRepository {
	items := make([]*model.Article, 0, 10)
	id := 1
	for ; id < 11; id++ {
		items = append(items, &model.Article{
			ID:     strconv.Itoa(id),
			Title:  fmt.Sprintf("ArticleTitle %d", id),
			Author: fmt.Sprintf("%d", id),
		})
	}
	return &ArticleRepository{id: id, items: items}
}
