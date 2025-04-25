package main

import (
	"Database_Go/collection"
	"Database_Go/database"
	User "Database_Go/user"
	"fmt"
)

func main() {
	// Create UserDB
	userDB := User.NewHashDB()

	// Add user
	user := User.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
		DB:       make(map[string]*database.Database),
	}
	userDB.AddUser("testuser", user)

	// Add database to user
	dbUser, _ := userDB.GetUser("testuser")
	dbUser.DB["AppDB"] = &database.Database{
		Data: make(map[string]*collection.Collection),
	}
	userDB.AddUser("testuser", dbUser)

	// Add collection to user's database
	collection1 := &collection.Collection{}
	dbUser.DB["AppDB"].Data["Users"] = collection1
	userDB.AddUser("testuser", dbUser)

	// Add JSON data to collection
	json1 := collection.JSONData{
		Data: map[string]interface{}{
			"id":    "001",
			"name":  "Alice",
			"value": 42,
		},
	}
	dbUser.DB["AppDB"].Data["Users"].Add(json1)

	// Fetch and print collection data
	collection := dbUser.DB["AppDB"].Data["Users"]
	for i := 0; i < collection.Len(); i++ {
		if data, ok := collection.Get(i); ok {
			fmt.Printf("Data %d: %+v\n", i, data.Data)
		}
	}

	// Delete the collection
	dbUser.DB["AppDB"].DropCollection("Users")

	// Drop entire database
	dbUser.DB["AppDB"].DropDB()
	delete(dbUser.DB, "AppDB")

	// Delete user
	delete(userDB.Data, "testuser")

	fmt.Println("All operations completed.")
}
