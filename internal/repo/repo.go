package repo

import "github.com/ozonva/ova-rule-api/internal/models"

// Repo - интерфейс хранилища для сущности Rule.
type Repo interface {
	AddRules(rules []models.Rule) error
	ListRules(limit, offset uint64) ([]models.Rule, error)
	DescribeRule(ruleID uint64) (*models.Rule, error)
}
