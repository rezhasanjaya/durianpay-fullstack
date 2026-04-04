package entity

type Payment struct {
	ID       string `json:"id"`
	Merchant string `json:"merchant"`
	Amount   string `json:"amount"`
	Status   string `json:"status"`
	CreatedAt string `json:"created_at"`
}
