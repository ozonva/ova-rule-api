package repo

import (
	"context"

	"github.com/rs/zerolog/log"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/ozonva/ova-rule-api/internal/models"
)

// Repo - интерфейс хранилища для сущности Rule.
type Repo interface {
	AddRules(rules []models.Rule) error
	ListRules(limit, offset uint64) ([]models.Rule, error)
	DescribeRule(ruleID uint64) (*models.Rule, error)
	RemoveRule(ruleID uint64) error
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

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query := psql.Insert("rule").Columns("id", "name", "user_id")
	for _, rule := range rules {
		query = query.Values(rule.ID, rule.Name, rule.UserID)
	}
	sql, args, err := query.ToSql()

	log.Info().Msgf("query: %s; args: %s", sql, args)

	_, err = conn.Exec(r.ctx, sql, args...)
	if err != nil {
		log.Info().Msg(err.Error())
		return err
	}

	return nil
}

func (r *repo) ListRules(limit, offset uint64) ([]models.Rule, error) {
	conn, err := r.pool.Acquire(r.ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	sql, _, err := sq.Select("id, name, user_id").
		From("rule").
		Limit(limit).
		Offset(offset).
		ToSql()
	if err != nil {
		return nil, err
	}

	log.Info().Msgf("query: %s", sql)

	var rulePtrs []*models.Rule
	if err = pgxscan.Select(r.ctx, conn, &rulePtrs, sql); err != nil {
		return nil, err
	}

	rules := make([]models.Rule, len(rulePtrs))
	for i, ptr := range rulePtrs {
		rules[i] = *ptr
	}

	return rules, nil
}

func (r *repo) DescribeRule(ruleID uint64) (*models.Rule, error) {
	conn, err := r.pool.Acquire(r.ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sql, args, err := psql.Select("id, name, user_id").
		From("rule").
		Where(sq.Eq{"id": ruleID}).
		ToSql()
	if err != nil {
		return nil, err
	}

	log.Info().Msgf("query: %s; args: %s", sql, args)

	rule := models.Rule{}
	if err = pgxscan.Get(r.ctx, conn, &rule, sql, args...); err != nil {
		return nil, err
	}

	return &rule, nil
}

func (r *repo) RemoveRule(ruleID uint64) error {
	conn, err := r.pool.Acquire(r.ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sql, args, err := psql.Delete("rule").
		Where(sq.Eq{"id": ruleID}).
		ToSql()

	log.Info().Msgf("query: %s; args: %s", sql, args)

	_, err = conn.Exec(r.ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
