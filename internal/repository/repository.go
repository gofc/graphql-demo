package repository

type Repository struct {
	Todo *TodoRepository
	User *UserRepository
}

func NewRepository() *Repository {
	return &Repository{
		Todo: NewTodoRepository(),
		User: NewUserRepository(),
	}
}
