package productRepository

import (
	"context"
	"errors"
	"log"
	"time"
	"github.com/jetsadawwts/go-microservices/modules/product"
	"github.com/jetsadawwts/go-microservices/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	ProductRepositoryService interface {
		IsUniqueProduct(pctx context.Context, title string) bool
		InsertOneProduct(pctx context.Context, req *product.Product) (primitive.ObjectID, error)
		FindOneProduct(pctx context.Context, productId string) (*product.Product, error)
		FindManyProducts(pctx context.Context,  filter primitive.D, opts []*options.FindOptions) ([]*product.ProductShowCase, error)
		CountProducts(pctx context.Context, filter primitive.D) (int64, error)
		UpdateOneProduct(pctx context.Context, productId string, req primitive.M) error
		EnableOrDisableProduct(pctx context.Context, productId string, isActive bool) error
	}

	productRepository struct {
		db *mongo.Client
	}
)

func NewProductRepository(db *mongo.Client) ProductRepositoryService {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) productDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("product_db")
}

func (r *productRepository) IsUniqueProduct(pctx context.Context, title string) bool {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.productDbConn(ctx)
	col := db.Collection("products")

	product := new(product.Product)
	if err := col.FindOne(ctx, bson.M{"title": title}).Decode(product); err != nil {
		log.Printf("Error: IsUniqueProduct: %s", err.Error())
		return true
	}

	return false
}

func (r *productRepository) InsertOneProduct(pctx context.Context, req *product.Product) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.productDbConn(ctx)
	col := db.Collection("products")

	productId, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOneProduct: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert one product failed")
	}

	return productId.InsertedID.(primitive.ObjectID), nil
} 

func (r *productRepository) FindOneProduct(pctx context.Context, productId string) (*product.Product, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.productDbConn(ctx)
	col := db.Collection("products")

	product := new(product.Product)
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(productId)}).Decode(product); err != nil {
		log.Printf("Error: FindOneProduct failed: %s", err.Error())
		return nil, errors.New("error: product not found")
	}

	return product, nil
}

func (r *productRepository) FindManyProducts(pctx context.Context, filter primitive.D, opts []*options.FindOptions) ([]*product.ProductShowCase, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.productDbConn(ctx)
	col := db.Collection("products")

	cursor, err := col.Find(ctx, filter, opts...)
	if err != nil {
		log.Printf("Error: FindManyProducts failed: %s", err.Error())
		return make([]*product.ProductShowCase, 0), errors.New("error: find many product failed")
	}

	results := make([]*product.ProductShowCase, 0)
	for cursor.Next(ctx) {
		result := new(product.Product)
		if err := cursor.Decode(result); err != nil {
			log.Printf("Error: FindManyProducts failed: %s", err.Error())
			return make([]*product.ProductShowCase, 0), errors.New("error: find many products failed")
		}
		results = append(results, &product.ProductShowCase{
			ProductId: "product:" + result.Id.Hex(),
			Title: result.Title,
			Price: result.Price,
			Damage: result.Damage,
			ImageUrl: result.ImageUrl,
		})
	}

	return results, nil
}

func (r *productRepository) CountProducts(pctx context.Context, filter primitive.D) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.productDbConn(ctx)
	col := db.Collection("products")

	count, err := col.CountDocuments(ctx, filter)
	if err != nil {
		log.Printf("Error: CountProducts failed: %s", err.Error())
		return -1, errors.New("error: count product failed")
	}
	
	return count, nil
}

func (r *productRepository) UpdateOneProduct(pctx context.Context, productId string, req primitive.M) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.productDbConn(ctx)
	col := db.Collection("products")

	result, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(productId)}, bson.M{"$set": req})
	if err != nil {
		log.Printf("Error: UpdateOneProduct failed: %s", err.Error())
		return errors.New("error: update product failed")
	}
	log.Printf("UpdateOneProduct result: %v", result.ModifiedCount)
	return nil
}

func (r *productRepository) EnableOrDisableProduct(pctx context.Context, productId string, isActive bool) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.productDbConn(ctx)
	col := db.Collection("products")

	result, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(productId)}, bson.M{"$set": bson.M{"usage_status": isActive}})
	if err != nil {
		log.Printf("Error: EnableOrDisableProduct failed: %s", err.Error())
		return errors.New("error: enable or disable product failed")
	}
	log.Printf("EnableOrDisableProduct result: %v", result.ModifiedCount)
	return nil
}



