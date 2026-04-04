package handler

import (
	"encoding/json"
	"net/http"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/durianpay/fullstack-boilerplate/internal/module/payment/repository"
	paymentUsecase "github.com/durianpay/fullstack-boilerplate/internal/module/payment/usecase"
	"github.com/durianpay/fullstack-boilerplate/internal/openapigen"
	"github.com/durianpay/fullstack-boilerplate/internal/transport"
)

type PaymentHandler struct {
	paymentUC paymentUsecase.PaymentUsecase
}

type PaymentResponse struct {
	ID       string `json:"id"`
	Merchant string `json:"merchant"`
	CreatedAt string `json:"created_at"`
	Amount   string `json:"amount"`
	Status   string `json:"status"`
}

type Response struct {
	Payments []PaymentResponse `json:"payments"`
}

type WidgetResponse struct {
	Total  int                 `json:"total"`
	Widget []repository.Widget `json:"widget"`
}

func NewPaymentHandler(paymentUC paymentUsecase.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{
		paymentUC: paymentUC,
	}
}

func (p *PaymentHandler) GetDashboardV1Payments(w http.ResponseWriter, r *http.Request, params openapigen.GetDashboardV1PaymentsParams) {
	var status string
	if params.Status != nil {
		status = *params.Status
	}

	var id string
	if params.Id != nil {
		id = *params.Id
	}

	data, err := p.paymentUC.GetPayments(id, status)
	if err != nil {
		transport.WriteError(w, err)
		return
	}

	var response []PaymentResponse
	for _, payment := range data {
		response = append(response, PaymentResponse{
			ID:       payment.ID,
			Merchant: payment.Merchant,
			CreatedAt: payment.CreatedAt,
			Amount:   payment.Amount,
			Status:   payment.Status,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Response{Payments: response}); err != nil {
		transport.WriteAppError(w, entity.ErrorInternal("internal server error"))
		return
	}
}

func (p *PaymentHandler) GetDashboardV1PaymentsWidget(w http.ResponseWriter, r *http.Request) {
	dataWidget, dataTotal, err := p.paymentUC.DashboardWidget()
	if err != nil {
		transport.WriteError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(WidgetResponse{Widget: dataWidget, Total: dataTotal.Total}); err != nil {
		transport.WriteAppError(w, entity.ErrorInternal("internal server error"))
		return
	}
}
