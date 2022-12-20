package main

import (
	"encoding/json"
	"fmt"

	"github.com/dyrector-io/xor/api/pkg/processor"
)

func main() {
	seq := processor.ReadJSONData()

	fmt.Printf("%d", len(seq))
	bytes, err := json.MarshalIndent(seq, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", bytes)
}
