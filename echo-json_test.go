package main

import (
	"reflect"
	"testing"
)

func TestReadData(t *testing.T) {
	tests := []struct {
		input []string
		want  fieldList
	}{
		{
			[]string{"foo", "bar"},
			fieldList{"foo": "bar"},
		},
		// uneven number of arguments: default value is empty string
		{
			[]string{"foo", "bar", "baz"},
			fieldList{"foo": "bar", "baz": ""},
		},
	}

	for _, test := range tests {
		got := readData(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("readData(%v) == %v, want: %v", test.input, got, test.want)
		}

	}
}
