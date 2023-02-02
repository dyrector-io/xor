package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dyrector-io/xor/api/pkg/processor"
)

func main() {
	seq := processor.ReadJSONData()

	toWrite := processor.QuizSequence{}
	for _, i := range seq {
		if i.GithubStars > 1000 {
			toWrite = append(toWrite, i)
		}
	}

	bytes, err := json.MarshalIndent(toWrite, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", bytes)

	os.WriteFile("landscape.json", bytes, os.ModePerm)
}
