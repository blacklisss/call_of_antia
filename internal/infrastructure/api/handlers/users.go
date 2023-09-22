package handlers

import (
	"antia/internal/entities/runeentity"
	"antia/internal/entities/teamentity"
	"antia/internal/entities/userentity"
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

var ErrUserNotFound = errors.New("user not found")
var ErrTeamNotFound = errors.New("teams not found")
var ErrRuneNotFound = errors.New("runes not found")

type UserResponse struct {
	User   userentity.User
	Teams  []*teamentity.Team
	Runes  []*runeentity.Rune
	Result map[ResultKey][]RuneResult
}

type RuneResult struct {
	ID     uint64
	RuneID uint64
	Name   string
}

type ResultKey struct {
	TeamID   uint64
	TeamName string
}

func (rt *Handlers) GetUserByID(ctx context.Context, id uint64) (*UserResponse, error) {
	user, err := rt.us.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &UserResponse{}, ErrUserNotFound
		}
		return &UserResponse{}, fmt.Errorf("error when reading: %w", err)
	}

	teams, err := rt.ts.GetTeams(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &UserResponse{}, ErrTeamNotFound
		}
		return &UserResponse{}, fmt.Errorf("error when reading: %w", err)
	}

	runes, err := rt.rs.GetRunes(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &UserResponse{}, ErrRuneNotFound
		}
		return &UserResponse{}, fmt.Errorf("error when reading: %w", err)
	}

	tmp, err := rt.rl.GetRelationByUserID(ctx, id)
	if err != nil {
		return &UserResponse{}, err
	}

	result := make(map[ResultKey][]RuneResult)

	for _, t := range tmp {
		rr := RuneResult{
			ID:     t.ID,
			RuneID: t.RuneID,
			Name:   t.RuneName,
		}

		tt := ResultKey{
			TeamID:   t.TeamID,
			TeamName: t.TeamName,
		}
		if _, ok := result[tt]; !ok {
			result[tt] = []RuneResult{rr}
		} else {
			result[tt] = append(result[tt], rr)
		}
	}

	fmt.Println(result)
	return &UserResponse{
		*user,
		teams,
		runes,
		result,
	}, nil
}
