package repository

import (
	"database/sql"
	"skeleton-fiber-clean-architecture/domain/user"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) user.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) GetUserByID(id int) (*user.User, error) {
	var u user.User
	err := r.DB.QueryRow("SELECT id, name, email, password FROM users WHERE id = ?", id).Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepositoryImpl) CreateUser(u *user.User) error {
	_, err := r.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", u.Name, u.Email, u.Password)
	return err
}
