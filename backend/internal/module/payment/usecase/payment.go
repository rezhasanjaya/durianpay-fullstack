package usecase

import (
	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/durianpay/fullstack-boilerplate/internal/module/payment/repository"
)

type PaymentUsecase interface {
	GetPayments(id string, status string) ([]*entity.Payment, error)
	DashboardWidget() (string, error)
}

type Payment struct {
	repo repository.PaymentRepository
}

func NewPaymentUsecase(repo repository.PaymentRepository) *Payment {
	return &Payment{repo: repo}
}

func (p *Payment) GetPayments(id string, status string) ([]*entity.Payment, error) {
	payments, err := p.repo.GetPayments(id, status)
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (p *Payment) DashboardWidget() (string, error) {
	widget, err := p.repo.DashboardWidget()
	if err != nil {
		return "", err
	}

	return widget, nil
}
