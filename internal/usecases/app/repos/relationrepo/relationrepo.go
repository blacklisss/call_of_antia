package relationrepo

import (
	"antia/internal/entities/relationentity"
	"context"

	"github.com/pkg/errors"
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.2 --name=RelationStore
type RelationStore interface {
	AddRelation(ctx context.Context, relation *relationentity.Relation) error
	GetRelationByUserID(ctx context.Context, userID uint64) ([]*relationentity.NamedRelation, error)
	DeleteRelationByID(ctx context.Context, id uint64) error
}

type Relations struct {
	rlstore RelationStore
}

func NewRelations(rlstore RelationStore) *Relations {
	return &Relations{
		rlstore,
	}
}

func (rls *Relations) AddRelation(ctx context.Context, relation *relationentity.Relation) error {
	err := rls.rlstore.AddRelation(ctx, relation)
	if err != nil {
		return errors.Wrap(err, "add relation error")
	}

	return nil
}

func (rls *Relations) GetRelationByUserID(ctx context.Context, userID uint64) ([]*relationentity.NamedRelation, error) {
	relations, err := rls.rlstore.GetRelationByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "add relation error")
	}

	return relations, nil
}

func (rls *Relations) DeleteRelationByID(ctx context.Context, id uint64) error {
	err := rls.rlstore.DeleteRelationByID(ctx, id)
	if err != nil {
		return errors.Wrap(err, "delete relation error")
	}

	return nil
}
