package container

import (
	"database/sql"
	historyRepo "skeleton-fiber-clean-architecture/application/history/repository"
	historyUsecase "skeleton-fiber-clean-architecture/application/history/usecase"
	userRepo "skeleton-fiber-clean-architecture/application/user/repository"
	userUsecase "skeleton-fiber-clean-architecture/application/user/usecase"
	"skeleton-fiber-clean-architecture/domain/history"
	"skeleton-fiber-clean-architecture/domain/user"
	"skeleton-fiber-clean-architecture/infrastructure/database"
	"skeleton-fiber-clean-architecture/infrastructure/logger"
)

type Container struct {
	DB                *sql.DB
	UserRepository    user.UserRepository
	HistoryRepository history.HistoryRepository
	UserUseCase       *userUsecase.UserUseCase
	HistoryUseCase    *historyUsecase.HistoryUseCase
}

func NewContainer() (*Container, error) {
	db, err := database.NewDBConnection()
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	userRepository := userRepo.NewUserRepository(db)
	historyRepository := historyRepo.NewHistoryRepository(db)

	historyUseCase := &historyUsecase.HistoryUseCase{HistoryRepository: historyRepository}
	userUseCase := &userUsecase.UserUseCase{UserRepository: userRepository, HistoryUseCase: historyUseCase}

	return &Container{
		DB:                db,
		UserRepository:    userRepository,
		HistoryRepository: historyRepository,
		UserUseCase:       userUseCase,
		HistoryUseCase:    historyUseCase,
	}, nil
}
