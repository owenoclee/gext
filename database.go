package main

import (
	"database/sql"
	"log"
)

type dbConfig struct {
	DataSourceName string
}

func initDB() *sql.DB {
	var config dbConfig
	loadConfig(&config, "cfg/database.json")
	db := openDB(config)
	createSchema(db)
	return db
}

func openDB(config dbConfig) *sql.DB {
	db, err := sql.Open("mysql", config.DataSourceName+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func createSchema(db *sql.DB) {
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS posts (
			id			INT	UNSIGNED	NOT NULL	AUTO_INCREMENT,
			reply_to	INT UNSIGNED	NULL,
			board		VARCHAR(16)		NULL,
    		subject		VARCHAR(32)		NULL,
			body		VARCHAR(4096)	NULL,
			created_at	TIMESTAMP		NOT NULL	DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			FOREIGN KEY (reply_to) REFERENCES posts(id)
		)`,
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`CREATE OR REPLACE VIEW latest_threads AS
		SELECT threads.id AS thread_id, COALESCE(MAX(replies.created_at), MAX(threads.created_at)) AS bumped_at, threads.board FROM posts AS threads
		LEFT JOIN posts AS replies ON replies.reply_to = threads.id
		WHERE threads.reply_to IS NULL
		GROUP BY thread_id`,
	)
	if err != nil {
		log.Fatal(err)
	}
}
