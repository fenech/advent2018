package main

import (
	"io"
	"strings"
	"testing"
)

func Test_react(t *testing.T) {
	type args struct {
		polymer string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"a and A react, leaving nothing behind",
			args{"aA"},
			"",
		},
		{
			"bB destroys itself, leaving aA. As above, this then destroys itself, leaving nothing",
			args{"abBA"},
			"",
		},
		{
			"no two adjacent units are of the same type, and so nothing happens",
			args{"abAB"},
			"abAB",
		},
		{
			"even though aa and AA are of the same type, their polarities match, and so nothing happens",
			args{"aabAAB"},
			"aabAAB",
		},
		{
			"longer example",
			args{"dabAcCaCBAcCcaDA"},
			"dabCBAcaDA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := react(tt.args.polymer); got != tt.want {
				t.Errorf("react() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pt1(t *testing.T) {
	type args struct {
		handle io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"longer example",
			args{strings.NewReader("dabAcCaCBAcCcaDA")},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pt1(tt.args.handle); got != tt.want {
				t.Errorf("pt1() = %v, want %v", got, tt.want)
			}
		})
	}
}
