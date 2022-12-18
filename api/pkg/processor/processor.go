package processor

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

type CNCFRecord struct {
	Name string `csv:"Name"`
	// Organization            string
	// Homepage                string
	Logo string
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
	// Description           string `csv:"Description"`
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

//go:embed landscape.json
var data []byte

func (c *CNCFRecord) String() string {
	return fmt.Sprintf("Company: %v, GH: %v, Crunchbase: %v",
		c.Name, c.GithubDescription, c.CrunchbaseDescription)
}

type CNCFSequence []*CNCFRecord

func (c *CNCFRecord) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ReadJSONData() CNCFSequence {
	list := CNCFSequence{}
	err := json.Unmarshal(data, &list)
	if err != nil {
		log.Fatal().Err(err)
		return nil
	}

	return list
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
			i.GithubDescription = Mask(i.GithubDescription, i.Name)
			i.CrunchbaseDescription = Mask(i.CrunchbaseDescription, i.Name)
		}
		result = append(result, i)
	}

	return result
}
