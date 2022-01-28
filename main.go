// Program flatjson reads JSON object from the stdin nd outputs the flattened
// version into stdout.
//
// Example: cat test.json | go run main.go
package main

import (
	"aegorov.personal/flatjson/flatjson"
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	if err := flatjson.FlatJSON(&flatjson.Params{
		Reader: reader,
		Writer: writer,
	}); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run flatter: %v\n", err)
		os.Exit(1)
	}
	writer.Flush()
}
