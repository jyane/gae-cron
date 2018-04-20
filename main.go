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

func publish(ctx context.Context, client *pubsub.Client, topic, msg string) (id string, err error) {
	t := client.Topic(topic)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	return result.Get(ctx)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func publishHandler(w http.ResponseWriter, r *http.Request) {
	if topic, found := strToTopic[r.URL.Query().Get("topic")]; found {
		ctx := appengine.NewContext(r)
		client, err := pubsub.NewClient(ctx, os.Getenv("PROJECT_ID"))
		if err != nil {
			log.Fatalf("Could not create pubsub client: %v", err)
		}
		id, err := publish(ctx, client, string(topic), "a")
		if err != nil {
			log.Fatalf("Could not publish: %v", err)
		}
		fmt.Fprintf(w, "hello %v", id)
	} else {
		fmt.Fprintf(w, "bad topic")
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/publish", publishHandler)
	appengine.Main()
}
