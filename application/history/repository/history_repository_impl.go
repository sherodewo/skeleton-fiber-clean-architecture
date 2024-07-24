package repository

import (
	"database/sql"
	"skeleton-fiber-clean-architecture/domain/history"
)

type HistoryRepositoryImpl struct {
	DB *sql.DB
}

func NewHistoryRepository(db *sql.DB) history.HistoryRepository {
	return &HistoryRepositoryImpl{DB: db}
}

func (r *HistoryRepositoryImpl) GetHistoryByID(id int) (*history.History, error) {
	var h history.History
	err := r.DB.QueryRow("SELECT id, item_name, quantity, action, created_by FROM history WHERE id = ?", id).Scan(&h.ID, &h.ItemName, &h.Quantity, &h.Action, &h.CreatedBy)
	if err != nil {
		return nil, err
	}
	return &h, nil
}

func (r *HistoryRepositoryImpl) CreateHistory(h *history.History) error {
	_, err := r.DB.Exec("INSERT INTO history (item_name, quantity, action, created_by) VALUES (?, ?, ?, ?)", h.ItemName, h.Quantity, h.Action, h.CreatedBy)
	return err
}
