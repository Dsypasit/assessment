package expense

type Expense struct {
	ID     int      `json:"id"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
	Title  string   `json:"title"`
}
