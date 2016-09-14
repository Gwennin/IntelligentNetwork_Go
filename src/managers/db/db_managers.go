package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
	"sync"
)

var _db *gorm.DB = nil
var mutex = &sync.Mutex{}

func InitializeDB() {
	connection_string := os.Getenv("DB_CONNECTION")
	var err error

	mutex.Lock()
	_db, err = gorm.Open("postgres", connection_string)
	_db.LogMode(true)
	mutex.Unlock()

	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() (*gorm.DB, *sync.Mutex) {
	mutex.Lock()
	return _db, mutex
}

func CloseDB() {
	mutex.Lock()
	_db.Close()
	_db = nil
	mutex.Unlock()
}
