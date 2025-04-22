package database

import (
	"encoding/json"
	"os"
)

type User struct {
    Username string
    Email    string
	Password string
}

type HashDB struct {
    data map[string]User
}

func NewHashDB() *HashDB {
    return &HashDB{data: make(map[string]User)}
}

func (db *HashDB) AddUser(key string, user User) {
    db.data[key] = user
}

func (db *HashDB) GetUser(key string) (User, bool) {
    user, exists := db.data[key]
    return user, exists
}

func (db *HashDB) SaveToFile(filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    return encoder.Encode(db.data)
}

func (db *HashDB) LoadFromFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    return decoder.Decode(&db.data)
}