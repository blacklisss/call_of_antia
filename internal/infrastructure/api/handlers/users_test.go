//go:build !integration

package handlers

import (
	"antia/internal/entities/relationentity"
	"antia/internal/entities/runeentity"
	"antia/internal/entities/teamentity"
	"antia/internal/entities/userentity"
	"antia/internal/usecases/app/repos/relationrepo"
	mocks4 "antia/internal/usecases/app/repos/relationrepo/mocks"
	"antia/internal/usecases/app/repos/runerepo"
	"antia/internal/usecases/app/repos/runerepo/mocks"
	"antia/internal/usecases/app/repos/teamrepo"
	mocks3 "antia/internal/usecases/app/repos/teamrepo/mocks"
	"antia/internal/usecases/app/repos/userrepo"
	mocks2 "antia/internal/usecases/app/repos/userrepo/mocks"
	"context"
	"database/sql"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	ctx := context.TODO()

	t.Run("successful GetUserByID", func(t *testing.T) {
		mockUserStore := mocks2.NewUserStore(t)
		mockUserRepo := userrepo.NewUsers(mockUserStore)

		mockTeamStore := mocks3.NewTeamStore(t)
		mockTeamRepo := teamrepo.NewTeams(mockTeamStore)

		mockRuneStore := mocks.NewRuneStore(t)
		mockRuneRepo := runerepo.NewRunes(mockRuneStore)

		mockRelationsStore := mocks4.NewRelationStore(t)
		mockRelationRepo := relationrepo.NewRelations(mockRelationsStore)
		handler := &Handlers{
			us: mockUserRepo,
			ts: mockTeamRepo,
			rs: mockRuneRepo,
			rl: mockRelationRepo,
		}

		// Set up mock methods
		mockUserStore.On("GetUserByID", ctx, uint64(1)).Return(&userentity.User{ID: 1, Name: "John"}, nil)
		mockTeamStore.On("GetTeams", ctx).Return([]*teamentity.Team{{ID: 2, Name: "Team A"}}, nil)
		mockRuneStore.On("GetRunes", ctx).Return([]*runeentity.Rune{{ID: 3, Name: "Rune X"}}, nil)
		mockRelationsStore.On("GetRelationByUserID", ctx, uint64(1)).Return([]*relationentity.NamedRelation{{ID: 4, TeamID: 2, RuneID: 3, TeamName: "Team A", RuneName: "Rune X"}}, nil)

		resp, err := handler.GetUserByID(ctx, 1)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "John", resp.User.Name)
		assert.Equal(t, "Team A", resp.Teams[0].Name)
		assert.Equal(t, "Rune X", resp.Runes[0].Name)
	})

	t.Run("error in GetUserByID", func(t *testing.T) {
		mockUserStore := mocks2.NewUserStore(t)
		mockUserRepo := userrepo.NewUsers(mockUserStore)
		handler := &Handlers{us: mockUserRepo}

		// Mock to return an error
		mockUserStore.On("GetUserByID", ctx, uint64(1)).Return(nil, sql.ErrNoRows)

		resp, err := handler.GetUserByID(ctx, 1)
		assert.Error(t, err)
		assert.Equal(t, ErrUserNotFound, err)
		assert.Nil(t, resp.Teams)
		assert.Nil(t, resp.Runes)

	})

	t.Run("error when fetching teams", func(t *testing.T) {
		mockUserStore := mocks2.NewUserStore(t)
		mockUserRepo := userrepo.NewUsers(mockUserStore)

		mockTeamStore := mocks3.NewTeamStore(t)
		mockTeamRepo := teamrepo.NewTeams(mockTeamStore)

		handler := &Handlers{
			us: mockUserRepo,
			ts: mockTeamRepo,
		}

		// Set up mock methods
		mockUserStore.On("GetUserByID", ctx, uint64(1)).Return(&userentity.User{ID: 1, Name: "John"}, nil)
		mockTeamStore.On("GetTeams", ctx).Return(nil, sql.ErrNoRows)

		resp, err := handler.GetUserByID(ctx, 1)
		assert.Error(t, err)
		assert.Equal(t, ErrTeamNotFound, err)
		assert.Nil(t, resp.Teams)
		assert.Nil(t, resp.Runes)
		assert.Nil(t, resp.Result)
	})

	t.Run("error when fetching runes", func(t *testing.T) {
		mockUserStore := mocks2.NewUserStore(t)
		mockUserRepo := userrepo.NewUsers(mockUserStore)

		mockTeamStore := mocks3.NewTeamStore(t)
		mockTeamRepo := teamrepo.NewTeams(mockTeamStore)

		mockRuneStore := mocks.NewRuneStore(t)
		mockRuneRepo := runerepo.NewRunes(mockRuneStore)

		handler := &Handlers{
			us: mockUserRepo,
			ts: mockTeamRepo,
			rs: mockRuneRepo,
		}

		// Set up mock methods
		mockUserStore.On("GetUserByID", ctx, uint64(1)).Return(&userentity.User{ID: 1, Name: "John"}, nil)
		mockTeamStore.On("GetTeams", ctx).Return([]*teamentity.Team{{ID: 2, Name: "Team A"}}, nil)
		mockRuneStore.On("GetRunes", ctx).Return(nil, sql.ErrNoRows)

		resp, err := handler.GetUserByID(ctx, 1)
		assert.Error(t, err)
		assert.Equal(t, ErrRuneNotFound, err)
		assert.Nil(t, resp.Runes)
		assert.Nil(t, resp.Result)
	})

	t.Run("error when fetching relations", func(t *testing.T) {
		mockUserStore := mocks2.NewUserStore(t)
		mockUserRepo := userrepo.NewUsers(mockUserStore)

		mockTeamStore := mocks3.NewTeamStore(t)
		mockTeamRepo := teamrepo.NewTeams(mockTeamStore)

		mockRuneStore := mocks.NewRuneStore(t)
		mockRuneRepo := runerepo.NewRunes(mockRuneStore)

		mockRelationsStore := mocks4.NewRelationStore(t)
		mockRelationRepo := relationrepo.NewRelations(mockRelationsStore)

		handler := &Handlers{
			us: mockUserRepo,
			ts: mockTeamRepo,
			rs: mockRuneRepo,
			rl: mockRelationRepo,
		}

		// Set up mock methods
		mockUserStore.On("GetUserByID", ctx, uint64(1)).Return(&userentity.User{ID: 1, Name: "John"}, nil)
		mockTeamStore.On("GetTeams", ctx).Return([]*teamentity.Team{{ID: 2, Name: "Team A"}}, nil)
		mockRuneStore.On("GetRunes", ctx).Return([]*runeentity.Rune{{ID: 3, Name: "Rune X"}}, nil)
		mockRelationsStore.On("GetRelationByUserID", ctx, uint64(1)).Return(nil, errors.New("some random error"))

		resp, err := handler.GetUserByID(ctx, 1)
		assert.Error(t, err)
		assert.Equal(t, "add relation error: some random error", err.Error())
		assert.Nil(t, resp.Result)
	})
}
