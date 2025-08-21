package repository

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRepository struct {
	users []User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: []User{}}
}

func (r *UserRepository) GetUsers() []User {
	return r.users
}

func (r *UserRepository) CreateUser(user User) User {
	user.ID = len(r.users) + 1
	r.users = append(r.users, user)
	return user
}
