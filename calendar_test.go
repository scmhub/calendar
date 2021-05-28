package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLocations(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("America/Mexico_City", Mexico.String())
	assert.Equal("America/Chicago", Chicago.String())
	assert.Equal("America/Toronto", Toronto.String())
	assert.Equal("America/New_York", NewYork.String())
	assert.Equal("America/Sao_Paulo", SaoPaulo.String())
	assert.Equal("Europe/London", London.String())
	assert.Equal("Europe/Lisbon", Lisbon.String())
	assert.Equal("Europe/Madrid", Madrid.String())
	assert.Equal("Europe/Amsterdam", Amsterdam.String())
	assert.Equal("Europe/Brussels", Brussels.String())
	assert.Equal("Europe/Paris", Paris.String())
	assert.Equal("Europe/Zurich", Zurich.String())
	assert.Equal("Europe/Rome", Milan.String())
	assert.Equal("Europe/Berlin", Franckfurt.String())
	assert.Equal("Europe/Moscow", Moscow.String())
	assert.Equal("Africa/Johannesburg", Johannesburg.String())
	assert.Equal("Asia/Dubai", Dubai.String())
	assert.Equal("Asia/Kolkata", Bombay.String())
	assert.Equal("Asia/Singapore", Singapore.String())
	assert.Equal("Asia/Bangkok", Bangkok.String())
	assert.Equal("Asia/Hong_Kong", HongKong.String())
	assert.Equal("Asia/Hong_Kong", Shenzhen.String())
	assert.Equal("Asia/Shanghai", Shanghai.String())
	assert.Equal("Asia/Seoul", Seoul.String())
	assert.Equal("Asia/Tokyo", Tokyo.String())
	assert.Equal("Australia/Sydney", Sydney.String())
}

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
	assert.Nil(c.early)
	assert.Nil(c.morning)
	assert.Nil(c.afternoon)
	assert.Nil(c.late)
}

func TestCalendarSessions(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago)
	assert.Nil(c.EarlySession())
	s1, s2 := c.CoreSessions()
	assert.Nil(s1)
	assert.Nil(s2)
	assert.Nil(c.LateSession())
	early := &Session{4 * time.Hour, 9 * time.Hour}
	morning := &Session{9 * time.Hour, 11*time.Hour + 30*time.Minute}
	afternoon := &Session{12*time.Hour + 30*time.Minute, 15 * time.Hour}
	late := &Session{15 * time.Hour, 22 * time.Hour}
	c.SetEarlySession(early)
	assert.Equal(early, c.EarlySession())
	assert.Panics(func() { c.SetCoreSessions(early, morning, afternoon) })
	c.SetCoreSessions(morning, afternoon)
	s1, s2 = c.CoreSessions()
	assert.Equal(morning, s1)
	assert.Equal(afternoon, s2)
	c.SetCoreSessions(morning)
	s1, s2 = c.CoreSessions()
	assert.Equal(morning, s1)
	assert.Nil(s2)
	c.SetLateSession(late)
	assert.Equal(late, c.LateSession())
}

func TestCalendarYears(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago)
	start, end := c.Years()
	assert.Equal(time.Now().Year()-YearsPast, start)
	assert.Equal(time.Now().Year()+YearsAhead, end)
	c = NewCalendar("Calendar", Chicago, 2015)
	c.AddHolidays(NewYear)
	ti, ho := c.NextHoliday(time.Time{})
	assert.Equal(time.Date(2015, 1, 1, 0, 0, 0, 0, Chicago), ti)
	assert.Equal(NewYear, ho)
	c.SetYears(2018, 2020)
	assert.Equal(2018, c.startYear)
	assert.Equal(2020, c.endYear)
	ti, ho = c.NextHoliday(time.Time{})
	assert.Equal(time.Date(2018, 1, 1, 0, 0, 0, 0, Chicago), ti)
	assert.Equal(NewYear, ho)
}

func TestCalendarAddHoliday(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago, 2011, 2015)
	assert.False(c.HasHoliday(NewYear))
	c.AddHolidays(NewYear)
	assert.True(c.HasHoliday(NewYear))
	assert.Equal(1, len(c.holidays))
	assert.Equal(NewYear, c.holidays[0])
	assert.Equal(3, len(c.calendar))
	assert.Equal(NewYear, c.calendar[time.Date(2015, 1, 1, 0, 0, 0, 0, Chicago).Unix()])
	for i := 2; i < len(c.calendar); i++ {
		assert.True(c.IsHoliday(time.Date(2011+i, 1, 1, 0, 0, 0, 0, Chicago)))
	}
	c = NewCalendar("Calendar", Chicago, 2013)
	c.AddHolidays(NewYear) // 1/1/2013 is a tuesday
	assert.True(c.IsHoliday(time.Date(2013, 1, 1, 0, 0, 0, 0, Chicago)))
	assert.False(c.IsHoliday(time.Date(2013, 1, 2, 0, 0, 0, 0, Chicago)))
}

func TestCalendarBusinessDay(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago, 2013)
	c.AddHolidays(NewYear) // 1/1/2013 is a tuesday
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

func TestNextBusinessDay(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago, 2014, 2015)
	c.AddHolidays(NewYear)
	ti, ho := c.NextHoliday(time.Date(2013, 1, 1, 0, 0, 0, 0, Chicago))
	assert.Equal(time.Date(2014, 1, 1, 0, 0, 0, 0, Chicago), ti)
	assert.Equal(NewYear, ho)
	ti, ho = c.NextHoliday(time.Date(2014, 1, 1, 0, 0, 0, 0, Chicago))
	assert.Equal(time.Date(2015, 1, 1, 0, 0, 0, 0, Chicago), ti)
	assert.Equal(NewYear, ho)
	ti, ho = c.NextHoliday(time.Date(2014, 2, 1, 0, 0, 0, 0, Chicago))
	assert.Equal(time.Date(2015, 1, 1, 0, 0, 0, 0, Chicago), ti)
	assert.Equal(NewYear, ho)
	ti, ho = c.NextHoliday(time.Date(2015, 2, 1, 0, 0, 0, 0, Chicago))
	assert.Equal(time.Time{}, ti)
	assert.Nil(ho)

}

func TestTimestamps(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago, 2010, 2012)
	c.AddHolidays(NewYear, Epiphany)
	var prev int64
	for _, t := range c.timestamps {
		assert.True(prev < t)
		assert.True(c.IsHoliday(BOD(time.Unix(t, 0).In(c.Loc))))
	}
}

func TestCalendarString(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago, 2011, 2015)
	assert.Equal("Calendar Calendar:\n", c.String())
	c.AddHolidays(NewYear)
	assert.Equal("Calendar Calendar:\n\t2013/01/01 New Year's Day\n\t2014/01/01 New Year's Day\n\t2015/01/01 New Year's Day\n", c.String())
}
