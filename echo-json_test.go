package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
		// types
		{
			[]string{"string:name", "alice", "int:age", "33", "float:score", "93.1", "bool:active", "1", "bool:admin", "false"},
			&pairList{"name": "alice", "age": int64(33), "score": 93.1, "active": true, "admin": false},
		},
		// type errors
		{
			[]string{"string:"},
			nil,
		},
		{
			[]string{"int:a", "123.4"},
			nil,
		},
		{
			[]string{"float:a", "asdf"},
			nil,
		},
		{
			[]string{"bool:a", "asdf"},
			nil,
		},
	}

	for _, test := range tests {
		got, err := readPairs(test.input)
		if test.want == nil && (got != nil || err == nil) {
			t.Errorf("readPairs(%v) should be nil (w/ error), got: %v", test.input, got)
		}
		if !cmp.Equal(got, test.want) {
			t.Errorf("readPairs(%v)\n%s", test.input, cmp.Diff(test.want, got))
		}

	}
}

func TestJSONResult(t *testing.T) {
	tests := []struct {
		input []string
		want  string // json result
	}{
		{
			[]string{"a", "b"},
			`{"a":"b"}`,
		},
		// raw type
		{
			[]string{"raw:x", "[1, 2, 3]"},
			`{"x":[1,2,3]}`,
		},
	}
	for _, test := range tests {
		got, err := args2JSON(test.input)
		if err != nil {
			t.Errorf("args2JSON(%v) should not fail, got error: %v", test.input, err)
		}
		if test.want != string(got) {
			t.Errorf("args2JSON(%v) == %v, got: %s", test.input, test.want, got)

		}

	}

}
