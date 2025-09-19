package issues_test

import (
	"testing"

	"github.com/ardibello/estimate/internal/application/issues"
	"github.com/stretchr/testify/assert"
)

func TestContainsEstimate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// success cases
		{"singular day", "Estimate: 1 day", true},
		{"integer estimate", "Estimate: 2 days", true},
		{"decimal estimate", "Estimate: 2.5 days", true},
		{"estimate with text", "Task scheduled. Estimate: 0.75 days remaining.", true},

		// failure cases
		{"empty body", "", false},
		{"negative number", "Estimate: -2 days", false},
		{"no number", "Estimate: days", false},
		{"no estimate text", "Task scheduled.", false},
		{"incorrect format", "Estimation: 2 days", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := issues.ContainsEstimate(tt.input)
			assert.Equal(t, tt.expected, got, "containsEstimate(%q)", tt.input)
		})
	}
}
