package processor

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type QuizRecord struct {
	Name            string
	Description     string
	CodeExample     string
	RandomFact      string
	GithubStars     int
	GithubLink      string
	WeeklyDownloads int
}

//go:embed javascript.json
var data []byte

func (c *QuizRecord) String() string {
	return fmt.Sprintf("Company: %v, GH: %v,"+
		c.Name, c.Description)
}

type QuizSequence []*QuizRecord

func (c *QuizRecord) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ReadJSONData() QuizSequence {
	list := QuizSequence{}
	fmt.Printf("%d", len(data))
	err := json.Unmarshal(data, &list)
	if err != nil {
		log.Fatal().Err(err)
		return nil
	}
	return list
}

func MaskAndFilter(list QuizSequence, startCountFilter int) QuizSequence {
	result := QuizSequence{}

	for _, i := range list {
		if startCountFilter > 0 && i.GithubStars < startCountFilter {
			continue
		}
		result = append(result, i)
	}

	return result
}
