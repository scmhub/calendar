package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewCalendar(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago)
	assert.Equal("Calendar", c.Name)
	assert.Equal(time.Now().Year()-YearsPast, c.startYear)
	assert.Equal(time.Now().Year()+YearsAhead, c.endYear)
	c = NewCalendar("Calendar", Chicago, 2000)
	assert.Equal(2000, c.startYear)
	assert.Equal(2000+YearsAhead+YearsPast, c.endYear)
	c = NewCalendar("Calendar", Chicago, 2010, 2030)
	assert.Equal(2010, c.startYear)
	assert.Equal(2030, c.endYear)
	c = NewCalendar("Calendar", Chicago, 2020, 15)
	assert.Equal(2020, c.startYear)
	assert.Equal(2035, c.endYear)
}

func TestCalendarYears(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago)
	start, end := c.Years()
	assert.Equal(time.Now().Year()-YearsPast, start)
	assert.Equal(time.Now().Year()+YearsAhead, end)
	c.SetYears(2010, 2020)
	assert.Equal(2010, c.startYear)
	assert.Equal(2020, c.endYear)
}

func TestCalendarAddHoliday(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago, 2010, 2012)
	assert.False(c.HasHoliday(NewYear))
	c.AddHoliday(NewYear)
	assert.True(c.HasHoliday(NewYear))
	assert.Equal(1, len(c.holidays))
	assert.Equal(NewYear, c.holidays[0])
	assert.Equal(3, len(c.calendar))
	assert.Equal(NewYear, c.calendar[time.Date(2010, 1, 1, 0, 0, 0, 0, Chicago).Unix()])
	for i := 0; i < len(c.calendar); i++ {
		assert.True(c.IsHoliday(time.Date(2010+i, 1, 1, 0, 0, 0, 0, Chicago)))
	}
	c = NewCalendar("Calendar", Chicago, 2013)
	c.AddHoliday(NewYear) // 1/1/2013 is a tuesday
	assert.True(c.IsHoliday(time.Date(2013, 1, 1, 0, 0, 0, 0, Chicago)))
	assert.False(c.IsHoliday(time.Date(2013, 1, 2, 0, 0, 0, 0, Chicago)))
}

func TestCalendarBusinessDay(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago, 2013)
	c.AddHoliday(NewYear) // 1/1/2013 is a tuesday
	assert.False(c.IsBusinessDay(time.Date(2013, 1, 1, 0, 0, 0, 0, Chicago)))
	assert.True(c.IsBusinessDay(time.Date(2013, 1, 2, 0, 0, 0, 0, Chicago)))
	assert.Equal(time.Date(2013, 1, 2, 0, 0, 0, 0, Chicago), c.NextBusinessDay(time.Date(2012, 12, 31, 0, 0, 0, 0, Chicago)))
	assert.Equal(time.Date(2013, 1, 2, 0, 0, 0, 0, Chicago), c.NextBusinessDay(time.Date(2013, 1, 1, 0, 0, 0, 0, Chicago)))
	assert.Equal(time.Date(2013, 1, 3, 0, 0, 0, 0, Chicago), c.NextBusinessDay(time.Date(2013, 1, 2, 0, 0, 0, 0, Chicago)))
	assert.Equal(time.Date(2013, 1, 4, 0, 0, 0, 0, Chicago), c.NextBusinessDay(time.Date(2013, 1, 3, 0, 0, 0, 0, Chicago)))
	assert.Equal(time.Date(2013, 1, 7, 0, 0, 0, 0, Chicago), c.NextBusinessDay(time.Date(2013, 1, 4, 0, 0, 0, 0, Chicago)))
	assert.Equal(time.Date(2013, 1, 7, 0, 0, 0, 0, Chicago), c.NextBusinessDay(time.Date(2013, 1, 5, 0, 0, 0, 0, Chicago)))
	assert.Equal(time.Date(2013, 1, 7, 0, 0, 0, 0, Chicago), c.NextBusinessDay(time.Date(2013, 1, 6, 0, 0, 0, 0, Chicago)))
	assert.Equal(time.Date(2013, 1, 8, 0, 0, 0, 0, Chicago), c.NextBusinessDay(time.Date(2013, 1, 7, 0, 0, 0, 0, Chicago)))

}
func TestSortedHolidaysTime(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago, 2010, 2012)
	c.AddHoliday(NewYear)
	c.AddHoliday(Epiphany)
	var prev time.Time
	for _, t := range c.sortedHolidaysTime() {
		assert.True(prev.Before(t))
		assert.True(c.IsHoliday(t))
	}
}

func TestCalendarString(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago, 2010, 2012)
	assert.Equal("Calendar Calendar:\n", c.String())
	c.AddHoliday(NewYear)
	assert.Equal("Calendar Calendar:\n\t2010/01/01 New Year's Day\n\t2011/01/01 New Year's Day\n\t2012/01/01 New Year's Day\n", c.String())
}
