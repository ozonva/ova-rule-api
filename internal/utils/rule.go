package utils

import (
	"errors"

	"github.com/ozonva/ova-rule-api/internal/models"
)

// SplitToBulks разделяет слайс правил на слайс слайсов.
func SplitToBulks(rules []models.Rule, batchSize int) ([][]models.Rule, error) {
	if batchSize <= 0 {
		return nil, errors.New("batchSize must be positive number")
	}

	if rules == nil || len(rules) == 0 {
		return nil, errors.New("rules is empty or nil")
	}

	rulesLength := len(rules)
	bulksLength := rulesLength / batchSize
	if bulksLength == 0 || rulesLength%batchSize > 0 {
		bulksLength++
	}

	bulks := make([][]models.Rule, bulksLength)

	for i := 0; i < bulksLength; i++ {
		start := i * batchSize
		end := i*batchSize + batchSize
		if end > rulesLength {
			end = rulesLength
		}
		bulks[i] = rules[start:end]
	}

	return bulks, nil
}
