package flusher

import (
	"github.com/ozonva/ova-rule-api/internal/models"
	"github.com/ozonva/ova-rule-api/internal/repo"
)

// Flusher - интерфейс для сброса задач в хранилище.
type Flusher interface {
	Flush(rules []models.Rule) []models.Rule
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения.
func NewFlusher(
	chunkSize int,
	ruleRepo repo.Repo,
) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		ruleRepo:  ruleRepo,
	}
}

type flusher struct {
	chunkSize int
	ruleRepo  repo.Repo
}

// Flush сбрасывает задачи в хранилище.
func (f *flusher) Flush(rules []models.Rule) []models.Rule {
	// TODO: имплементировать логику
	return rules
}
