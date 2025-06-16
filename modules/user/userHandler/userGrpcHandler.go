package userHandler

import (
	"context"

	userPb "github.com/jetsadawwts/go-microservices/modules/user/userPb"
	"github.com/jetsadawwts/go-microservices/modules/user/userUsecase"
)

type (
	userGrpcHandler struct {
		userPb.UnimplementedUserGrpcServiceServer
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserGrpcHandler(userUsecase userUsecase.UserUsecaseService) *userGrpcHandler {
	return &userGrpcHandler{
		userUsecase: userUsecase,
	}
}

func (g *userGrpcHandler) CredentialSearch(ctx context.Context, req *userPb.CredentialSearchReq) (*userPb.UserProfile, error) {
	return nil, nil
}

func (g *userGrpcHandler) FindOneUserProfileToRefresh(ctx context.Context, req *userPb.FindOneUserProfileToRefreshReq) (*userPb.UserProfile, error) {
	return nil, nil
}

func (g *userGrpcHandler) GetUserSavingAccount(ctx context.Context, req *userPb.GetUserSavingAccountReq) (*userPb.GetUserSavingAccountRes, error) {
	return nil, nil
}
