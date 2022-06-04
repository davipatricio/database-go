package database

import (
	"encoding/json"
	"os"
)

// Represents a database
type Database struct {
	Path   string
	Data   map[string]interface{}
	Loaded bool
}

// Initializes a instance of a database
// Database should be loaded before using it!
//  db := New("database.json")
//  db.Load()
func New(path string) *Database {
	return &Database{
		Path:   path,
		Data:   make(map[string]interface{}),
		Loaded: false,
	}
}

// Loads the database from the database file
func (db *Database) Load() error {
	file, err := os.Open(db.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&db.Data)
	if err != nil {
		return err
	}

	return nil
}

// Saves the database to the database file
// You should call this method everytime you make changes
func (db *Database) Save() error {
	file, err := os.Create(db.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&db.Data)
	if err != nil {
		return err
	}

	return nil
}

// Returns the value of a key
//   val := db.Get("key")
//   if val == nil {
//     fmt.Println("key not found")
//   }
//
//   fmt.Println(val)
func (db *Database) Get(key string) interface{} {
	return db.Data[key]
}

// Sets the value of a key
//   db.Set("key", "value")
//   db.Set("ints", 12)
//   db.Set("maps", map[string]string{"a": "b"})
//   db.Set("slices", []string{"a", "b"})
//   db.Save()
func (db *Database) Set(key string, value interface{}) {
	db.Data[key] = value
}

// Deletes a key from the database
//   db.Delete("key")
//   db.Save()
func (db *Database) Delete(key string) {
	delete(db.Data, key)
}

// Whether the key exists in the database
//  if db.Has("key") {
//    fmt.Println("key exists")
//  }
func (db *Database) Has(key string) bool {
	_, ok := db.Data[key]
	return ok
}

// Returns a slice of all keys in the database
//   keys := db.Keys()
//   for _, key := range keys {
//     fmt.Println(key)
//   }
func (db *Database) Keys() []string {
	keys := make([]string, 0, len(db.Data))
	for key := range db.Data {
		keys = append(keys, key)
	}
	return keys
}

// Returns a slice of all values in the database
//   values := db.Values()
//   for _, value := range values {
//     fmt.Println(value)
//   }
func (db *Database) Values() []interface{} {
	values := make([]interface{}, 0, len(db.Data))
	for _, value := range db.Data {
		values = append(values, value)
	}
	return values
}

// Deletes all keys from the database
//   db.Clear()
//   db.Save()
func (db *Database) Clear() {
	db.Data = make(map[string]interface{})
}

// Returns the number of keys in the database
//   fmt.Println(db.Len())
func (db *Database) Size() int {
	return len(db.Data)
}

// Whether the database is empty
//   if db.IsEmpty() {
//     fmt.Println("database is empty")
//   }
func (db *Database) IsEmpty() bool {
	return len(db.Data) == 0
}

// Whether the database is loaded
//   if db.IsLoaded() {
//     fmt.Println("database is loaded")
//   } else {
//     db.Load()
//     defer db.Save()
//     db.Set("key", "value")
//   }
func (db *Database) IsLoaded() bool {
	return db.Loaded
}
