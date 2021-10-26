package db

import (
	"context"
	"fmt"
	"log"
	"os"

	// Autoload .env file
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
)

var Db *db.Client

func init() {
	Db = databaseConnection()

	if Db == nil {
		log.Printf("failed to connect to Database")
		os.Exit(1)
	}
}

func databaseConnection() *db.Client {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: fmt.Sprintf("https://%s.firebaseio.com/", os.Getenv("DATABASE_NAME")),
	}
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	return client
}
