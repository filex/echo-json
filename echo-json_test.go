package main

import (
	"reflect"
	"testing"
)

func TestReadData(t *testing.T) {
	tests := []struct {
		input []string
		want  *pairList
	}{
		{
			[]string{"foo", "bar"},
			&pairList{"foo": "bar"},
		},
		{
			[]string{"foo1", "bar", "foo2", "baz", "foo3", "baq"},
			&pairList{"foo1": "bar", "foo2": "baz", "foo3": "baq"},
		},
		// uneven number of arguments: default value is empty string
		{
			[]string{"foo", "bar", "baz"},
			&pairList{"foo": "bar", "baz": ""},
		},
		// key must not be empty
		{
			[]string{"", "bar"},
			nil,
		},
	}

	for _, test := range tests {
		got, err := readPairs(test.input)
		if test.want == nil && (got != nil || err == nil) {
			t.Errorf("readPairs(%v) should be nil (w/ error), got: %v", test.input, got)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("readPairs(%v) == %v, want: %v", test.input, got, test.want)
		}

	}
}
