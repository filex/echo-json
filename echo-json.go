package main

import (
	"encoding/json"
	"flag"
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
	type_raw
)

var (
	showVersion = flag.Bool("v", false, "show version information")
)

// set on build time with -ldflags "-X …"
var Version string

func main() {
	flag.Parse()

	if *showVersion {
		printVersion()
		os.Exit(0)
	}

	b, err := args2JSON(flag.Args())
	if err != nil {
		if _, ok := err.(*json.MarshalerError); ok {
			printError("Argument Error: A raw value is not valid JSON\n")
		}
		printError("JSON encode error: %T %v\n", err, err)
	}

	fmt.Printf("%s\n", b)
}

func printVersion() {
	fmt.Printf("echo-json\nVersion: %s\nMore info at https://github.com/filex/echo-json\n", version())
}

func version() string {
	if Version == "" {
		Version = "development"
	}
	if Version[:1] == "v" {
		return Version[1:]
	}
	return Version
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

	var i int
	var k, v string

	isLast := func() bool {
		return i >= num
	}
	useDefault := func() bool {
		return isLast() || v == ""
	}

	for ; i < num; i++ {
		k = args[i]
		i++
		if !isLast() {
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
			if useDefault() {
				tv = 0
			} else if tv, err = strconv.ParseInt(v, 10, 64); err != nil {
				return nil, fmt.Errorf("value (%v) for key \"%v\" is not an int: %v", v, k, err)
			}
		case type_float:
			if useDefault() {
				tv = 0.0
			} else if tv, err = strconv.ParseFloat(v, 64); err != nil {
				return nil, fmt.Errorf("value (%v) for key \"%v\" is not a float: %v", v, k, err)
			}
		case type_bool:
			if useDefault() {
				tv = false
			} else if tv, err = strconv.ParseBool(v); err != nil {
				return nil, fmt.Errorf("value (%v) for key \"%v\" is not a bool: %v", v, k, err)
			}
		case type_raw:
			if useDefault() {
				tv = nil
			} else {
				tv = json.RawMessage([]byte(v))
			}
		}
		pairs[k] = tv
	}
	return &pairs, nil
}

func getType(key string) (argType, string) {
	if pos := strings.LastIndexByte(key, ':'); pos > -1 {
		// key:type or namespaced:key:type
		t := key[pos+1:]
		k := key[:pos]
		switch t {
		case "int":
			// age:int, type int, key "age"
			return type_int, k
		case "float":
			return type_float, k
		case "bool":
			return type_bool, k
		case "string":
			return type_string, k
		case "raw":
			return type_raw, k
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
