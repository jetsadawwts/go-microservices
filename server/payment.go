package server

import (
	"github.com/jetsadawwts/go-microservices/modules/payment/paymentHandler"
	"github.com/jetsadawwts/go-microservices/modules/payment/paymentRepository"
	"github.com/jetsadawwts/go-microservices/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	repo := paymentRepository.NewPaymentRepository(s.db)
	usecase := paymentUsecase.NewPaymentUsecase(repo)
	httpHandler := paymentHandler.NewPaymentHttpHandler(s.cfg, usecase)
	grpcHandler := paymentHandler.NewPaymentGrpcHandler(usecase)
	queueHandler := paymentHandler.NewPaymentQueueHandler(usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	payment := s.app.Group("/payment_v1")

	payment.GET("", s.healthCheckService)
}
