package sqlitestore

import (
	"antia/internal/entities/userentity"
	"antia/internal/usecases/app/repos/userrepo"
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestUserMethods(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %s", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Err(fmt.Errorf("Users.TestUserMethods error: %w", err))
		}
	}(db)

	repo := &SQLiteRepository{db: db}

	// Test CreateUser method
	t.Run("CreateUser", func(t *testing.T) {
		arg := &userrepo.CreateUserParams{
			ID:   uint64(1),
			Name: sql.NullString{Valid: true, String: "John Doe"},
		}

		mock.ExpectQuery("INSERT INTO users").WithArgs(arg.ID, arg.Name).WillReturnRows(sqlmock.NewRows([]string{}))

		err := repo.CreateUser(context.Background(), arg)

		assert.NoError(t, err)
	})

	// Test CreateUser error scenario
	t.Run("CreateUser_Error", func(t *testing.T) {
		arg := &userrepo.CreateUserParams{
			ID:   2,
			Name: sql.NullString{Valid: true, String: "Jane Doe"},
		}

		mock.ExpectQuery("INSERT INTO users").WithArgs(arg.ID, arg.Name).WillReturnError(sql.ErrTxDone)

		err := repo.CreateUser(context.Background(), arg)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "sqllitestore.users.CreateUser error")
	})

	// Test DeleteUser method
	t.Run("DeleteUser", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM users").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.DeleteUser(context.Background(), 1)

		assert.NoError(t, err)
	})

	// Test GetUserByID method
	t.Run("GetUserByID", func(t *testing.T) {
		expectedUser := &userentity.User{
			ID:   1,
			Name: "John Doe",
		}

		mock.ExpectQuery("SELECT id, name FROM users").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(expectedUser.ID, expectedUser.Name))

		user, err := repo.GetUserByID(context.Background(), 1)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser.ID, user.ID)
		assert.Equal(t, expectedUser.Name, user.Name)
	})

	t.Run("GetUserByID_Error", func(t *testing.T) {
		userID := uint64(2)

		mock.ExpectQuery("SELECT id, name FROM users").WithArgs(userID).WillReturnError(sql.ErrTxDone)

		user, err := repo.GetUserByID(context.Background(), userID)

		assert.Error(t, err)
		assert.Equal(t, &userentity.User{}, user)
	})

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "there were unfulfilled expectations")
}
