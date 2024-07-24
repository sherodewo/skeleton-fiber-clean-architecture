package history

type History struct {
	ID        int
	ItemName  string
	Quantity  int
	Action    string // "in" atau "out"
	CreatedBy int    // ID User
}
