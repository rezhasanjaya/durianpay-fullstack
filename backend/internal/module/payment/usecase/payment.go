package usecase

import (
	"fmt"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/durianpay/fullstack-boilerplate/internal/module/payment/repository"
)

type PaymentUsecase interface {
	GetPayments() ([]*entity.Payment, error)
}

type Payment struct {
	repo repository.PaymentRepository
}

func NewPaymentUsecase(repo repository.PaymentRepository) *Payment {
	return &Payment{repo: repo}
}

func (p *Payment) GetPayments() ([]*entity.Payment, error) {
	payments, err := p.repo.GetPayments()
	if err != nil {
		return nil, err
	}
	fmt.Println("Testing Payments : ", payments)

	return payments, nil
}
