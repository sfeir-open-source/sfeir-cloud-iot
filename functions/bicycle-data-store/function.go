// Package p contains a Pub/Sub Cloud Function.
package bicycle

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// HelloPubSub consumes a Pub/Sub message.
func Bicycle(ctx context.Context, m PubSubMessage) error {
	log.Println(string(m.Data))
	logToFirebase(ctx)
	return nil
}

func logToFirebase(ctx context.Context) {
	databaseName := os.Getenv("DATABASE_NAME")
	conf := &firebase.Config{
		DatabaseURL: fmt.Sprintf("https://%s.firebaseio.com/", databaseName),
	}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	// As an admin, the app has access to read and write all data, regradless of Security Rules
	ref := client.NewRef("/bicycle_data/radius")
	var data map[string]interface{}
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	fmt.Println(data)
}
