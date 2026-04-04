package repository

import (
	"database/sql"
	"strings"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
)

type PaymentRepository interface {
	GetPayments(id string, status string) ([]*entity.Payment, error)
	DashboardWidget() ([]Widget, error)
	GetTotal() (int, error)
}

type Payment struct {
	db *sql.DB
}

type Widget struct {
	Status string `json:"status"`
	Total  int    `json:"total"`
}

func NewPaymentRepo(db *sql.DB) *Payment {
	return &Payment{db: db}
}

func (r *Payment) GetPayments(id string, status string) ([]*entity.Payment, error) {
	query := "SELECT id, merchant, created_at, amount, status FROM payments"
	conditions := []string{}
	args := []any{}

	if id != "" {
		conditions = append(conditions, "id = ?")
		args = append(args, id)
	}
	if status != "" {
		conditions = append(conditions, "status = ?")
		args = append(args, status)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	data, err := r.db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer data.Close()

	var payments []*entity.Payment
	for data.Next() {
		var payment entity.Payment
		if err := data.Scan(&payment.ID, &payment.Merchant, &payment.CreatedAt, &payment.Amount, &payment.Status); err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}

	return payments, nil
}

func (r *Payment) GetTotal() (int, error) {
	countData := "SELECT COUNT(*) FROM payments"
	var total int
	err := r.db.QueryRow(countData).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *Payment) DashboardWidget() ([]Widget, error) {
	query := `SELECT status, COUNT(*) AS TOTAL FROM payments GROUP BY status`

	data, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	var result []Widget
	for data.Next() {
		var status string
		var total int
		if err := data.Scan(&status, &total); err != nil {
			return nil, err
		}
		result = append(result, Widget{Status: status, Total: total})
	}

	return result, nil
}
