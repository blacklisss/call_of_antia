//go:build !integration

package sqlitestore

import (
	"antia/internal/entities/teamentity"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestGetTeams(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %s", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Err(fmt.Errorf("Teams.TestGetTeams error: %w", err))
		}
	}(db)

	repo := &SQLiteRepository{db: db}

	tests := []struct {
		name       string
		mockRows   *sqlmock.Rows
		mockErr    error
		wantResult []*teamentity.Team
		wantErr    bool
	}{
		{
			name: "Success",
			mockRows: sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, "TeamA").
				AddRow(2, "TeamB"),
			wantResult: []*teamentity.Team{
				{ID: 1, Name: "TeamA"},
				{ID: 2, Name: "TeamB"},
			},
			wantErr: false,
		},
		{
			name:     "DB Error",
			mockRows: sqlmock.NewRows([]string{}),
			mockErr:  errors.New("mock DB error"),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectQuery("SELECT .* FROM teams").WillReturnRows(tt.mockRows).WillReturnError(tt.mockErr)

			// Call the method
			result, err := repo.GetTeams(context.Background())

			// Check results
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantResult, result)
			}

			// Ensure all mock expectations were met
			err = mock.ExpectationsWereMet()
			assert.NoError(t, err, "there were unfulfilled expectations for this test")
		})
	}
}
