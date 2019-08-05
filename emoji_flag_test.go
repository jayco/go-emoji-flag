package emojiflag

import (
	"testing"
	"unicode/utf8"
)

func Test_getFlag(t *testing.T) {
	type args struct {
		country string
	}
	tests := []struct {
		name        string
		args        args
		expectedLen int
	}{
		{
			"Should handle correct input",
			args{"AUS"},
			8,
		},
		{
			"Should return empty string if match cannot be found",
			args{"BOB"},
			0,
		},
		{
			"Should uppercase input",
			args{"aus"},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetFlag(tt.args.country)
			if !utf8.ValidString(got) {
				t.Errorf("GetFlag() expected valid flag got %v", got)
			}
			if len(got) != tt.expectedLen {
				t.Errorf("expected length emoji of %v got %v", tt.expectedLen, got)
			}
		})
	}
}
