package inventoryRepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jetsadawwts/go-microservices/modules/inventory"
	productPb "github.com/jetsadawwts/go-microservices/modules/product/productPb"
	"github.com/jetsadawwts/go-microservices/pkg/grpcconn"
	"github.com/jetsadawwts/go-microservices/pkg/jwtauth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	InventoryRepositoryService interface {
		FindProductsInIds(pctx context.Context, grpcUrl string, req *productPb.FindProductsInIdsReq) (*productPb.FindProductsInIdsRes, error)
		FindUserProducts(pctx context.Context, filter primitive.D, opts []*options.FindOptions) ([]*inventory.Inventory, error) 
		CountUserProducts(pctx context.Context, userId string) (int64, error)
	}

	inventoryRepository struct {
		db *mongo.Client
	}
)

func NewInventoryRepository(db *mongo.Client) InventoryRepositoryService {
	return &inventoryRepository{
		db: db,
	}
}

func (r *inventoryRepository) inventoryDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("inventory_db")
}

func (r *inventoryRepository) FindProductsInIds(pctx context.Context, grpcUrl string, req *productPb.FindProductsInIdsReq) (*productPb.FindProductsInIdsRes, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	jwtauth.SetApiKeyInContext(&ctx)
	conn, err := grpcconn.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %s", err.Error())
		return nil, errors.New("error: gRPC connection failed")
	}

	result, err := conn.Product().FindProductsInIds(ctx, req)


	if err != nil {
		log.Printf("Error: FindProductsInIds failed: %s", err.Error())
		return nil, errors.New("error: products not found")
	}

	if result == nil {
		log.Printf("Error: FindProductsInIds failed: %s", err.Error())
		return nil, errors.New("error: products not found")
	}

	if len(result.Products) == 0 {
		log.Printf("Error: FindProductsInIds failed: %s", err.Error())
		return nil, errors.New("error: products not found")
	}

	return result, nil
}

func (r *inventoryRepository) FindUserProducts(pctx context.Context, filter primitive.D, opts []*options.FindOptions) ([]*inventory.Inventory, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.inventoryDbConn(ctx)
	col := db.Collection("users_inventory")

	cursors, err := col.Find(ctx, filter, opts...)
	if err != nil {
		log.Printf("Error: FindUserProducts failed: %s", err.Error())
		return nil, errors.New("error: user products not found")
	}

	results := make([]*inventory.Inventory, 0)
	for cursors.Next(ctx) {
		result := new(inventory.Inventory)
		if err := cursors.Decode(result); err != nil {
			log.Printf("Error: FindUserProducts failed: %s", err.Error())
			return nil, errors.New("error: user products not found")
		}

		results = append(results, result)
	}

	return results, nil
}


func (r *inventoryRepository) CountUserProducts(pctx context.Context, userId string) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.inventoryDbConn(ctx)
	col := db.Collection("users_inventory")

	count, err := col.CountDocuments(ctx, bson.M{"user_id": userId})
	if err != nil {
		log.Printf("Error: CountUserProducts failed: %s", err.Error())
		return -1, errors.New("error: count user products failed")
	}
	
	return count, nil
}