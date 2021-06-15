package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	thursday  = time.Date(2020, 1, 2, 0, 0, 0, 0, Paris)
	friday    = time.Date(2020, 1, 3, 0, 0, 0, 0, Paris)
	saturday  = time.Date(2020, 1, 4, 0, 0, 0, 0, Paris)
	sunday    = time.Date(2020, 1, 5, 0, 0, 0, 0, Paris)
	monday    = time.Date(2020, 1, 6, 0, 0, 0, 0, Paris)
	tuesday   = time.Date(2020, 1, 7, 0, 0, 0, 0, Paris)
	wednesday = time.Date(2020, 1, 8, 0, 0, 0, 0, Paris)
)

func TestNextMonday(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(thursday, nextMonday(thursday))
	assert.Equal(friday, nextMonday(friday))
	assert.Equal(monday, nextMonday(saturday))
	assert.Equal(monday, nextMonday(sunday))
	assert.Equal(monday, nextMonday(monday))
	assert.Equal(tuesday, nextMonday(tuesday))
	assert.Equal(wednesday, nextMonday(wednesday))
}

func TestPreviousFriday(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(thursday, previousFriday(thursday))
	assert.Equal(friday, previousFriday(friday))
	assert.Equal(friday, previousFriday(saturday))
	assert.Equal(friday, previousFriday(sunday))
	assert.Equal(monday, previousFriday(monday))
	assert.Equal(tuesday, previousFriday(tuesday))
	assert.Equal(wednesday, previousFriday(wednesday))
}

func TestNearestWorkday(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(thursday, nearestWorkday(thursday))
	assert.Equal(friday, nearestWorkday(friday))
	assert.Equal(friday, nearestWorkday(saturday))
	assert.Equal(monday, nearestWorkday(sunday))
	assert.Equal(monday, nearestWorkday(monday))
	assert.Equal(tuesday, nearestWorkday(tuesday))
	assert.Equal(wednesday, nearestWorkday(wednesday))
}

func TestSundayToMonday(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(thursday, sundayToMonday(thursday))
	assert.Equal(friday, sundayToMonday(friday))
	assert.Equal(saturday, sundayToMonday(saturday))
	assert.Equal(monday, sundayToMonday(sunday))
	assert.Equal(monday, sundayToMonday(monday))
	assert.Equal(tuesday, sundayToMonday(tuesday))
	assert.Equal(wednesday, sundayToMonday(wednesday))
}
func TestSundayToTuesday(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(thursday, sundayToTuesday(thursday))
	assert.Equal(friday, sundayToTuesday(friday))
	assert.Equal(saturday, sundayToTuesday(saturday))
	assert.Equal(tuesday, sundayToTuesday(sunday))
	assert.Equal(monday, sundayToTuesday(monday))
	assert.Equal(tuesday, sundayToTuesday(tuesday))
	assert.Equal(wednesday, sundayToTuesday(wednesday))
}
func TestSundayToWednesday(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(thursday, sundayToWednesday(thursday))
	assert.Equal(friday, sundayToWednesday(friday))
	assert.Equal(saturday, sundayToWednesday(saturday))
	assert.Equal(wednesday, sundayToWednesday(sunday))
	assert.Equal(monday, sundayToWednesday(monday))
	assert.Equal(tuesday, sundayToWednesday(tuesday))
	assert.Equal(wednesday, sundayToWednesday(wednesday))
}
func TestExeptOnWeekdays(t *testing.T) {
	assert := assert.New(t)
	exeptFriday := exeptOnWeekdays(time.Friday)
	assert.Equal(thursday, exeptFriday(thursday))
	assert.Equal(time.Time{}, exeptFriday(friday))
	assert.Equal(saturday, exeptFriday(saturday))
	assert.Equal(sunday, exeptFriday(sunday))
	assert.Equal(monday, exeptFriday(monday))
	assert.Equal(tuesday, exeptFriday(tuesday))
	assert.Equal(wednesday, exeptFriday(wednesday))
	exeptMTT := exeptOnWeekdays(time.Monday, time.Tuesday, time.Thursday)
	assert.Equal(time.Time{}, exeptMTT(thursday))
	assert.Equal(friday, exeptMTT(friday))
	assert.Equal(saturday, exeptMTT(saturday))
	assert.Equal(sunday, exeptMTT(sunday))
	assert.Equal(time.Time{}, exeptMTT(monday))
	assert.Equal(time.Time{}, exeptMTT(tuesday))
	assert.Equal(wednesday, exeptMTT(wednesday))
}
func TestOnlyOnWeekdays(t *testing.T) {
	assert := assert.New(t)
	onlyFriday := onlyOnWeekdays(time.Friday)
	assert.Equal(time.Time{}, onlyFriday(thursday))
	assert.Equal(friday, onlyFriday(friday))
	assert.Equal(time.Time{}, onlyFriday(saturday))
	assert.Equal(time.Time{}, onlyFriday(sunday))
	assert.Equal(time.Time{}, onlyFriday(monday))
	assert.Equal(time.Time{}, onlyFriday(tuesday))
	assert.Equal(time.Time{}, onlyFriday(wednesday))
	onlyMTT := onlyOnWeekdays(time.Monday, time.Tuesday, time.Thursday)
	assert.Equal(thursday, onlyMTT(thursday))
	assert.Equal(time.Time{}, onlyMTT(friday))
	assert.Equal(time.Time{}, onlyMTT(saturday))
	assert.Equal(time.Time{}, onlyMTT(sunday))
	assert.Equal(monday, onlyMTT(monday))
	assert.Equal(tuesday, onlyMTT(tuesday))
	assert.Equal(time.Time{}, onlyMTT(wednesday))
}
