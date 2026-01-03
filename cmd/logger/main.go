package main

import (
	"context"
	"fmt"
	"log"

	"go-log-aggregator/internal/aggregator"
	"go-log-aggregator/internal/producer"
	mypubsub "go-log-aggregator/internal/pubsub"

	"google.golang.org/api/iterator"
)

func main() {
	fmt.Println("This new project will be awesome!!!")
	/*
		Creates Pub / Sub client
		and connects to my Google Account
	*/

	ctx := context.Background()
	client, err := mypubsub.NewClient(ctx, "go-log-aggregator")
	if err != nil {
		log.Fatal("Failed to create pub/sub client: ", err)
	}
	defer client.Close()

	// check connectivity
	fmt.Println("Checking GCP connection...")
	it := client.Topics(ctx)
	for {
		topic, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error listing topics: %v", err)
		}
		fmt.Printf("Found Topics: %s\n", topic.ID())
	}
	fmt.Println("Connection Successful!")

	/*
		Produces a set ammount of logs
		and Uploads to Pub / Sub
	*/

	producer.ProduceLog()

	/*
		Receives logs from Pub / Sub
	*/

	aggregator.StartSubscriber(client)

	/*
		Parses logs
	*/

	// parser here
	// Make templates here?
}
