package convert

import "testing"

func TestReadableSize(t *testing.T) {
	tests := []struct {
		in       float64
		expected string
	}{
		{1023, "1023.0B"},
		{1023 * 1024, "1023.0KB"},
		{1023 * 1024 * 1024, "1023.0MB"},
		{1023 * 1024 * 1024 * 1024, "1023.0GB"},
		{1023 * 1024 * 1024 * 1024 * 1024, "1023.0TB"},
		{1023 * 1024 * 1024 * 1024 * 1024 * 1024, "1023.0PB"},
		{1023 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024, "TooLarge"},
	}
	for _, test := range tests {
		if out := ReadableSize(test.in); out != test.expected {
			t.Errorf("ReadableSize()=%v, want:%v", out, test.expected)
		}
	}
}
