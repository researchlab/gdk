package gdk

import "testing"

func TestStringReverse(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "olleh"},
		{"123", "321"},
		{"112", "211"},
	}
	for _, tt := range tests {
		if got, err := StringReverse(tt.input); got != tt.want && err != nil {
			t.Errorf("got %v, want %v, err %v", got, tt.want, err)
		}
	}
}
