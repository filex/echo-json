package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	data := ReadData(flag.Args())

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "json encode error: %v", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", b)
}

func ReadData(args []string) map[string]interface{} {
	data := make(map[string]interface{})

	for i := 0; i < len(args); i++ {
		data[flag.Arg(i)] = flag.Arg(i + 1)
		i++
	}
	return data
}
