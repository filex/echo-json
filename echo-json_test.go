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
			[]string{"name:string", "alice", "age:int", "33", "score:float", "93.1", "active:bool", "1", "admin:bool", "false"},
			&pairList{"name": "alice", "age": int64(33), "score": 93.1, "active": true, "admin": false},
		},
		// types vs namespaced keys
		{
			[]string{"namespace:key:bool", "true"},
			&pairList{"namespace:key": true},
		},

		// unknown type -> keep
		{
			[]string{"bing:boo", "bar", "bing:go:bo", "ba"},
			&pairList{"bing:boo": "bar", "bing:go:bo": "ba"},
		},
		{
			[]string{"a:int", "123.4"},
			nil,
		},
		{
			[]string{"a:float", "asdf"},
			nil,
		},
		{
			[]string{"a:bool", "asdf"},
			nil,
		},
		// edge cases: type hint for last arg with missing value
		{
			[]string{"foo:int"},
			&pairList{"foo": 0},
		},
		{
			[]string{"foo:float"},
			&pairList{"foo": 0.0},
		},
		{
			[]string{"foo:bool"},
			&pairList{"foo": false},
		},
		{
			[]string{"foo:string"},
			&pairList{"foo": ""},
		},
		{
			[]string{"foo:raw"},
			&pairList{"foo": nil},
		},
		// #2 typed default values when empty string is given
		{
			[]string{"foo:float", "", "z"},
			&pairList{"foo": 0.0, "z": ""},
		},
		{
			[]string{"foo:int", "", "z"},
			&pairList{"foo": 0, "z": ""},
		},
		{
			[]string{"foo:bool", "", "z"},
			&pairList{"foo": false, "z": ""},
		},
		{
			[]string{"foo:raw", "", "z"},
			&pairList{"foo": nil, "z": ""},
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
		in      []string
		want    string // json result
		wantErr bool
	}{
		{
			in:   []string{"a", "b"},
			want: `{"a":"b"}`,
		},
		// raw type
		{
			in:   []string{"x:raw", "[1, 2, 3]"},
			want: `{"x":[1,2,3]}`,
		},
		// raw default
		{
			in:   []string{"x:raw"},
			want: `{"x":null}`,
		},
		// raw null
		{
			in:   []string{"x:raw", "null"},
			want: `{"x":null}`,
		},
		// raw number
		{
			in:   []string{"x:raw", "123.3"},
			want: `{"x":123.3}`,
		},
		// raw: most complicated string
		{
			in:   []string{"x:raw", "\"foo\""},
			want: `{"x":"foo"}`,
		},
		// error
		{
			in:      []string{""},
			want:    "Argument Error: key (arg 1) may not be empty\n",
			wantErr: true,
		},
		{
			in:      []string{"a:int", "NaN"},
			want:    `Argument Error: value (NaN) for key "a" is not an int: strconv.ParseInt: parsing "NaN": invalid syntax`,
			wantErr: true,
		},
		{
			in:      []string{"a:bool", "NaB"},
			want:    `Argument Error: value (NaB) for key "a" is not a bool: strconv.ParseBool: parsing "NaB": invalid syntax`,
			wantErr: true,
		},
		{
			in:      []string{"a:float", "NaF"},
			want:    `Argument Error: value (NaF) for key "a" is not a float: strconv.ParseFloat: parsing "NaF": invalid syntax`,
			wantErr: true,
		},
		{
			in: []string{"a:raw", "no valid json here"},
			// message is not printed, main() wraps MarshalerError
			want:    `json: error calling MarshalJSON for type json.RawMessage: invalid character 'o' in literal null (expecting 'u')`,
			wantErr: true,
		},
	}
	for _, test := range tests {
		got, err := args2JSON(test.in)
		if test.wantErr {
			if err == nil {
				t.Fatalf("args2JSON(%v) should fail, got: %v", test.in, got)
			}
			if test.want != err.Error() {
				t.Errorf("args2JSON(%v) should fail with %v", test.in, cmp.Diff(test.want, err.Error()))
			}
		} else {
			if test.want != string(got) {
				t.Errorf("args2JSON(%v) == %v\n%v", test.in, test.want, cmp.Diff(test.want, string(got)))
			}
		}
	}
}

func TestVersion(t *testing.T) {
	tests := []struct {
		version string
		want    string
	}{
		{"", "development"},
		{"v1.0", "1.0"},
		{"foo", "foo"},
	}

	for _, test := range tests {
		Version = test.version
		got := version()
		if got != test.want {
			t.Errorf("version(%s) == %s, got: %s", test.version, test.want, got)
		}
	}
}
