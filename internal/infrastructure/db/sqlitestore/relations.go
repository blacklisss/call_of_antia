package sqlitestore

import (
	"antia/internal/entities/relationentity"
	"antia/internal/usecases/app/repos/relationrepo"
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

var _ relationrepo.RelationStore = &SQLiteRepository{}

const addRelation = `-- name: AddRelation :exec
INSERT INTO rune_relations (user_id, team_id, rune_id) VALUES (?,?,?)
`

func (q *SQLiteRepository) AddRelation(ctx context.Context, relation *relationentity.Relation) error {
	stmt, err := q.db.Prepare(addRelation)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Err(fmt.Errorf("Relations.AddRelation error: %w", err))
		}
	}(stmt)

	_, err = stmt.ExecContext(ctx, relation.UserID, relation.TeamID, relation.RuneID)
	if err != nil {
		return err
	}

	return nil
}

const getRelationByUserID = `-- name: GetRelationByUserID :list
SELECT rl.*, t.name as team_name, r.name as rune_name FROM rune_relations as rl
LEFT JOIN teams as t ON rl.team_id = t.id
LEFT JOIN rune_characteristics as r ON rl.rune_id = r.id
WHERE user_id = ?
ORDER BY r.name
`

func (q *SQLiteRepository) GetRelationByUserID(ctx context.Context, userID uint64) ([]*relationentity.NamedRelation, error) {
	stmt, err := q.db.Prepare(getRelationByUserID)
	if err != nil {
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Err(fmt.Errorf("Relations.GetRelationByUserID error: %w", err))
		}
	}(stmt)

	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		return nil, err
	}

	var items []*relationentity.NamedRelation
	for rows.Next() {
		var i relationentity.NamedRelation

		err = rows.Scan(&i.ID, &i.UserID, &i.TeamID, &i.RuneID, &i.TeamName, &i.RuneName)
		if err != nil {
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

const deleteRelation = `-- name: DeleteRelation :exec
DELETE FROM rune_relations where id = ?
`

func (q *SQLiteRepository) DeleteRelationByID(ctx context.Context, id uint64) error {
	fmt.Println(id)

	stmt, err := q.db.Prepare(deleteRelation)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Err(fmt.Errorf("Relations.DeleteRelationByID error: %w", err))
		}
	}(stmt)

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
