package productRepository

import "go.mongodb.org/mongo-driver/mongo"


type (
	ProductRepositoryService interface {}

	productRepository struct {
		db *mongo.Client
	}
)


func NewProductRepository(db *mongo.Client) ProductRepositoryService {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) productDbConn() *mongo.Database {
	return r.db.Database("product-db")
}