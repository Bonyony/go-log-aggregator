package aggregator

import (
	"fmt"

	"cloud.google.com/go/pubsub"
)

func AggragateLogs() {
	fmt.Println("Aggregator Starting...")
}

func StartSubscriber(client *pubsub.Client) {
	fmt.Println("Received client: ", client)
}
