package sqlite

import (
	"database/sql"
)

type Sqlite struct {
	db *sql.DB
}

func New() (*Sqlite, error) {
	db, err := sql.Open("sqlite3", config.AppConfig.StoreDB)
	if err != nil {
		return nil, err
	}

	// ping
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Sqlite{
		db: db,
	}, nil
}

func (s *Sqlite) Close() error {
	return s.db.Close()
}
