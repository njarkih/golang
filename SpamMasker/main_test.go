package main

import (
	"testing"
)

// Тестирование функции HideLinks
func TestHideLinks(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"Here's my spammy page: http://hehefouls.netHAHAHA see you.",
			"Here's my spammy page: http://******************* see you.",
		},
		{
			"No any links here!",
			"No any links here!",
		},
		{
			"http://mylink",
			"http://******",
		},
		{
			"Many many http://links http://are http://here",
			"Many many http://***** http://*** http://****",
		},
		{
			"Broken link http://",
			"Broken link http://",
		},
	}

	for _, test := range tests {
		result := HideLinks(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected '%s' but got '%s'", test.input, test.expected, result)
		}
	}
}
