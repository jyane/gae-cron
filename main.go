package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
	"google.golang.org/appengine"
)

func publish(client *pubsub.Client, topic, msg string) error {
	ctx := context.Background()
	t := client.Topic(topic)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	id, err := result.Get(ctx)
	if err != nil {
		return err
	}

	log.Printf("Published a message; msg ID: %v", id)
	return nil
}

func main() {
	ctx := context.Background()
	proj := os.Getenv("PROJECT_ID")

	_, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		log.Fatalf("Could not create pubsub client: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})

	http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		publish(client, "minutely-tick", "")
		fmt.Fprintln(w, path)
	})

	appengine.Main()
}
