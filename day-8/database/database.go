package database

import (
	"sync"

	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
	// we use sync.Once for make sure we create connection only once
	once sync.Once
)

// CreateConnection is a function for creating new connection with database
func CreateConnection() {

	conf := dbConfig{
		User: "root",
		Pass: "123456",
		Host: "127.0.0.1",
		Port: "3306",
		Name: "alterra-agmc-day-2",
	}

	mysql := mysqlConfig{dbConfig: conf}
	// if you use postgres, you can uncomment code bellow

	//postgres := postgresqlConfig{dbConfig: conf}

	once.Do(func() {
		mysql.Connect()
		//postgres.Connect()
	})
}

// GetConnection is a faction for return connection or return value dbConn
// because we set var dbConn is private
func GetConnection() *gorm.DB {
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}
