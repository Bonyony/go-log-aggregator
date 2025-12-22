package main

import (
	"fmt"

	"go-log-aggregator/internal/producer"
)

func main() {
	fmt.Println("This new project will be awesome!!!")
	producer.ProduceLog()
}