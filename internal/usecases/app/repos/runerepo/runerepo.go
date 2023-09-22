package runerepo

import (
	"antia/internal/entities/runeentity"
	"context"

	"github.com/pkg/errors"
)

type RuneStore interface {
	GetRunes(ctx context.Context) ([]*runeentity.Rune, error)
}

type Runes struct {
	rstore RuneStore
}

func NewRunes(rstore RuneStore) *Runes {
	return &Runes{
		rstore,
	}
}

func (rs *Runes) GetRunes(ctx context.Context) ([]*runeentity.Rune, error) {
	runes, err := rs.rstore.GetRunes(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get runes error")
	}

	return runes, nil
}
