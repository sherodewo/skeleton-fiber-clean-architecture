package usecase

import (
	"skeleton-fiber-clean-architecture/domain/history"
)

type HistoryUseCase struct {
	HistoryRepository history.HistoryRepository
}

func (uc *HistoryUseCase) GetHistory(id int) (*history.History, error) {
	return uc.HistoryRepository.GetHistoryByID(id)
}

func (uc *HistoryUseCase) CreateHistory(itemName string, quantity int, action string, createdBy int) error {
	newHistory := &history.History{
		ItemName:  itemName,
		Quantity:  quantity,
		Action:    action,
		CreatedBy: createdBy,
	}

	return uc.HistoryRepository.CreateHistory(newHistory)
}
