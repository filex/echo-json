package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type pairList map[string]interface{}

func main() {
	pairs, err := readPairs(os.Args[1:])
	if err != nil {
		printError("argument error: %v", err)
	}

	b, err := json.Marshal(pairs)
	if err != nil {
		printError("json encode error: %v\n", err)
	}
	fmt.Printf("%s\n", b)
}

func readPairs(args []string) (*pairList, error) {
	num := len(args)
	pairs := make(pairList, num/2+num%2)

	for i := 0; i < num; i++ {
		var k, v string
		k = args[i]
		if k == "" {
			return nil, fmt.Errorf("key (arg %v) may not be empty\n", i)
		}
		i++
		if i < num {
			v = args[i]
		} else {
			v = ""
		}
		pairs[k] = v
	}
	return &pairs, nil
}

func printError(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}
