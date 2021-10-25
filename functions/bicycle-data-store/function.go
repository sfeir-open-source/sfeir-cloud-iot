// Package p contains a Pub/Sub Cloud Function.
package bicycle

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

type Metric struct {
	Rpm         int    `json:"rpm"`
	Revolutions int    `json:"revolutions"`
	Time        string `json:"time"`
}

// HelloPubSub consumes a Pub/Sub message.
func Bicycle(ctx context.Context, m PubSubMessage) error {
	log.Println(string(m.Data))
	dbClient := newFirebaseApp(ctx)
	addDocument(ctx, *dbClient, m.Data)
	return nil
}

func newFirebaseApp(ctx context.Context) *db.Client {
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

	return client

}

func addDocument(ctx context.Context, dbClient db.Client, message []byte) {
	metric := &Metric{}
	json.Unmarshal(message, &metric)

	// retrieve a byte slice from bytes.Buffer
	err := dbClient.NewRef(fmt.Sprintf("bicycle_data/%s", metric.Time)).Set(ctx, metric)
	if err != nil {
		log.Fatalln("Error set message", err)
	}

}
