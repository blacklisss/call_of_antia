//go:build !integration

package sqlitestore

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSQLiteRepository(t *testing.T) {
	// Mock database setup
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %s", err)
	}

	// Test NewSQLiteRepository
	t.Run("NewSQLiteRepository", func(t *testing.T) {
		repo := NewSQLiteRepository(db)

		assert.NotNil(t, repo)
		assert.Equal(t, db, repo.db)
	})

	// Test Close method
	t.Run("Close", func(t *testing.T) {
		mock.ExpectClose()

		repo := NewSQLiteRepository(db)
		repo.Close()

		err := mock.ExpectationsWereMet()
		assert.NoError(t, err, "there were unfulfilled expectations: %s", err)
	})
}
