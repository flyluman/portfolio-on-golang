package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	db_init_q := `CREATE TABLE IF NOT EXISTS log (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ip TEXT, isp TEXT, city TEXT, country TEXT,
		date TEXT, path TEXT, useragent TEXT
	);
		CREATE TABLE IF NOT EXISTS foreign_log (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ip TEXT, isp TEXT, city TEXT, country TEXT,
		date TEXT, path TEXT, useragent TEXT
	);
		CREATE TABLE IF NOT EXISTS msg (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ip TEXT, isp TEXT, city TEXT, country TEXT,
		date TEXT, useragent TEXT, name TEXT, email TEXT, msg TEXT
	);`

	_, err = DB.Exec(db_init_q)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, db_init_q)
	}

	log.Println("✅ SQLite DB Initialized")
}
