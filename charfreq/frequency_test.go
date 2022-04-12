package main

import (
	"reflect"
	"testing"
)

func TestFrequency(t *testing.T) {
	got := frequency("ass")
	want := map[string]int{
		"a": 1, "s": 2,
	}
	if !reflect.DeepEqual(got, want) {

		t.Errorf("got ass = %v; want map[a:1 s:2]", got)
	}
}
