package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to mysql storage")
	err = db.Ping()
	if err != nil {
		log.Fatal("Something went wrong on database connection", err)
	}
	return &MySQLStorage{
		db,
	}
}

func (s *MySQLStorage) Init() (*sql.DB, error) {
	//initialize the tables
	return s.db, nil

}
