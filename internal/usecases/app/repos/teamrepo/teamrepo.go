package teamrepo

import (
	"antia/internal/entities/teamentity"
	"context"

	"github.com/pkg/errors"
)

type TeamStore interface {
	GetTeams(ctx context.Context) ([]*teamentity.Team, error)
}

type Teams struct {
	tstore TeamStore
}

func NewTeams(tstore TeamStore) *Teams {
	return &Teams{
		tstore,
	}
}

func (ts *Teams) GetTeams(ctx context.Context) ([]*teamentity.Team, error) {
	teams, err := ts.tstore.GetTeams(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get teams error")
	}

	return teams, nil
}
