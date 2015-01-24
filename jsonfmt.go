package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/* Usage:
 * % echo "{\"one\": \"two\"}" | go run jsonfmt.go
 */
func main() {
	reader := bufio.NewReader(os.Stdin)
	decoder := json.NewDecoder(reader)

	var o interface{}
	if err := decoder.Decode(&o); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//b, err := json.MarshalIndent(o, "", "	")
	b, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(b))
}
