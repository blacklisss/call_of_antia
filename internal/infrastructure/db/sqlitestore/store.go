package sqlitestore

import (
	"database/sql"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (srepo *SQLiteRepository) Close() {
	srepo.db.Close()
}
