package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleMethods(t *testing.T) {
	t.Parallel()

	rule := Rule{Name: "Number One", UserID: 123}
	assert.Equal(t, "<Number One: Count = 0>", rule.String())

	rule.IncCount()
	assert.Equal(t, "<Number One: Count = 1>", rule.String())
}
