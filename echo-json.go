package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pairList map[string]interface{}

type argType int

const (
	type_string argType = iota
	type_int
	type_float
	type_bool
)

func main() {
	b, err := args2JSON(os.Args[1:])
	if err != nil {
		printError("json encode error: %v\n", err)
	}
	fmt.Printf("%s\n", b)
}

func args2JSON(args []string) ([]byte, error) {
	pairs, err := readPairs(args)
	if err != nil {
		return []byte(""), fmt.Errorf("Argument Error: %v", err)
	}

	return json.Marshal(pairs)
}

func readPairs(args []string) (*pairList, error) {
	num := len(args)
	pairs := make(pairList, num/2+num%2)

	for i := 0; i < num; i++ {
		var k, v string
		k = args[i]
		i++
		if i < num {
			v = args[i]
		} else {
			v = ""
		}

		// k can get a new name here
		t, k := getType(k)
		if k == "" {
			return nil, fmt.Errorf("key (arg %v) may not be empty\n", i)
		}
		var tv interface{}
		var err error
		switch t {
		case type_string:
			tv = v
		case type_int:
			if tv, err = strconv.ParseInt(v, 10, 64); err != nil {
				return nil, fmt.Errorf("value (%v) for key \"%v\" is not an int: %v", v, k, err)
			}
		case type_float:
			if tv, err = strconv.ParseFloat(v, 64); err != nil {
				return nil, fmt.Errorf("value (%v) for key \"%v\" is not a float: %v", v, k, err)
			}
		case type_bool:
			if tv, err = strconv.ParseBool(v); err != nil {
				return nil, fmt.Errorf("value (%v) for key \"%v\" is not a bool: %v", v, k, err)
			}
		}
		pairs[k] = tv
	}
	return &pairs, nil
}

func getType(key string) (argType, string) {
	if strings.IndexByte(key, ':') > -1 {
		parts := strings.SplitN(key, ":", 2)
		t := parts[0]
		k := parts[1]
		switch t {
		case "int":
			// int:age, type int, key "age"
			return type_int, k
		case "float":
			return type_float, k
		case "bool":
			return type_bool, k
		case "string":
			return type_string, k
		default:
			// foo:bar is string, key is "foo:bar"
			return type_string, key // return _key_ here!
		}
	}
	return type_string, key
}

func printError(msg string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}
