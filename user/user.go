package User

import (
	"Database_Go/collection"
	"Database_Go/database"
	"encoding/json"
	"os"
)

type User struct {
    Username string
    Email    string
	Password string
    DB map[string]*database.Database
}

type UserDB struct {
    Data map[string]User
}

func NewHashDB() *UserDB {
    return &UserDB{Data: make(map[string]User)}
}

func (db *UserDB) AddUser(key string, user User) {
    db.Data[key] = user
}

func (db *UserDB) AddDatabase(userKey, dbName string) {
    user, exists := db.Data[userKey]
    if !exists {
        return
    }

    if user.DB == nil {
        user.DB = make(map[string]*database.Database)
    }

    user.DB[dbName] = &database.Database{
        Data: make(map[string]*collection.Collection),
    }

    db.Data[userKey] = user
}

func (db *UserDB) DeleteDatabase(userKey, dbName string) {
    user, exists := db.Data[userKey]
    if !exists {
        return
    }

    if user.DB != nil {
        delete(user.DB, dbName)
    }

    db.Data[userKey] = user
}

func (db *UserDB) GetUser(key string) (User, bool) {
    user, exists := db.Data[key]
    return user, exists
}

func (db *UserDB) SaveToFile(filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    return encoder.Encode(db.Data)
}

func (db *UserDB) LoadFromFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    return decoder.Decode(&db.Data)
}