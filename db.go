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
	return &MySQLStorage{
		db,
	}
}

func (s *MySQLStorage) Init() (*sql.DB, error) {
	//initialize the tables
	return s.db, nil

}
