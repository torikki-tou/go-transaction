package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func InitDB(db *sql.DB) {
	var exists bool
	rows, err := db.Query(
		"SELECT EXISTS(SELECT * FROM information_schema.tables WHERE table_name = 'clients')",
	)
	if err != nil {
		panic(err)
	}
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)
	for rows.Next() {
		err = rows.Scan(&exists)
	}
	if exists {
		return
	}

	_, err = db.Exec(
		"CREATE TABLE clients (id varchar(7) primary key, balance int)",
	)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		_, err = db.Exec(
			fmt.Sprintf(
				"INSERT INTO clients (id, balance) VALUES ('%s', '%d')", randStr(7), rand.Intn(90)+10),
		)
		if err != nil {
			panic(err)
		}
	}
}

func SetupDatabaseConnection() *sql.DB {
	connStr := "sslmode=disable user=postgres host=db"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

func CloseDatabaseConnection(db *sql.DB) {
	_ = db.Close()
}
