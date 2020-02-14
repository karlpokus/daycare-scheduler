package main

import "testing"

var testTable = []struct {
	isoweek, m1week int
}{
	{7, 3},
	{8, 1},
	{9, 2},
	{10, 3},
	{11, 1},
}

func TestIndex(t *testing.T) {
	for _, tt := range testTable {
		if Schedule[Index(tt.isoweek)] != tt.m1week {
			t.Errorf("oops!")
		}
	}
}
