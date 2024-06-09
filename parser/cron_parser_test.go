package parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	parser := &DefaultCronParser{}

	tests := []struct {
		input    string
		expected Cron
	}{
		{
			input: "*/15 0 1,15 * 1-5 /usr/bin/find",
			expected: Cron{
				Minute:     "0 15 30 45",
				Hour:       "0",
				DayOfMonth: "1 15",
				Month:      "1 2 3 4 5 6 7 8 9 10 11 12",
				DayOfWeek:  "1 2 3 4 5",
				Command:    "/usr/bin/find",
			},
		},
		{
			input: "0 12 * * 0-6 /usr/bin/backup",
			expected: Cron{
				Minute:     "0",
				Hour:       "12",
				DayOfMonth: "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31",
				Month:      "1 2 3 4 5 6 7 8 9 10 11 12",
				DayOfWeek:  "0 1 2 3 4 5 6",
				Command:    "/usr/bin/backup",
			},
		},
	}

	for _, test := range tests {
		result, err := parser.Parse(test.input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != test.expected {
			t.Errorf("For input '%s', expected %v but got %v", test.input, test.expected, result)
		}
	}
}

func TestInvalidInput(t *testing.T) {
	parser := &DefaultCronParser{}
	_, err := parser.Parse("invalid cron string")
	if err == nil {
		t.Errorf("Expected an error for invalid input, but got none")
	}
}
