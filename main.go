package main

import (
	"context"
	"log"
	"os"

	"github.com/jetsadawwts/go-microservices/config"
	"github.com/jetsadawwts/go-microservices/pkg/database"
	"github.com/jetsadawwts/go-microservices/server"
)

func main() {
	ctx := context.Background()

	//Initiate config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			return ".env"
		}
		return os.Args[1]
	}())

	//Database connection
	db := database.DbConn(ctx, &cfg)
	log.Println(db)

	//Start server
	server.Start(ctx, &cfg, db)

}
