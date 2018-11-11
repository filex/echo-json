package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type fieldList map[string]interface{}

func main() {
	flag.Parse()

	data := readData(flag.Args())

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "json encode error: %v", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", b)
}

func readData(args []string) fieldList {
	num := len(args)
	count := num / 2
	count += num % 2
	data := make(map[string]interface{}, 2)

	for i := 0; i <= count; i++ {
		var k string
		k = args[i]
		i++
		var v interface{}
		if i < num {
			v = args[i]
		} else {
			v = ""
		}
		data[k] = v
	}
	return data
}
