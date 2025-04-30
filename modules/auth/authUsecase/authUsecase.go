package authUsecase

import "githib.coom/jetsadawwts/go-microservices/modules/auth/authRepository"


type (
	AuthUsecaseService interface {}

	authUsecase struct {
		authRepository authRepository.AuthRepositoryService
	}
)

func NewAuthUsecase(authRepository authRepository.AuthRepositoryService) AuthUsecaseService {
	return &authUsecase{
		authRepository: authRepository,
	}
}
