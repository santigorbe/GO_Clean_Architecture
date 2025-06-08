package repository

import (
	"errors"

	"github.com/santigorbe/clean_arq/internal/domain"
)

type InMemoryUserRepo struct {
	users map[string]*domain.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{users: make(map[string]*domain.User)}
}

func (r *InMemoryUserRepo) GetByID(id string) (*domain.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepo) GetAll() ([]*domain.User, error) {
	var result []*domain.User
	for _, u := range r.users {
		result = append(result, u)
	}
	return result, nil
}

func (r *InMemoryUserRepo) Create(user *domain.User) error {
	if _, exists := r.users[user.ID]; exists {
		return errors.New("user already exists")
	}
	r.users[user.ID] = user
	return nil
}
