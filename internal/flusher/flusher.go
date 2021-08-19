package flusher

import (
	"github.com/ozonva/ova-rule-api/internal/models"
	"github.com/ozonva/ova-rule-api/internal/repo"
	"github.com/ozonva/ova-rule-api/internal/utils"
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
	var result []models.Rule

	chunks, err := utils.SplitToBulks(rules, f.chunkSize)
	if err != nil {
		return rules
	}

	for i, chunk := range chunks {
		if err = f.ruleRepo.AddRules(chunk); err != nil {
			result = append(result, chunks[i]...)
		}
	}

	return result
}
