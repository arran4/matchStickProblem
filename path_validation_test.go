package main

import (
	"testing"
)

func TestValidatePath(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"Valid relative path", "output.gif", false},
		{"Valid subdir path", "subdir/output.gif", false},
		{"Invalid absolute path", "/tmp/output.gif", true},
		{"Invalid directory traversal", "../output.gif", true},
		{"Invalid deep traversal", "subdir/../../output.gif", true},
		{"Invalid parent directory", "..", true},
		{"Valid explicit relative", "./output.gif", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePath(tt.path); (err != nil) != tt.wantErr {
				t.Errorf("validatePath(%q) error = %v, wantErr %v", tt.path, err, tt.wantErr)
			}
		})
	}
}
