package api

import (
	"net/http"

	ah "github.com/durianpay/fullstack-boilerplate/internal/module/auth/handler"
	py "github.com/durianpay/fullstack-boilerplate/internal/module/payment/handler"
	"github.com/durianpay/fullstack-boilerplate/internal/openapigen"
)

type APIHandler struct {
	Auth    *ah.AuthHandler
	Payment *py.PaymentHandler
}

var _ openapigen.ServerInterface = (*APIHandler)(nil)

func (h *APIHandler) PostDashboardV1AuthLogin(w http.ResponseWriter, r *http.Request) {
	h.Auth.PostDashboardV1AuthLogin(w, r)
}

func (h *APIHandler) GetDashboardV1Payments(w http.ResponseWriter, r *http.Request, body openapigen.GetDashboardV1PaymentsParams) {
	h.Payment.GetDashboardV1Payments(w, r, body)
}

func (h *APIHandler) GetDashboardV1Widget(w http.ResponseWriter, r *http.Request) {
	h.Payment.GetDashboardV1PaymentsWidget(w, r)
}
