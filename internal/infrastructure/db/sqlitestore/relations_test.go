//go:build !integration

package sqlitestore

import (
	"antia/internal/entities/relationentity"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestAddRelation(t *testing.T) {
	// Create a mock DB connection and mock expectations
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %s", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Err(fmt.Errorf("Relations.TestAddRelation error: %w", err))
		}
	}(db)

	repo := &SQLiteRepository{db: db}

	tests := []struct {
		name     string
		relation *relationentity.Relation
		mockErr  error
		wantErr  bool
	}{
		{
			name: "Success",
			relation: &relationentity.Relation{
				UserID: 1,
				TeamID: 2,
				RuneID: 3,
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "DB Error",
			relation: &relationentity.Relation{
				UserID: 4,
				TeamID: 5,
				RuneID: 6,
			},
			mockErr: errors.New("mock DB error"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up mock expectations
			prep := mock.ExpectPrepare("INSERT INTO rune_relations")
			if tt.mockErr != nil {
				prep.ExpectExec().WithArgs(tt.relation.UserID, tt.relation.TeamID, tt.relation.RuneID).WillReturnError(tt.mockErr)
			} else {
				prep.ExpectExec().WithArgs(tt.relation.UserID, tt.relation.TeamID, tt.relation.RuneID).WillReturnResult(sqlmock.NewResult(1, 1))
			}

			// Call the method
			err := repo.AddRelation(context.Background(), tt.relation)

			// Check results
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Ensure all mock expectations were met
			err = mock.ExpectationsWereMet()
			assert.NoError(t, err, "there were unfulfilled expectations for this test")
		})
	}
}

func TestGetRelationByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %s", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Err(fmt.Errorf("Relations.TestGetRelationByUserID error: %w", err))
		}
	}(db)

	repo := &SQLiteRepository{db: db}

	tests := []struct {
		name       string
		userID     uint64
		mockRows   *sqlmock.Rows
		mockErr    error
		wantResult []*relationentity.NamedRelation
		wantErr    bool
	}{
		{
			name:   "Success",
			userID: 1,
			mockRows: sqlmock.NewRows([]string{"id", "user_id", "team_id", "rune_id", "team_name", "rune_name"}).
				AddRow(1, 1, 2, 3, "TeamA", "RuneA").
				AddRow(2, 1, 4, 5, "TeamB", "RuneB"),
			wantResult: []*relationentity.NamedRelation{
				{ID: 1, UserID: 1, TeamID: 2, RuneID: 3, TeamName: "TeamA", RuneName: "RuneA"},
				{ID: 2, UserID: 1, TeamID: 4, RuneID: 5, TeamName: "TeamB", RuneName: "RuneB"},
			},
			wantErr: false,
		},
		{
			name:    "DB Error",
			userID:  2,
			mockErr: errors.New("mock DB error"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prep := mock.ExpectPrepare("SELECT .* FROM rune_relations")
			if tt.mockErr != nil {
				prep.ExpectQuery().WithArgs(tt.userID).WillReturnError(tt.mockErr)
			} else {
				prep.ExpectQuery().WithArgs(tt.userID).WillReturnRows(tt.mockRows)
			}

			// Call the method
			result, err := repo.GetRelationByUserID(context.Background(), tt.userID)

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

func TestDeleteRelationByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %s", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Err(fmt.Errorf("Relations.TestDeleteRelationByID error: %w", err))
		}
	}(db)

	repo := &SQLiteRepository{db: db}

	tests := []struct {
		name    string
		id      uint64
		mockErr error
		wantErr bool
	}{
		{
			name:    "Success",
			id:      1,
			mockErr: nil,
			wantErr: false,
		},
		{
			name:    "DB Error",
			id:      2,
			mockErr: errors.New("mock DB error"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prep := mock.ExpectPrepare("DELETE FROM rune_relations")
			if tt.mockErr != nil {
				prep.ExpectExec().WithArgs(tt.id).WillReturnError(tt.mockErr)
			} else {
				prep.ExpectExec().WithArgs(tt.id).WillReturnResult(sqlmock.NewResult(0, 1))
			}

			// Call the method
			err := repo.DeleteRelationByID(context.Background(), tt.id)

			// Check results
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Ensure all mock expectations were met
			err = mock.ExpectationsWereMet()
			assert.NoError(t, err, "there were unfulfilled expectations for this test")
		})
	}
}
