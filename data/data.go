package data

import (
	"encoding/json"
	"os"
)

type JSONData struct {
	Data map[string]interface{}
}

func NewJSONDatabase() *JSONData {
	return &JSONData{Data: make(map[string]interface{})}
}

// Add a user (or any record) as a JSON object
func (db *JSONData) AddJSONRecord(key string, jsonData map[string]interface{}) {
	db.Data[key] = jsonData
}

// Get a JSON record by key
func (db *JSONData) GetJSONRecord(key string) (map[string]interface{}, bool) {
	data, exists := db.Data[key]
	if !exists {
		return nil, false
	}
	jsonObj, ok := data.(map[string]interface{})
	return jsonObj, ok
}

// Save the entire database to a JSON file
func (db *JSONData) SaveToJSONFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(db.Data)
}

func (db *JSONData) LoadFromJSONFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&db.Data)
}
