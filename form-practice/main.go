package main

import (
	"context"
	"flag"
	"log"

	"cloud.google.com/go/firestore"
)

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "dares-app-346910"

	// [END firestore_setup_client_create]
	// Override with -project flags
	flag.StringVar(&projectID, "project", projectID, "The Google Cloud Platform project ID.")
	flag.Parse()

	// [START firestore_setup_client_create]
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

func main() {

	ctx := context.Background()
	client := createClient(ctx)

	defer client.Close()

	//add a dare
	_, _, err := client.Collection("dares").Add(ctx, map[string]interface{}{
		"id":       "1",
		"title":    "Sockhands",
		"text":     "socks on your hands",
		"savagery": 3,
	})

	if err != nil {
		log.Fatalf("failed making sockhands: %v", err)
	}

	//add another dare
	_, _, err = client.Collection("dares").Add(ctx, map[string]interface{}{
		"id":       "2",
		"title":    "Tallahassee Bubble Tea",
		"text":     "dogfood in your beer",
		"savagery": 5,
	})

	if err != nil {
		log.Fatalf("failed making tbt: %v", err)
	}

}

//cool, so this works
