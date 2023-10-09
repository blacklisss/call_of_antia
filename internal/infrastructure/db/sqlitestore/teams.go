package sqlitestore

import (
	"antia/internal/entities/teamentity"
	"antia/internal/usecases/app/repos/teamrepo"
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

var _ teamrepo.TeamStore = &SQLiteRepository{}

const getTeams = `-- name: GetRunes :list
SELECT * FROM teams
`

func (q *SQLiteRepository) GetTeams(ctx context.Context) ([]*teamentity.Team, error) {
	rows, err := q.db.QueryContext(ctx, getTeams)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Err(fmt.Errorf("Teams.GetTeams error: %w", err))
		}
	}(rows)

	var items []*teamentity.Team
	for rows.Next() {
		var i teamentity.Team
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
