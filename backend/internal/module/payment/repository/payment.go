package repository

import (
	"database/sql"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
)

type PaymentRepository interface {
	GetPayments() ([]*entity.Payment, error)
}

type Payment struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) *Payment {
	return &Payment{db: db}
}

func (r *Payment) GetPayments() ([]*entity.Payment, error) {
	data, err := r.db.Query(`SELECT id, merchant, date, amount, status FROM payments`)
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
