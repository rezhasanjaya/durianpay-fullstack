package entity

type Payment struct {
	ID       string `json:"id"`
	Merchant string `json:"merchant"`
	Date     string `json:"created_at"`
	Amount   string `json:"amount"`
	Status   string `json:"status"`
}
