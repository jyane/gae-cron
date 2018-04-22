package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

"cloud.google.com/go/pubsub"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func publish(ctx context.Context, topic Topic) (string, error) {
	projectId := os.Getenv("PROJECT_ID")
	if projectId == "" {
		log.Errorf(ctx, "Missing project id")
		return "", fmt.Errorf("Missing project id")
	}
	client, err := pubsub.NewClient(ctx, os.Getenv(projectId))
	if err != nil {
		log.Errorf(ctx, "Failed to create the pubsub client %v", err)
		return "", err
	}
	result := client.Topic(string(topic)).Publish(ctx, &pubsub.Message{
		Data: []byte(""),
	})
	id, err := result.Get(ctx)
	if err != nil {
		log.Errorf(ctx, "Failed to publish a message %v", err)
		return "", err
	}
	return id, nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func publishHandler(w http.ResponseWriter, r *http.Request) {
	if topic, found := strToTopic[r.URL.Query().Get("topic")]; found {
		ctx := appengine.NewContext(r)
		if id, err := publish(ctx, topic); err != nil {
			http.Error(w, "internal", http.StatusInternalServerError)
		} else {
			fmt.Fprintf(w, "published: %v", id)
		}
	} else {
		http.Error(w, "bad topic", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/publish", publishHandler)
	appengine.Main()
}
