package sqlitestore

import (
	"antia/internal/entities/runeentity"
	"antia/internal/usecases/app/repos/runerepo"
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

var _ runerepo.RuneStore = &SQLiteRepository{}

const getRunes = `-- name: GetRunes :list
SELECT * FROM rune_characteristics
`

func (q *SQLiteRepository) GetRunes(ctx context.Context) ([]*runeentity.Rune, error) {
	rows, err := q.db.QueryContext(ctx, getRunes)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Err(fmt.Errorf("Runes.GetRunes error: %w", err))
		}
	}(rows)

	var items []*runeentity.Rune
	for rows.Next() {
		var i runeentity.Rune
		if err := rows.Scan(
			&i.ID,
			&i.Name,
		); err != nil {
			return nil, err
		}

		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
