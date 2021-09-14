package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ozonva/ova-rule-api/internal/models"
)

func TestSplitToBulksPositive(t *testing.T) {
	t.Parallel()

	cases := []struct {
		rules     []models.Rule
		batchSize int
		expected  [][]models.Rule
	}{
		{
			rules:     []models.Rule{{UserID: 1}, {UserID: 2}, {UserID: 3}, {UserID: 4}},
			batchSize: 3,
			expected:  [][]models.Rule{{{UserID: 1}, {UserID: 2}, {UserID: 3}}, {{UserID: 4}}},
		},
		{
			rules:     []models.Rule{{UserID: 1}, {UserID: 2}, {UserID: 3}},
			batchSize: 1,
			expected:  [][]models.Rule{{{UserID: 1}}, {{UserID: 2}}, {{UserID: 3}}},
		},
		{
			rules:     []models.Rule{{UserID: 1}, {UserID: 2}, {UserID: 3}},
			batchSize: 5,
			expected:  [][]models.Rule{{{UserID: 1}, {UserID: 2}, {UserID: 3}}},
		},
	}

	for _, c := range cases {
		actual, err := SplitToBulks(c.rules, c.batchSize)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, c.expected, actual)
	}
}

func TestSplitToBulksNegative(t *testing.T) {
	t.Parallel()

	emptyRulesError := errors.New("rules is empty or nil")
	notPositiveBatchSizeError := errors.New("batchSize must be positive number")

	cases := []struct {
		rules     []models.Rule
		batchSize int
		err       error
	}{
		{
			rules:     []models.Rule{},
			batchSize: 3,
			err:       emptyRulesError,
		},
		{
			rules:     nil,
			batchSize: 3,
			err:       emptyRulesError,
		},
		{
			rules:     []models.Rule{{UserID: 1}, {UserID: 2}, {UserID: 3}},
			batchSize: 0,
			err:       notPositiveBatchSizeError,
		},
		{
			rules:     []models.Rule{{UserID: 1}, {UserID: 2}, {UserID: 3}},
			batchSize: -1,
			err:       notPositiveBatchSizeError,
		},
	}

	for _, c := range cases {
		_, err := SplitToBulks(c.rules, c.batchSize)
		assert.EqualError(t, err, c.err.Error())
	}
}

func TestMapRulesPositive(t *testing.T) {
	t.Parallel()

	cases := []struct {
		rules    []models.Rule
		expected map[uint64]models.Rule
	}{
		{
			rules:    []models.Rule{{UserID: 1}, {UserID: 2}, {UserID: 3}},
			expected: map[uint64]models.Rule{1: {UserID: 1}, 2: {UserID: 2}, 3: {UserID: 3}},
		},
		{
			rules:    []models.Rule{},
			expected: map[uint64]models.Rule{},
		},
		{
			rules:    nil,
			expected: map[uint64]models.Rule{},
		},
	}

	for _, c := range cases {
		actual, err := MapRules(c.rules)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, c.expected, actual)
	}
}

func TestMapRulesNegative(t *testing.T) {
	t.Parallel()

	rules := []models.Rule{{UserID: 2}, {UserID: 2}}
	_, err := MapRules(rules)
	assert.EqualError(t, err, "duplicate key")
}
