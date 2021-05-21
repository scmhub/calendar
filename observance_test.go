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
