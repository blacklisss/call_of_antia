package handlers

import (
	"antia/internal/entities/relationentity"
	"antia/internal/usecases/app/repos/relationrepo"
	"antia/internal/usecases/app/repos/relationrepo/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandlers(t *testing.T) {
	ctx := context.TODO()

	// Test AddRuneForTeam
	t.Run("AddRuneForTeam", func(t *testing.T) {
		mockStore := mocks.NewRelationStore(t)
		mockRepo := relationrepo.NewRelations(mockStore)
		handler := &Handlers{rl: mockRepo}

		// Define input
		request := &RelationsRequest{
			UserID:  1,
			TeamID:  2,
			RunesId: []uint64{3, 4},
		}

		// Mock the AddRelation method
		mockStore.On("AddRelation", ctx, mock.AnythingOfType("*relationentity.Relation")).Return(nil)

		// Mock the GetRelationByUserID method
		mockStore.On("GetRelationByUserID", ctx, request.UserID).Return([]*relationentity.NamedRelation{{}}, nil)

		runes, err := handler.AddRuneForTeam(ctx, request)
		assert.NoError(t, err)
		assert.NotNil(t, runes)

		mockStore.AssertExpectations(t)
	})

	// Test DeleteRelationByID
	t.Run("DeleteRelationByID", func(t *testing.T) {
		mockStore := mocks.NewRelationStore(t)
		mockRepo := relationrepo.NewRelations(mockStore)
		handler := &Handlers{rl: mockRepo}

		// Define input
		request := &DeleteRelationsRequest{
			ID: 5,
		}

		// Mock the DeleteRelationByID method
		mockStore.On("DeleteRelationByID", ctx, request.ID).Return(nil)

		err := handler.DeleteRelationByID(ctx, request)
		assert.NoError(t, err)

		mockStore.AssertExpectations(t)
	})

	// Test AddRuneForTeam with a failing AddRelation (just as an example for error scenarios)
	t.Run("AddRuneForTeam with failing AddRelation", func(t *testing.T) {
		mockStore := mocks.NewRelationStore(t)
		mockRepo := relationrepo.NewRelations(mockStore)
		handler := &Handlers{rl: mockRepo}

		// Define input
		request := &RelationsRequest{
			UserID:  1,
			TeamID:  2,
			RunesId: []uint64{3, 4},
		}

		// Mock the AddRelation method to return an error
		mockStore.On("AddRelation", ctx, mock.AnythingOfType("*relationentity.Relation")).Return(errors.New("mock error"))

		_, err := handler.AddRuneForTeam(ctx, request)
		assert.Error(t, err)

		mockStore.AssertExpectations(t)
	})
}
