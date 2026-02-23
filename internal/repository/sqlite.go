package repository

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func Open(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(`PRAGMA foreign_keys = ON`); err != nil {
		return nil, err
	}

	return db, nil
}

func InitDB(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS guests (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tg_id INTEGER UNIQUE,
		name TEXT,
		username TEXT,
		username_fio TEXT,
		attending BOOLEAN,
		plus_one BOOLEAN,
		plus_one_fio TEXT,
		meal TEXT,
		drinks BOOLEAN,
		drink_type TEXT,
		table_id INTEGER,
		state TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS tables (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		capacity INTEGER
	);
	`

	_, err := db.Exec(schema)
	return err
}
