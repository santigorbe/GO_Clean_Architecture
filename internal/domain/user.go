package domain

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository interface {
	GetByID(id string) (*User, error)
	GetAll() ([]*User, error)
	Create(user *User) error
}
