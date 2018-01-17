package db

import (
	"database/sql"
	"grafana-simplejson-datacreator/common/config"
	"log"

	"github.com/go-sql-driver/mysql"
)

func GetDBData(dbName, command string) (*sql.Rows, error) {

	db, err := open(dbName)
	defer close(db)

	if err != nil {
		log.Printf("an error occurred while opening db connection\r\ndb: %s\r\n%s", config.Config.DB_Address, err)
		return nil, err
	}

	results, err := db.Query(command)

	if err != nil {
		log.Printf("an error occurred while quering data\r\ndb: %s\r\n%s", config.Config.DB_Address, err)
		return nil, err
	}

	return results, nil
}

func open(dbName string) (*sql.DB, error) {

	dbConfig := mysql.Config{
		Net:    "tcp",
		Addr:   config.Config.DB_Address,
		User:   config.Config.DB_User,
		Passwd: config.Config.DB_Password,
		DBName: dbName,
	}

	db, err := sql.Open("mysql", dbConfig.FormatDSN())

	return db, err
}

func close(db *sql.DB) {
	if db != nil {
		db.Close()
	}
}
