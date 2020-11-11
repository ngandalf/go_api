package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
)

const (
  host     = "172.18.0.1"
  port     = 5432
  user     = "postgres"
  password = "postgres"
  dbname   = "postgres"
)

var db *sql.DB

func DatabaseInit() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	// Create Table cars if not exists
	createCarsTable()
}

func createCarsTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS cars(id serial,manufacturer varchar(20), design varchar(20), style varchar(20), doors int, created_at timestamp default NULL, updated_at timestamp default NULL, constraint pk primary key(id))")

	if err != nil {
		log.Fatal(err)
	}
}

// Getter for db var
func Db() *sql.DB {
	return db
}
