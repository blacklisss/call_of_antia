package sqlitestore

import (
	"antia/internal/entities/userentity"
	"antia/internal/usecases/app/repos/userrepo"
	"context"
)

var _ userrepo.UserStore = &SQLiteRepository{}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    id,
    name
) VALUES ( $1, $2,)
`

func (q *SQLiteRepository) CreateUser(ctx context.Context, arg *userrepo.CreateUserParams) (*userentity.User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Name,
	)

	var user *userentity.User
	err := row.Scan(
		&user.ID,
		&user.Name,
	)

	return user, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1
`

func (q *SQLiteRepository) DeleteUser(ctx context.Context, id uint64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, name FROM users
WHERE id = $1 LIMIT 1
`

func (q *SQLiteRepository) GetUserByID(ctx context.Context, id uint64) (*userentity.User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)

	var user userentity.User
	err := row.Scan(
		&user.ID,
		&user.Name,
	)

	return &user, err
}
