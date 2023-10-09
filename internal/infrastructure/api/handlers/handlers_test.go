package handlers

import (
	"antia/internal/usecases/app/repos/relationrepo"
	"antia/internal/usecases/app/repos/runerepo"
	"antia/internal/usecases/app/repos/teamrepo"
	"antia/internal/usecases/app/repos/userrepo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHandlers(t *testing.T) {
	// Mock repositories
	mockUsers := &userrepo.Users{}
	mockRunes := &runerepo.Runes{}
	mockTeams := &teamrepo.Teams{}
	mockRelations := &relationrepo.Relations{}

	// Create a new Handlers instance using the NewHandlers function
	handlers := NewHandlers(mockUsers, mockRunes, mockTeams, mockRelations)

	// Assertions
	assert.NotNil(t, handlers)
	assert.Equal(t, mockUsers, handlers.us, "Expected user repo to be initialized correctly")
	assert.Equal(t, mockRunes, handlers.rs, "Expected runes repo to be initialized correctly")
	assert.Equal(t, mockTeams, handlers.ts, "Expected teams repo to be initialized correctly")
	assert.Equal(t, mockRelations, handlers.rl, "Expected relations repo to be initialized correctly")
}
