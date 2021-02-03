package main

import (
	"log"
	"sort"

	"github.com/romnn/flatbson"
	"go.mongodb.org/mongo-driver/bson"
)

func getLongestKey() string {
	input := bson.D{
		{Key: "user", Value: bson.D{{Key: "email", Value: "test"}}},
		{Key: "metadata", Value: bson.D{{Key: "city", Value: bson.D{{Key: "name", Value: "Berlin"}}}}},
	}
	// Flatten the document
	flattened, err := flatbson.Flattened(input, ".")
	if err != nil {
		log.Fatalf("Failed to flatten: %v", err)
	}
	log.Printf("Flattened: %v\n", flattened)

	// Get the longest key
	keys := make([]string, 0, len(flattened))
	for k := range flattened.Map() {
		keys = append(keys, k)
	}
	sort.Sort(sort.StringSlice(keys))
	return keys[0]
}

func main() {
	log.Println(getLongestKey())
}
