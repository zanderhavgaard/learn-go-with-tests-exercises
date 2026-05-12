package main

import "testing"

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		"struct with one string field",
		struct{
			Name string
		}{"Chris"},
		[]string{"Chris"},
	},

}
