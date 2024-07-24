package usecase

import (
	"golang.org/x/crypto/bcrypt"
	historyUsecase "skeleton-fiber-clean-architecture/application/history/usecase"
	"skeleton-fiber-clean-architecture/domain/user"
)

type UserUseCase struct {
	UserRepository user.UserRepository
	HistoryUseCase *historyUsecase.HistoryUseCase
}

func (uc *UserUseCase) GetUser(id int) (*user.User, error) {
	return uc.UserRepository.GetUserByID(id)
}

func (uc *UserUseCase) CreateUser(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &user.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}
	// Buat history entri
	err = uc.HistoryUseCase.CreateHistory("User Created", 1, "in", newUser.ID)
	if err != nil {
		return err
	}
	return uc.UserRepository.CreateUser(newUser)
}
