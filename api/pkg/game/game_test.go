package game_test

import (
	"testing"
	"time"

	"github.com/dyrector-io/xor/api/pkg/game"
	"github.com/stretchr/testify/assert"
)

func TestDatePicking(t *testing.T) {
	const pickPerDay = 5
	dat, _ := time.Parse("2006-01-02", "2022-12-12")
	picks := game.PickByDate(dat, pickPerDay, 15, []int{})
	assert.Equalf(t, []int{12, 7, 2, 8, 3}, picks, "These numbers should've been picked on the first day")
	excl := picks
	picks = game.PickByDate(dat.AddDate(0, 0, 1), pickPerDay, 15, excl)
	assert.Equalf(t, []int{13, 4, 14, 9, 5}, picks, "These numbers should've been picked on the second day")
	excl = append(excl, picks...)
	picks = game.PickByDate(dat.AddDate(0, 0, 1), pickPerDay, 15, excl)
	assert.Equalf(t, []int{10, 6, 0, 11, 1}, picks, "These numbers should've been picked on the third day")
}
