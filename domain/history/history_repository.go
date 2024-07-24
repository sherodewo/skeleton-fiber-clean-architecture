package history

type HistoryRepository interface {
	GetHistoryByID(id int) (*History, error)
	CreateHistory(history *History) error
}
