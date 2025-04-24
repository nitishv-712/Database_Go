package main

import (
	"Database_Go/collection"
	"fmt"
)

func main() {
	coll := collection.Collection{}

	// Create and add first dataset
	ds1 := collection.JSONData{Data: map[string]interface{}{
		"id":    "001",
		"name":  "temperature",
		"value": 36.6,
	}}

	coll.Add(ds1)

	// Create and add second dataset
	ds2 := collection.JSONData{Data: map[string]interface{}{
		"id":    "002",
		"name":  "humidity",
		"value": 80,
	}}

	coll.Add(ds2)

	// Find by key-value
	results := coll.FindByKeyValue("name", "humidity")
	fmt.Println("Results:", results)

	// Get dataset
	if ds, ok := coll.Get(1); ok {
		fmt.Println("Got dataset:", ds)
	}

	// Remove dataset
	// coll.Remove(0)
	fmt.Println("After remove:", coll)
}
