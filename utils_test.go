package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestYearRange(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int{2010, 2011, 2012}, YearRange(2010, 2012))
}

func TestYearInRange(t *testing.T) {
	assert := assert.New(t)
	var yr []int
	assert.False(YearInRange(2015, yr))
	yr = YearRange(2010, 2012)
	assert.False(YearInRange(2015, yr))
	assert.True(YearInRange(2010, yr))
	assert.True(YearInRange(2011, yr))
	assert.True(YearInRange(2011, yr))
}

func TestIsWeekend(t *testing.T) {
	assert := assert.New(t)
	assert.False(IsWeekend(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)))
	assert.True(IsWeekend(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)))
	assert.True(IsWeekend(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)))
	assert.False(IsWeekend(time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)))
}

func TestBOD(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(
		time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC),
		BOD(time.Date(2006, 1, 2, 15, 4, 5, 789, time.UTC)),
	)
}

func TestEOD(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(
		time.Date(2006, 1, 2, 23, 59, 59, 999999999, time.UTC),
		EOD(time.Date(2006, 1, 2, 15, 4, 5, 789, time.UTC)),
	)
}
