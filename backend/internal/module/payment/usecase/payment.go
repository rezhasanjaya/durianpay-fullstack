package usecase

import (
	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/durianpay/fullstack-boilerplate/internal/module/payment/repository"
)

type PaymentUsecase interface {
	GetPayments(id string, status string) ([]*entity.Payment, error)
	DashboardWidget() ([]repository.Widget, Total, error)
}

type Payment struct {
	repo repository.PaymentRepository
}

type Total struct {
	Total int `json:"total"`
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

func (p *Payment) DashboardWidget() ([]repository.Widget, Total, error) {
	widget, err := p.repo.DashboardWidget()
	countAll, errCount := p.repo.GetTotal()

	if err != nil {
		return nil, Total{}, err
	}

	if errCount != nil {
		return nil, Total{}, errCount
	}

	return widget, Total{Total: countAll}, nil
}
