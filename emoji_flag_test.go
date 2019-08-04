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
		name string
		args args
	}{
		{
			"Should error if flag is incorrect",
			args{"AUS"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetFlag(tt.args.country)
			if !utf8.ValidString(got) {
				t.Errorf("GetFlag() expected valid flag got %v", got)
			}
		})
	}
}
