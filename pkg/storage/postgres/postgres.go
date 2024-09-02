package postgres

import (
	"Kode_test/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //init postgres driver
)

type Storage struct {
	*sql.DB
}

func New(pgCfg config.Postgres) (*Storage, error) {
	dns := pgCfg.GetDNS()

	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, fmt.Errorf("postgres - New - Open: %v", err)
	}
	//TODO  это нужно отсюда убрать и вызывать отдельно

	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS users(
    	id SERIAL PRIMARY KEY NOT NULL,
    	email TEXT NOT NULL,
    	password TEXT NOT NULL);
	`)
	if err != nil {
		return nil, fmt.Errorf("postgres - New - prepareUser: %v", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("postgres - New - exec: %v", err)
	}

	stmt, err = db.Prepare(`CREATE TABLE IF NOT EXISTS notes(
    	id	serial	primary key,
    	note	text	not null,
    	user_id integer not null
			constraint notes_users_id_fk
				references "users"
			);
	`)
	if err != nil {
		return nil, fmt.Errorf("postgres - New - prepareNote: %v", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("postgres - New - exec: %v", err)
	}

	return &Storage{db}, nil
}
