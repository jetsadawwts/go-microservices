package middlewareHandler

import (
	"githib.coom/jetsadawwts/go-microservices/config"
	"githib.coom/jetsadawwts/go-microservices/modules/middleware/middlewareUsecase"
)

type (
	MiddlewareHandlerService interface{}

	middlewareHandler struct {
		cfg *config.Config
		middlewareUsecase middlewareUsecase.MiddlewareUsecaseHandler 
	}
)

func NewMiddlewareHandler(cfg *config.Config, middlewareUsecase middlewareUsecase.MiddlewareUsecaseHandler) MiddlewareHandlerService {
	return &middlewareHandler{
		cfg: cfg,
		middlewareUsecase: middlewareUsecase,
	}
}
