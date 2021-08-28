package repo

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ozonva/ova-rule-api/internal/models"
)

// Repo - интерфейс хранилища для сущности Rule.
type Repo interface {
	AddRules(rules []models.Rule) error
	ListRules(limit, offset uint64) ([]models.Rule, error)
	DescribeRule(ruleID uint64) (*models.Rule, error)
}

func NewRepo(ctx context.Context, pool *pgxpool.Pool) Repo {
	return &repo{
		ctx:  ctx,
		pool: pool,
	}
}

type repo struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func (r *repo) AddRules(rules []models.Rule) error {
	conn, err := r.pool.Acquire(r.ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	return nil
}

func (r *repo) ListRules(limit, offset uint64) ([]models.Rule, error) {
	conn, err := r.pool.Acquire(r.ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	return nil, nil
}

func (r *repo) DescribeRule(ruleID uint64) (*models.Rule, error) {
	conn, err := r.pool.Acquire(r.ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	return nil, nil
}
