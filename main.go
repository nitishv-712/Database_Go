package main

import (
	User "Database_Go/user"
	"fmt"
)

func main() {
    // db := User.NewHashDB()

    // db.AddUser("john", User.User{Username: "john", Email: "john@example.com"})
    // db.AddUser("alice", User.User{Username: "alice", Email: "alice@example.com"})

    // db.SaveToFile("db.json")

    // // Create new instance and load data
    newDB := User.NewHashDB()
    newDB.LoadFromFile("db.json")
	for key, user := range newDB.Data {
        fmt.Printf("Key: %s, Username: %s, Email: %s\n", key, user.Username, user.Email)
    }
    // if user, found := newDB.GetUser("alice"); found {
    //     fmt.Println("Found user:", user)
    // }
	
}
