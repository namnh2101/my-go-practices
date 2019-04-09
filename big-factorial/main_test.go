package main

import (
	"regexp"
	"testing"
)

func Test_isIPv4(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIPv4(tt.args.input); got != tt.want {
				t.Errorf("isIPv4() = %v, want %v", got, tt.want)
			}
		})
	}
	regexp.MatchString("")
}
