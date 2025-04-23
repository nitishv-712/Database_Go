package main

import (
	"Database_Go/data"
	"fmt"
)

func main() {
	db := data.NewJSONDatabase()

	// Add JSON-like user data
	db.AddJSONRecord("user1", map[string]interface{}{
		"name":   "SampurnaKart",
		"active": true,
		"city":   "Jalandhar",
	})

	// Save to file
	db.SaveToJSONFile("db.json")

	// Load from file
	newDB := data.NewJSONDatabase()
	newDB.LoadFromJSONFile("db.json")

	// Fetch and print
	if user, ok := newDB.GetJSONRecord("user1"); ok {
		fmt.Println("User1:", user)
	}
}
