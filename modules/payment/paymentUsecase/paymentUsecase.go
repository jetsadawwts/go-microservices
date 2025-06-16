package paymentUsecase

import "github.com/jetsadawwts/go-microservices/modules/payment/paymentRepository"

type (
	PaymentUsecaseService interface{}
	paymentUsecase        struct {
		paymentRepository paymentRepository.PaymentRepositoryService
	}
)

func NewPaymentUsecase(paymentRepository paymentRepository.PaymentRepositoryService) PaymentUsecaseService {
	return &paymentUsecase{
		paymentRepository: paymentRepository,
	}
}
