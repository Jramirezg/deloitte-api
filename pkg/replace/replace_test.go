package replace

import (
	"testing"
)

func TestReplace(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		// Test cases.
		{"We really like the new security features of Google Cloud and Amazon", "We really like the new security features of Google© Cloud and Amazon©"},
		{"We really like the new security features of Oracle", "We really like the new security features of Oracle©"},
		{"We really like the new security features of Microsoft", "We really like the new security features of Microsoft©"},
		{"We really like the new security features of Amazon", "We really like the new security features of Amazon©"},
		{"We really like the new security features of Deloitte", "We really like the new security features of Deloitte©"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := Replace(tt.input)
			if result != tt.want {
				t.Errorf("Return is incorrect, got %s", result)
			}
		})
	}
}

var bench_string = "We really like the new security features of Google Cloud"

func BenchmarkReplace(b *testing.B) {
	// TODO: Initialize
	for i := 0; i < b.N; i++ {
		// TODO: Your Code Here
		Replace(bench_string)
	}
}
