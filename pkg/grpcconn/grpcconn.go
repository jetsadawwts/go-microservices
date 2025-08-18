package grpcconn

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/jetsadawwts/go-microservices/config"
	"github.com/jetsadawwts/go-microservices/pkg/jwtauth"

	authPb "github.com/jetsadawwts/go-microservices/modules/auth/authPb"
	inventoryPb "github.com/jetsadawwts/go-microservices/modules/inventory/inventoryPb"
	productPb "github.com/jetsadawwts/go-microservices/modules/product/productPb"
	userPb "github.com/jetsadawwts/go-microservices/modules/user/userPb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type (
	GrpcClientFactoryHandle interface {
		Auth() authPb.AuthGrpcServiceClient
		User() userPb.UserGrpcServiceClient
		Product() productPb.ProductGrpcServiceClient
		Inventory() inventoryPb.InventoryGrpcServiceClient
	}

	grpcClientFactory struct {
		client *grpc.ClientConn
	}

	grpcAuth struct {
		secretKey string
	}
)

func (g *grpcAuth) unaryAuthorization(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("Error: Metadata. not found")
		return nil, errors.New("error: metadata not found")
	}

	authHandler, ok := md["auth"]
	if !ok {
		log.Printf("Error: Metadata not found")
		return nil, errors.New("error: metadata not found")
	}

	if (len(authHandler) == 0) {
		log.Printf("Error: Metadata not found")
		return nil, errors.New("error: metadata not found")
	}

	claims, err := jwtauth.ParseToken(g.secretKey, string(authHandler[0]))
	if err != nil {
		log.Printf("Error: Parse token failed: %s", err.Error())
		return nil, errors.New("error: token is invalid")
	}

	log.Printf("claims: %v", claims)

	return handler(ctx, req)
}

func (g *grpcClientFactory) Auth() authPb.AuthGrpcServiceClient {
	return authPb.NewAuthGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) User() userPb.UserGrpcServiceClient {
	return userPb.NewUserGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Product() productPb.ProductGrpcServiceClient {
	return productPb.NewProductGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Inventory() inventoryPb.InventoryGrpcServiceClient {
	return inventoryPb.NewInventoryGrpcServiceClient(g.client)
}

func NewGrpcClient(host string) (GrpcClientFactoryHandle, error) {
	opts := make([]grpc.DialOption, 0)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	clientConn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Printf("Error: Grpc client connection failed: %s", err.Error())
		return nil, errors.New("error: grpc client connection failed")
	}
	return &grpcClientFactory{
		client: clientConn,
	}, nil
}

func NewGrpcServer(cfg *config.Jwt, host string) (*grpc.Server, net.Listener) {
	opts := make([]grpc.ServerOption, 0)

	grpcAuth := &grpcAuth{
		secretKey: cfg.ApiSceretKey,
	}

	opts = append(opts, grpc.UnaryInterceptor(grpcAuth.unaryAuthorization))

	grpcServer := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", host)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return grpcServer, lis
}
