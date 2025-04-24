package database

import (
	"Database_Go/collection"
)

type Database struct {
	Data map[string]*collection.Collection
}

func CreateNewDB() *Database {
	return &Database{Data: make(map[string]*collection.Collection)}
}

func (db *Database) DropDB() bool {
	for _, col := range db.Data {
		col.Data = nil
	}
	db.Data = nil 
	return true
}


func (db *Database) AddCollection(CollectionName string) (*Database, bool) {
	if db.Data == nil {
		db.Data = make(map[string]*collection.Collection)
	}

	if _, exists := db.Data[CollectionName]; !exists {
		db.Data[CollectionName] = &collection.Collection{}
		return db, true
	}

	return db, false
}

func (db *Database) DropCollection(CollectionName string) (*Database, bool){

	if _, exists := db.Data[CollectionName]; exists {
		db.Data[CollectionName].DeleteAll()
		delete(db.Data,CollectionName)
		return db, true
	}
	return nil, false
}
