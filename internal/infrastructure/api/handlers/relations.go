package handlers

import (
	"antia/internal/entities/relationentity"
	"context"
)

type RelationsRequest struct {
	UserID  uint64   `form:"user_id" binding:"required"`
	TeamID  uint64   `form:"team_id" binding:"required"`
	RunesId []uint64 `form:"runes_id" binding:"required"`
}

func (rt *Handlers) AddRuneForTeam(ctx context.Context, request *RelationsRequest) ([]*relationentity.NamedRelation, error) {
	for _, r := range request.RunesId {
		rlt := &relationentity.Relation{
			UserID: request.UserID,
			TeamID: request.TeamID,
			RuneID: r,
		}
		err := rt.rl.AddRelation(ctx, rlt)
		if err != nil {
			return nil, err
		}
	}

	runes, err := rt.rl.GetRelationByUserID(ctx, request.UserID)
	if err != nil {
		return nil, err
	}

	return runes, nil
}

type DeleteRelationsRequest struct {
	ID uint64 `form:"id" json:"id" uri:"id" binding:"required"`
}

func (rt *Handlers) DeleteRelationByID(ctx context.Context, request *DeleteRelationsRequest) error {
	err := rt.rl.DeleteRelationByID(ctx, request.ID)
	if err != nil {
		return err
	}

	return nil
}
