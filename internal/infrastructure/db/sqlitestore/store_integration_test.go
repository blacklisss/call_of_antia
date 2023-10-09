//go:build integration

package sqlitestore

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func setupDatabase(t *testing.T) *sql.DB {
	// Create an in-memory SQLite database.
	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err)

	// For more complex testing, you might want to initialize your database schema here.
	_, err = db.Exec(`CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, password TEXT)`)
	require.NoError(t, err)

	return db
}

func TestSQLiteRepositoryIn(t *testing.T) {
	// Setup in-memory SQLite database.
	db := setupDatabase(t)
	// Create SQLiteRepository instance.
	repo := NewSQLiteRepository(db)
	require.NotNil(t, repo)

	repo.Close()

}
