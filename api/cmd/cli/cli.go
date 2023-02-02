package main

import (
	"fmt"
	"strings"

	"github.com/dyrector-io/xor/api/pkg/game"
	"github.com/dyrector-io/xor/api/pkg/processor"
)

const gradingSplit = 3

func printAnswer(right bool) string {
	if right {
		return "++"
	}
	return "--"
}

func quiz() {
	records := processor.ReadJSONData()

	qNum := 5
	attempts := 3

	indices := game.PickRandom(qNum, len(records)-1)
	answers := map[int]bool{}

	fmt.Printf("How well do you know the landscape? Find out who has this as a value proposition / GitHub description.")
	for i := 0; i < qNum; i++ {
		fmt.Printf("\n%d. Question\n%s\nStars:%d\n",
			i+1, records[indices[i]].Description, records[indices[i]].GithubStars)
		fmt.Printf("Code sample:\n%s\n", records[indices[i]].CodeExample)
		if records[indices[i]].RandomFact != "" {
			fmt.Printf("Random fact:\n%s\n", records[indices[i]].RandomFact)
		}
		for j := attempts; j > 0; j-- {
			input := ""
			fmt.Scanln(&input)
			if strings.EqualFold(records[indices[i]].Name, input) ||
				(len(input) > 4 && strings.Contains(records[indices[i]].Name, input)) {
				answers[i] = true
				fmt.Printf("Nailed it! Onto the next one...\n")
				break
			}
			if j == attempts {
				fmt.Printf("Hint: %s\n", records[indices[i]].CodeExample)
			}
			if input == "" {
				fmt.Printf("You have to write something...\n")
				j++
			}
			fmt.Printf("Not: %v. Try again! %d attempts left\n", input, j-1)
		}
		fmt.Printf("It was: %s\n\n", records[indices[i]].Name)
	}

	good := 0
	for _, v := range answers {
		if v {
			good++
		}
	}
	fmt.Printf("Your result is: %d/%d\n", good, qNum)
	if good < gradingSplit {
		fmt.Printf("Better luck next time.\n")
	} else if good > gradingSplit {
		fmt.Printf("You've got it! GG!")
	}
	for i := 0; i < qNum; i++ {
		fmt.Printf("%d %s | %s\n", i+1, records[indices[i]].Name, printAnswer(answers[i]))
	}
}

func main() {
	quiz()
}
