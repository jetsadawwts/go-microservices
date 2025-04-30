package paymentHandler

import "githib.coom/jetsadawwts/go-microservices/modules/payment/paymentUsecase"

type (
	paymentGrpcHandler struct {
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentGrpcHandler(paymentUsecase paymentUsecase.PaymentUsecaseService) paymentGrpcHandler {
	return paymentGrpcHandler{
		paymentUsecase: paymentUsecase,
	}
}
