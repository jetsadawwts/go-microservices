package paymentHandler

import (
	"githib.coom/jetsadawwts/go-microservices/config"
	"githib.coom/jetsadawwts/go-microservices/modules/payment/paymentUsecase"
)

type (
	PaymentHttpHandler interface{}
	paymentHttpHandler        struct {
		cfg            *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentHttpHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentHttpHandler {
	return &paymentHttpHandler{
		cfg:            cfg,
		paymentUsecase: paymentUsecase,
	}
}
