package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/trimmer-io/go-csv"
)

type CNCFRecord struct {
	Name string `csv:"Name"`
	// Organization            string
	// Homepage                string
	// Logo                    string
	// Twitter                 string
	// CrunchbaseURL           string
	// MarketCap               string
	// Ticker                  string
	// Funding                 string
	// Member                  string
	// Relation                string
	// License                 string
	// Headquarters            string
	// LatestTweetDate         string
	// Description             string
	// CrunchbaseDescription   string
	// CrunchbaseHomepage      string
	// CrunchbaseCity          string
	// CrunchbaseRegion        string
	// CrunchbaseCountry       string
	// CrunchbaseTwitter       string
	// CrunchbaseLinkedin      string
	// CrunchbaseTicker        string
	// CrunchbaseKind          string
	// CrunchbaseMinEmployees  string
	// CrunchbaseMaxEmployees  string
	Category string `csv:"Category"`
	// Subcategory             string
	// OSS                     string
	// GithubRepo              string
	GithubStars       string `csv:"Github Stars"`
	GithubDescription string `csv:"Github Description"`
	// GithubLatestCommitDate  string
	// GithubLatestCommitLink  string
	// GithubReleaseDate       string
	// GithubReleaseLink       string
	// GithubStartCommitDate   string
	// GithubStartCommitLink   string
	// GithubContributorsCount string
	// GithubContributorsLink  string
	// Accepted                string
	// Incubation              string
	// Graduated               string
	// DevStatsUrl             string
	// ArtworkUrl              string
	// BlogUrl                 string
	// MailingListUrl          string
	// SlackUrl                string
	// YoutubeUrl              string
	// ChatChannel             string
}

type CNCFSequence []*CNCFRecord

func main() {
	rand.Seed(time.Now().UnixNano())
	filePath := "interactive-landscape.csv"

	// records, err := ReadFile(filePath)
	// if err != nil {
	// 	panic(err)
	// }

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	records, err := ReadStream(file)
	if err != nil {
		panic(err)
	}

	quiz(records)
}

func ReadStream(r io.Reader) (CNCFSequence, error) {
	records := CNCFSequence{}
	dec := csv.NewDecoder(r)

	// read and decode the file header
	line, err := dec.ReadLine()
	if err != nil {
		return nil, err
	}
	if _, err = dec.DecodeHeader(line); err != nil {
		return nil, err
	}

	// loop until EOF (i.e. dec.ReadLine returns an empty line and nil error);
	// any other error during read will result in a non-nil error
	for {
		// read the next line from stream
		line, err = dec.ReadLine()

		// check for read errors other than EOF
		if err != nil {
			return nil, err
		}

		// check for EOF condition
		if line == "" {
			break
		}

		// decode the record
		v := &CNCFRecord{}
		if err = dec.DecodeRecord(v, line); err != nil {
			return nil, err
		}

		v.Name = strings.TrimSuffix(v.Name, "\"")
		v.GithubDescription = strings.TrimSuffix(v.GithubDescription, "\"")
		v.Category = strings.TrimSuffix(v.Category, "\"")

		// process the record here
		records = append(records, v)
	}
	return records, nil
}

func quiz(records []*CNCFRecord) {
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

func printAnswer(right bool) string {
	if right {
		return "++"
	} else {
		return "--"
	}
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
