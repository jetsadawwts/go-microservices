package paymentHandler

import "github.com/jetsadawwts/go-microservices/modules/payment/paymentUsecase"

type (
	paymentQueueHandler struct {
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentQueueHandler(paymentUsecase paymentUsecase.PaymentUsecaseService) paymentQueueHandler {
	return paymentQueueHandler{
		paymentUsecase: paymentUsecase,
	}
}
