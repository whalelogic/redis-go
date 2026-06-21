package main_test

import(
	"github.com/whalelogic/redis-go-quick/misc"
	"testing"
)

func TestFinder(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		str    []string
		target string
		want   string
		want2  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := main.Finder(tt.str, tt.target)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("Finder() = %v, want %v", got, tt.want)
			}
			if true {
				t.Errorf("Finder() = %v, want %v", got2, tt.want2)
			}
		})
	}
}

