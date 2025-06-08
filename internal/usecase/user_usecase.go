package usecase

import "github.com/santigorbe/clean_arq/internal/domain"

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(r domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (uc *UserUseCase) GetUser(id string) (*domain.User, error) {
	return uc.repo.GetByID(id)
}

func (uc *UserUseCase) ListUsers() ([]*domain.User, error) {
	return uc.repo.GetAll()
}

func (uc *UserUseCase) CreateUser(user *domain.User) error {
	return uc.repo.Create(user)
}
