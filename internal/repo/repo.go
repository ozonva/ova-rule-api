package repo

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/ozonva/ova-rule-api/internal/kafka"
	"github.com/ozonva/ova-rule-api/internal/models"
)

// Repo - интерфейс хранилища для сущности Rule.
type Repo interface {
	AddRules(rules []models.Rule) error
	ListRules(limit, offset uint64) ([]models.Rule, error)
	DescribeRule(ruleID uint64) (*models.Rule, error)
	RemoveRule(ruleID uint64) error
}

func NewRepo(ctx context.Context, pool *pgxpool.Pool, producer kafka.AsyncProducer) Repo {
	return &repo{
		ctx:      ctx,
		pool:     pool,
		producer: producer,
	}
}

type repo struct {
	ctx      context.Context
	pool     *pgxpool.Pool
	producer kafka.AsyncProducer
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

	for _, rule := range rules {
		body := struct {
			Name   string `json:"name"`
			UserID uint64 `json:"user_id"`
		}{
			Name:   rule.Name,
			UserID: rule.UserID,
		}

		msg, err := encodeMessageToJSON(body)
		if err != nil {
			return err
		}

		preparedMsg := kafka.PrepareMessage(kafka.CreateRuleTopic, msg)
		r.producer.SendMessageWithContext(r.ctx, preparedMsg)

		log.Info().Msgf("Отправили в очередь событие про создание нового правила: %s", rule.Name)
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

	body := struct {
		ID uint64 `json:"id"`
	}{
		ID: ruleID,
	}

	msg, err := encodeMessageToJSON(body)
	if err != nil {
		return err
	}

	preparedMsg := kafka.PrepareMessage(kafka.RemoveRuleTopic, msg)
	r.producer.SendMessageWithContext(r.ctx, preparedMsg)

	log.Info().Msgf("Отправили в очередь событие про удаление правила с id=%d", ruleID)

	return nil
}

func encodeMessageToJSON(body interface{}) (string, error) {
	result, err := json.Marshal(body)
	if err != nil {
		log.Error().Msg("encode error")
		return "", err
	}

	return string(result), nil
}
