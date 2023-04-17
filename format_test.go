package gdk

import "testing"

func TestBytesToReadable(t *testing.T) {
	t.Run("no precision", func(t *testing.T) {
		tests := []struct {
			in       float64
			expected string
		}{
			{1023, "1023B"},
			{1023 * 1024, "1023KB"},
			{1023 * 1024 * 1024, "1023MB"},
			{1023 * 1024 * 1024 * 1024, "1023GB"},
			{1023 * 1024 * 1024 * 1024 * 1024, "1023TB"},
			{1023 * 1024 * 1024 * 1024 * 1024 * 1024, "1023PB"},
			{1023 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024, "1023EB"},
			{1023 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024, "1023ZB"},
			{1023 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024, "1023YB"},
			{1023 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024, "TooLarge"},
		}
		for _, test := range tests {
			if out := BytesToReadable(test.in); out != test.expected {
				t.Errorf("BytesToReadable()=%v, want:%v", out, test.expected)
			}
		}
	})

	t.Run("has precision", func(t *testing.T) {
		tests := []struct {
			in        float64
			precision int
			expected  string
		}{
			{1023, 1, "1023.0B"},
			{1023 * 1024, 2, "1023.00KB"},
			{1023 * 1024 * 1024, 2, "1023.00MB"},
			{1023 * 1024 * 1024 * 1024, 1, "1023.0GB"},
			{1023 * 1024 * 1024 * 1024 * 1024, 1, "1023.0TB"},
			{1023 * 1024 * 1024 * 1024 * 1024 * 1024, 3, "1023.000PB"},
			{1023 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024, 2, "1023.00EB"},
			{1023 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024, 1, "1023.0ZB"},
			{1023 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024, 1, "1023.0YB"},
			{1023 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024, 2, "TooLarge"},
		}
		for _, test := range tests {
			if out := BytesToReadable(test.in, test.precision); out != test.expected {
				t.Errorf("BytesToReadable()=%v, want:%v", out, test.expected)
			}
		}
	})
}
