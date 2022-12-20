package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dyrector-io/xor/api/internal/config"
	"github.com/dyrector-io/xor/api/pkg/processor"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/proullon/ramsql/driver" // ramsql
)

const (
	SimpleDateFormat        = "2006-01-02"
	DBconnectAttempt        = 3
	DBconnectBackoffSeconds = 10
)

type Pick struct {
	gorm.Model
	// 1,2,3,4,5 format
	Picks string
	// day of pick
	Date time.Time `gorm:"unique"`
}

type History struct {
	Date     string
	Projects string
}

func (c *History) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type HistoryList []*History

func SimpleDay(t time.Time) time.Time {
	str := t.Format(SimpleDateFormat)
	val, _ := time.Parse(SimpleDateFormat, str)
	return val
}

func PersistPicks(db *gorm.DB, today time.Time, picks []int) error {
	strArr := []string{}
	for _, p := range picks {
		strArr = append(strArr, fmt.Sprintf("%d", p))
	}

	str := strings.Join(strArr, ",")
	db.Create(&Pick{Picks: str, Date: SimpleDay(today)})
	return nil
}

func GetPicksForDay(db *gorm.DB, day time.Time) []int {
	p := Pick{}

	prunedDay := SimpleDay(day)
	err := db.Where(&Pick{Date: prunedDay}).First(&p)
	if err.Error != nil {
		return []int{}
	}

	return ParsePicksToIntArr(p.Picks)
}

func ParsePicksToIntArr(pickStr string) []int {
	str := strings.Split(pickStr, ",")
	picks := []int{}
	for _, v := range str {
		i, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			log.Error().Err(err).Msg("very unlikely str to int parse error")
		}
		picks = append(picks, int(i))
	}
	return picks
}

func GetHistoryDB(db *gorm.DB) HistoryList {
	picks := []Pick{}
	result := db.Where("Date < ?", SimpleDay(time.Now())).Find(&picks)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("could not load history returning empty set")
		return HistoryList{}
	}
	all := processor.ReadJSONData()
	hist := HistoryList{}
	for i := range picks {
		pickList := ParsePicksToIntArr(picks[i].Picks)
		projectNames := []string{}
		for _, index := range pickList {
			projectNames = append(projectNames, all[index].Name)
		}
		hist = append(hist, &History{
			Date:     picks[i].Date.Format(SimpleDateFormat),
			Projects: strings.Join(projectNames, ","),
		})
	}
	return hist
}

func GetExclusionList(db *gorm.DB) []int {
	picks := []Pick{}
	result := db.Find(&picks)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("could not load history returning empty set")
		return []int{}
	}
	xList := []int{}
	for i := range picks {
		pickList := ParsePicksToIntArr(picks[i].Picks)
		xList = append(xList, pickList...)
	}
	return xList
}

func InitPostgres(cfg *config.AppConfig) *gorm.DB {
	db := connect(cfg)
	err := db.AutoMigrate(&Pick{})
	if err != nil {
		log.Fatal().Msgf("sql migrate: %s\n", err)
	}
	return db
}

func connect(cfg *config.AppConfig) *gorm.DB {
	if cfg.DSN == "" {
		sqlDB, err := sql.Open("ramsql", "test")
		if err != nil {
			log.Fatal().Msgf("ramsql.Open: %s\n", err)
		}

		db, err := gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{})
		if err != nil {
			log.Fatal().Msgf("sql.Open: %s\n", err)
		}

		return db
	}
	for i := 0; i < DBconnectAttempt; i++ {
		db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
		if err != nil {
			log.Error().Err(err).Msgf("sql error, backoff %d %d/%d\n", DBconnectBackoffSeconds, i+1, DBconnectAttempt)
			time.Sleep(DBconnectBackoffSeconds * time.Second)
		} else {
			return db
		}
		if i == DBconnectAttempt {
			log.Fatal().Msg("could not open sql connection")
		}
	}
	log.Fatal().Msg("db connect unreachable code")
	return nil
}
