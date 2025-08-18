package middlewareRepository

import (
	"context"
	"errors"
	"log"
	"time"

	authPb "github.com/jetsadawwts/go-microservices/modules/auth/authPb"
	"github.com/jetsadawwts/go-microservices/pkg/grpcconn"
	"github.com/jetsadawwts/go-microservices/pkg/jwtauth"
)

type (
	MiddlewareRepositoryHandler interface {
		AccessTokenSearch(pctx context.Context, grpcUrl,  accessToken string) error 
		RoleCount(pctx context.Context, grpcUrl string) (int64, error)
	}

	middlewareRepository struct {}
)


func NewMiddlewareRepository() MiddlewareRepositoryHandler {
	return &middlewareRepository{}
}

func (r *middlewareRepository) AccessTokenSearch(pctx context.Context, grpcUrl,  accessToken string) error {

	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()
	
	jwtauth.SetApiKeyInContext(&ctx)	
	conn, err := grpcconn.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %v", err.Error())
		return errors.New("error: gRPC connection failed")
	}

	result, err := conn.Auth().AccessTokenSearch(ctx, &authPb.AccessTokenSearchReq{
		AccessToken: accessToken,
	})
	if err != nil {
		log.Printf("Error: CredentialSearch failed: %s", err.Error())
		return errors.New("error: email or password is incorrect")
	}

	if result == nil {
		log.Printf("Error: access token is invalid")
		return errors.New("error: access token is invalid")
	}

	if !result.IsValid {
		log.Printf("Error: access token is invalid")
		return errors.New("error: access token. is invalid")
	}

	return nil
}

func (r *middlewareRepository) RoleCount(pctx context.Context, grpcUrl string) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	jwtauth.SetApiKeyInContext(&ctx)	
	conn, err := grpcconn.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %v", err.Error())
		return -1, errors.New("error: gRPC connection failed")
	}

	result, err := conn.Auth().RolesCount(ctx, &authPb.RolesCountReq{})
	if err != nil {
		log.Printf("Error: CredentialSearch failed: %s", err.Error())
		return -1, errors.New("error: email or password is incorrect")
	}

	if result == nil {
		log.Printf("Error: roles count invalid")
		return -1, errors.New("error: roles count invalid")
	}

	return result.Count, nil
}