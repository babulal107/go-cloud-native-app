package migration

import (
	"database/sql"
	"log"
)

func CreateUserTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
	 	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`
	if _, err := db.Exec(query); err != nil {
		log.Fatalf("failed to create users table error: %v\n", err)
	}
}
