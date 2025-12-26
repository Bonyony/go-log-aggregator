package main

import (
	"context"
	"fmt"
	"log"

	"go-log-aggregator/internal/producer"
	mypubsub "go-log-aggregator/internal/pubsub"
)

func main() {
	fmt.Println("This new project will be awesome!!!")
	/* 
		Creates Pub / Sub client
		and connects to my Google Account
	*/
	ctx := context.Background()

	client, err := mypubsub.NewClient(ctx, "go-log-aggregator")
	if (err != nil) {
		log.Fatal("Failed to create pub/sub client: ", err)
	}
	defer client.Close()
	/* 
		Produces a set ammount of logs 
		and Uploads to Pub / Sub
	*/
	producer.ProduceLog()

	/* 
		Receives logs from Pub / Sub
	*/
	// aggregator here

	/* 
		Parses logs
	*/
	// parser here

	// Make templates here?
}