package authRepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jetsadawwts/go-microservices/modules/auth"
	userPb "github.com/jetsadawwts/go-microservices/modules/user/userPb"
	"github.com/jetsadawwts/go-microservices/pkg/grpcconn"
	"github.com/jetsadawwts/go-microservices/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	AuthRepositoryService interface {
		CredentialSearch(pctx context.Context, grpcUrl string, req *userPb.CredentialSearchReq) (*userPb.UserProfile, error)
		FindOneUserProfileToRefresh(pctx context.Context, grpcUrl string, req *userPb.FindOneUserProfileToRefreshReq) (*userPb.UserProfile, error)
		InsertOneUserCredential(pctx context.Context, req *auth.Credential) (primitive.ObjectID, error)
		FindOneUserCredential(pctx context.Context, credentialId string) (*auth.Credential, error)
		UpdateOneUserCredential(pctx context.Context, credentialId string, req *auth.UpdateRefreshTokenReq) error
		DeleteOneUserCredential(pctx context.Context, credentialId string) (int64, error) 
	}

	authRepository struct {
		db *mongo.Client
	}
)

func NewAuthRepository(db *mongo.Client) AuthRepositoryService {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) authDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("auth_db")
}

func (r *authRepository) CredentialSearch(pctx context.Context, grpcUrl string, req *userPb.CredentialSearchReq) (*userPb.UserProfile, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	conn, err := grpcconn.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %v", err.Error())
		return nil, errors.New("error: gRPC connection failed")
	}

	result, err := conn.User().CredentialSearch(ctx, req)
	if err != nil {
		log.Printf("Error: CredentialSearch failed: %s", err.Error())
		return nil, errors.New("error: email or password is incorrect")
	}

	return result, nil
}

func (r *authRepository) FindOneUserProfileToRefresh(pctx context.Context, grpcUrl string, req *userPb.FindOneUserProfileToRefreshReq) (*userPb.UserProfile, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	conn, err := grpcconn.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %v", err.Error())
		return nil, errors.New("error: gRPC connection failed")
	}

	result, err := conn.User().FindOneUserProfileToRefresh(ctx, req)
	if err != nil {
		log.Printf("Error: FindOneUserProfileToRefresh failed: %s", err.Error())
		return nil, errors.New("error: user profile not found")
	}

	return result, nil
}

func (r *authRepository) InsertOneUserCredential(pctx context.Context, req *auth.Credential) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()
	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	result, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOneUserCredential failed: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert one user credential failed")
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *authRepository) FindOneUserCredential(pctx context.Context, credentialId string) (*auth.Credential, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	result := new(auth.Credential)

	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(credentialId)}).Decode(result); err != nil {
		log.Printf("Error: FindOneUserCredential failed: %s", err.Error())
		return nil, errors.New("error: find one user credential failed")
	}

	return result, nil

}

func (r *authRepository) UpdateOneUserCredential(pctx context.Context, credentialId string, req *auth.UpdateRefreshTokenReq) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	_, err := col.UpdateOne(ctx, bson.M{"refresh_token": utils.ConvertToObjectId(credentialId)}, bson.M{
			"$set": bson.M{
				"user_id": req.UserId,
				"access_token":  req.AccessToken,
				"refresh_token": req.RefreshToken,
				"updated_at":   req.UpdateAt,
			},
		},
	)

	if err != nil {
		log.Printf("Error: UpdateOneUserCredential failed: %s", err.Error())
		return errors.New("error:user credential not found")
	}

	return  nil
}

func (r *authRepository) DeleteOneUserCredential(pctx context.Context, credentialId string) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	result, err := col.DeleteOne(ctx, bson.M{"_id": utils.ConvertToObjectId(credentialId)})

	if err != nil {
		log.Printf("Error: DeleteOneUserCredential failed: %s", err.Error())
		return -1, errors.New("error: delete user credential failed")
	}

	log.Printf("DeleteOneUserCredential: result: %v", result)

	return result.DeletedCount, nil
}