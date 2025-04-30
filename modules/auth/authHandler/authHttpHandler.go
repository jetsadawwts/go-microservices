package authHandler

import (
	"githib.coom/jetsadawwts/go-microservices/config"
	"githib.coom/jetsadawwts/go-microservices/modules/auth/authUsecase"
)


type (
	AuthHttpHandlerService interface {}

	authHttpHandler struct {
		cfg *config.Config
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthHttpHandler(cfg *config.Config, authUsecase authUsecase.AuthUsecaseService) AuthHttpHandlerService {
	return &authHttpHandler{
		cfg: cfg,
		authUsecase: authUsecase,
	}
}
