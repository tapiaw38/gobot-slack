package db

import (
	"bot-slack/env"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Connect to the database
func Connect() *sql.DB {

	//load envs
	env.LoadEnv()

	client := "user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=disable"

	db, err := sql.Open(
		"postgres",
		client,
	)

	if err != nil {
		log.Fatal(err)
	}

	// ... use the db
	fmt.Println("Connected to the database successfully")

	return db
}
