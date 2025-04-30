package authHandler

import "githib.coom/jetsadawwts/go-microservices/modules/auth/authUsecase"

type (
	authGrpcHandler struct {
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) authGrpcHandler {
	return authGrpcHandler{
		authUsecase: authUsecase,
	}
}
