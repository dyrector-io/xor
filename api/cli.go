package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func printAnswer(right bool) string {
	if right {
		return "++"
	}
	return "--"
}

func pick(amount, upperLimit int) []int {
	picked := []int{}
	min := 0
	for i := 0; i < amount; i++ {
		picked = append(picked, rand.Intn(upperLimit-min+1)+min)
	}
	return picked
}

func mask(original, mask string) string {
	return strings.ReplaceAll(original, mask, "...")
}

func quiz() {
	filePath := "interactive-landscape.csv"

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	records, err := ReadStream(file)
	if err != nil {
		panic(err)
	}
	qNum := 5
	attempts := 3

	indices := pick(qNum, len(records)-1)
	answers := map[int]bool{}
	fmt.Printf("How well do you know the landscape? Find out who has this as a value proposition / GitHub description.")
	for i := 0; i < qNum; i++ {
		// fmt.Printf("\n %d. Question \n%s\n", i+1, mask(records[indices[i]].GithubDescription, records[indices[i]].Name))
		fmt.Printf("\n%d. Question\n%s\nStars:%s\n", i+1, mask(records[indices[i]].GithubDescription, records[indices[i]].Name), records[indices[i]].GithubStars)
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
				fmt.Printf("Hint: %s\n", records[indices[i]].Category)
			}
			if input == "" {
				fmt.Printf("You have to write something...\n")
				j++
			}
			fmt.Printf("Not: %v. Try again! %d attemps left\n", input, j-1)
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
	if good < 3 {
		fmt.Printf("Better luck next time.\n")
	} else if good > 3 {
		fmt.Printf("You've got it! GG!")
	}
	for i := 0; i < qNum; i++ {
		fmt.Printf("%d %s | %s\n", i+1, records[indices[i]].Name, printAnswer(answers[i]))
	}
}
