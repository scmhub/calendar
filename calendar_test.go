package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewCalendar(t *testing.T) {
	assert := assert.New(t)
	assert.Panics(func() { NewCalendar("Calendar", "US/Chicago") })
	c := NewCalendar("Calendar", "America/Chicago")
	assert.Equal("Calendar", c.Name)
	assert.Equal(time.Now().Year()-YearsPast, c.startYear)
	assert.Equal(time.Now().Year()+YearsAhead, c.endYear)
	c = NewCalendar("Calendar", "America/Chicago", 2000)
	assert.Equal(2000, c.startYear)
	assert.Equal(2000+YearsAhead+YearsPast, c.endYear)
	c = NewCalendar("Calendar", "America/Chicago", 2010, 2030)
	assert.Equal(2010, c.startYear)
	assert.Equal(2030, c.endYear)
	c = NewCalendar("Calendar", "America/Chicago", 2020, 15)
	assert.Equal(2020, c.startYear)
	assert.Equal(2035, c.endYear)
}

func TestCalendarYears(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", "America/Chicago")
	start, end := c.Years()
	assert.Equal(time.Now().Year()-YearsPast, start)
	assert.Equal(time.Now().Year()+YearsAhead, end)
	c.SetYears(2010, 2020)
	assert.Equal(2010, c.startYear)
	assert.Equal(2020, c.endYear)
}

func TestCalendarAddHoliday(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", "America/Chicago", 2010, 2012)
	chicago, _ := time.LoadLocation("America/Chicago")
	c.AddHoliday(NewYear)
	assert.Equal(1, len(c.holidays))
	assert.Equal(NewYear, c.holidays[0])
	assert.Equal(3, len(c.calendar))
	assert.Equal(NewYear, c.calendar[time.Date(2010, 1, 1, 0, 0, 0, 0, chicago).Unix()])
	i := 0
	for ts, ho := range c.calendar {
		assert.Equal(time.Date(2010+i, 1, 1, 0, 0, 0, 0, chicago).Unix(), ts)
		assert.Equal(NewYear, ho)
		i++
	}
	assert.True(c.IsHoliday(time.Date(2010, 1, 1, 0, 0, 0, 0, chicago)))
}

func TestCalendarString(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", "America/Chicago", 2010, 2012)
	assert.Equal("Calendar Calendar:\n", c.String())
	c.AddHoliday(NewYear)
	assert.Equal("Calendar Calendar:\n\t2010/01/01 New Year's Day\n\t2011/01/01 New Year's Day\n\t2012/01/01 New Year's Day\n", c.String())
}
