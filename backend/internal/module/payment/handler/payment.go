package handler

import (
	"encoding/json"
	"net/http"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	paymentUsecase "github.com/durianpay/fullstack-boilerplate/internal/module/payment/usecase"
	"github.com/durianpay/fullstack-boilerplate/internal/transport"
)

type PaymentHandler struct {
	paymentUC paymentUsecase.PaymentUsecase
}

type PaymentResponse struct {
	ID       string `json:"id"`
	Merchant string `json:"merchant"`
	Date     string `json:"date"`
	Amount   string `json:"amount"`
	Status   string `json:"status"`
}

func NewPaymentHandler(paymentUC paymentUsecase.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{
		paymentUC: paymentUC,
	}
}

func (p *PaymentHandler) GetDashboardV1Payments(w http.ResponseWriter, r *http.Request) {
	// TODO add implementations
	data, err := p.paymentUC.GetPayments()
	if err != nil {
		transport.WriteError(w, err)
		return
	}

	var response []PaymentResponse
	for _, payment := range data {
		response = append(response, PaymentResponse{
			ID:       payment.ID,
			Merchant: payment.Merchant,
			Date:     payment.Date,
			Amount:   payment.Amount,
			Status:   payment.Status,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		transport.WriteAppError(w, entity.ErrorInternal("internal server error"))
		return
	}
}
