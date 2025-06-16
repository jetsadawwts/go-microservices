package migration

import (
	"context"
	"log"

	"github.com/jetsadawwts/go-microservices/config"
	"github.com/jetsadawwts/go-microservices/modules/product"
	"github.com/jetsadawwts/go-microservices/pkg/database"
	"github.com/jetsadawwts/go-microservices/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func productDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("product_db")
}

func ProductMigrate(pctx context.Context, cfg *config.Config) {
	db := productDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("products")

	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"_id", 1}}},
		{Keys: bson.D{{"title", 1}}},
	})

	for _, index := range indexs {
		log.Printf("Index: %s", index)
	}

	documents := func() []any {
		roles := []*product.Product{
			{
				Title:       "Diamond Sword",
				Price:       1000,
				ImageUrl:    "https://i.imgur.com/1Y8tQZM.png",
				UsageStatus: true,
				Damage:      100,
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
			},
			{
				Title:       "Iron Sword",
				Price:       500,
				ImageUrl:    "https://i.imgur.com/1Y8tQZM.png",
				UsageStatus: true,
				Damage:      50,
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
			},
			{
				Title:       "Wooden Sword",
				Price:       100,
				ImageUrl:    "https://i.imgur.com/1Y8tQZM.png",
				UsageStatus: true,
				Damage:      20,
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
			},
		}
		docs := make([]any, 0)
		for _, r := range roles {
			docs = append(docs, r)
		}
		return docs
	}()

	results, err := col.InsertMany(pctx, documents, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate product completed: ", results)
}
