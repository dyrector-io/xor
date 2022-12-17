package main

import (
	"io"
	"strings"

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
		if err := dec.DecodeRecord(v, line); err != nil {
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
