package processor

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
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
	Description           string `csv:"Description"`
	CrunchbaseDescription string `csv:"Crunchbase Description"`
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
	Category    string `csv:"Category"`
	Subcategory string
	// OSS                     string
	// GithubRepo              string
	GithubStars       int    `csv:"Github Stars"`
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

func (c *CNCFRecord) String() string {
	return fmt.Sprintf("Company: %v, Short Description: %v, GH: %v, Crunchbase: %v",
		c.Name, c.Description, c.GithubDescription, c.CrunchbaseDescription)
}

type CNCFSequence []*CNCFRecord

func (c *CNCFRecord) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ReadJSONData() (CNCFSequence, error) {
	file, err := os.Open("landscape.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	list := CNCFSequence{}
	err = json.Unmarshal(bytes, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func ReadCSVStream() (CNCFSequence, error) {
	filePath := "interactive-landscape.csv"

	r, err := os.Open(filePath)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

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

		// process the record here
		v.Name = strings.TrimSuffix(v.Name, "\"")
		v.Description = strings.TrimSuffix(v.Description, "\"")
		v.CrunchbaseDescription = strings.TrimSuffix(v.CrunchbaseDescription, "\"")
		v.GithubDescription = strings.TrimSuffix(v.GithubDescription, "\"")
		v.Category = strings.TrimSuffix(v.Category, "\"")

		log.Info().Msgf("%v", v)
		records = append(records, v)
	}
	return records, nil
}

func Mask(original, mask string) string {
	return strings.ReplaceAll(original, mask, "...")
}

func MaskAndFilter(list CNCFSequence, masked bool, startCountFilter int) CNCFSequence {
	result := CNCFSequence{}

	for _, i := range list {
		if startCountFilter > 0 && i.GithubStars < startCountFilter {
			continue
		}

		if masked {
			i.Description = Mask(i.Description, i.Name)
		}
		result = append(result, i)
	}

	return result
}
