//go:build integration

package sqlitestore

import (
	"antia/internal/entities/userentity"
	"antia/internal/usecases/app/repos/userrepo"
	"context"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupUserDatabase(t *testing.T) *sql.DB {
	// Create an in-memory SQLite database.
	db, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)

	// Initialize your database schema.
	// Assuming the users table looks like: id INTEGER PRIMARY KEY, name TEXT
	_, err = db.Exec(`CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, password TEXT)`)
	assert.NoError(t, err)

	return db
}

func TestUserMethods(t *testing.T) {
	// Setup in-memory SQLite database.
	db := setupUserDatabase(t)

	repo := &SQLiteRepository{db: db}
	defer repo.Close()

	// Test CreateUser method
	t.Run("CreateUser", func(t *testing.T) {
		arg := &userrepo.CreateUserParams{
			ID:   uint64(1),
			Name: sql.NullString{Valid: true, String: "John Doe"},
		}

		err := repo.CreateUser(context.Background(), arg)
		assert.NoError(t, err)
	})

	// Test GetUserByID method
	t.Run("GetUserByID", func(t *testing.T) {
		expectedUser := &userentity.User{
			ID:   1,
			Name: "John Doe",
		}

		user, err := repo.GetUserByID(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser.ID, user.ID)
		assert.Equal(t, expectedUser.Name, user.Name)
	})

	// Test DeleteUser method
	t.Run("DeleteUser", func(t *testing.T) {
		err := repo.DeleteUser(context.Background(), 1)
		assert.NoError(t, err)
	})

}
