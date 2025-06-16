package middlewareUsecase

import "github.com/jetsadawwts/go-microservices/modules/middleware/middlewareRepository"

type (
	MiddlewareUsecaseHandler interface{}

	middlewareUsecase struct {
		middlewareRepository middlewareRepository.MiddlewareRepositoryHandler
	}
)

func NewMiddlewareUsecase(middlewareRepository middlewareRepository.MiddlewareRepositoryHandler) MiddlewareUsecaseHandler {
	return &middlewareUsecase{
		middlewareRepository: middlewareRepository,
	}
}
