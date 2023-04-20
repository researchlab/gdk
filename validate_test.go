package gdk

import "testing"

func TestIsEmail(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"good case with qq", "122@qq.com", true},
		{"good case with gmail", "mike@gmail.com", true},
		{"bad case", "122.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmail(tt.input); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmailRFC(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"good case with qq", "122@qq.com", true},
		{"good case with gmail", "mike@gmail.com", true},
		{"bad case", "122.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmailRFC(tt.input); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUrl(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"good case", "http://www.baidu.com", true},
		{"good case with https", "https://www.baidu.com", true},
		{"good case with ftp", "ftp://www.baidu.com", true},
		{"bad case", "www.baidu.com", false},
	}
	for _, tt := range tests {
		if got := IsUrl(tt.input); got != tt.want {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
