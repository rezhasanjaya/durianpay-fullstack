package repository

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
)

type PaymentRepository interface {
	GetPayments(id string, status string) ([]*entity.Payment, error)
	DashboardWidget() (string, error)
}

type Payment struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) *Payment {
	return &Payment{db: db}
}

func (r *Payment) GetPayments(id string, status string) ([]*entity.Payment, error) {
	query := "SELECT id, merchant, date, amount, status FROM payments"
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
		if err := data.Scan(&payment.ID, &payment.Merchant, &payment.Date, &payment.Amount, &payment.Status); err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}

	return payments, nil
}

func (r *Payment) DashboardWidget() (string, error) {
	query := `SELECT status, COUNT(*) AS TOTAL FROM payments GROUP BY status`

	data, err := r.db.Query(query)
	if err != nil {
		return "", err
	}
	defer data.Close()

	var result string
	for data.Next() {
		var status string
		var total int
		if err := data.Scan(&status, &total); err != nil {
			return "", err
		}
		result += status + ": " + strconv.Itoa(total) + "\n"
	}

	return result, nil
}
