package calendar

import (
	"fmt"
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
	assert.Equal("Asia/Kolkata", Mumbai.String())
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
}

func TestCalendarSessions(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago)
	assert.Equal(&Session{}, c.Session())
	assert.True(c.Session().IsZero())
	earlyOpen := 4 * time.Hour
	open := 9 * time.Hour
	close := 15 * time.Hour
	lateClose := 20 * time.Hour
	session := &Session{
		EarlyOpen: earlyOpen,
		Open:      open,
		Close:     close,
		LateClose: lateClose,
	}
	c.SetSession(session)
	assert.Equal(earlyOpen, c.Session().EarlyOpen)
	assert.Equal(open, c.Session().Open)
	assert.Equal(time.Duration(0), c.Session().BreakStart)
	assert.Equal(time.Duration(0), c.Session().BreakStop)
	assert.Equal(close, c.Session().Close)
	assert.Equal(lateClose, c.Session().LateClose)
	assert.False(session.HasBreak())
	breakStart := 11*time.Hour + 30*time.Minute
	breakStop := 12*time.Hour + 30*time.Minute
	session.BreakStart = breakStart
	session.BreakStop = breakStop
	c.SetSession(session)
	assert.Equal(open, c.Session().Open)
	assert.Equal(breakStart, c.Session().BreakStart)
	assert.Equal(breakStop, c.Session().BreakStop)
	assert.True(session.HasBreak())
	earlyClose := 13 * time.Hour
	session.EarlyClose = earlyClose
	c.SetSession(session)
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
	assert.Equal(1, len(c.h))
	assert.Equal(NewYear, c.h[0])
	assert.Equal(3, len(c.hmap))
	assert.Equal(NewYear, c.hmap[time.Date(2015, 1, 1, 0, 0, 0, 0, Chicago).Unix()])
	for i := 2; i < len(c.hmap); i++ {
		assert.True(c.IsHoliday(time.Date(2011+i, 1, 1, 0, 0, 0, 0, Chicago)))
	}
	c = NewCalendar("Calendar", Chicago, 2013)
	c.AddHolidays(NewYear) // 1/1/2013 is a tuesday
	assert.True(c.IsHoliday(time.Date(2013, 1, 1, 0, 0, 0, 0, Chicago)))
	assert.False(c.IsHoliday(time.Date(2013, 1, 2, 0, 0, 0, 0, Chicago)))
	assert.True(c.IsHoliday(time.Date(2013, 1, 1, 15, 30, 0, 0, Chicago)))
	assert.False(c.IsHoliday(time.Date(2013, 1, 2, 15, 30, 0, 0, Chicago)))
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
func TestIsOpen(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago, 2014, 2015)
	c.AddHolidays(NewYear)
	assert.PanicsWithError(fmt.Sprint(errNoSession), func() { c.IsOpen(time.Date(2014, 12, 31, 10, 30, 0, 0, Chicago)) })
	c.SetSession(&Session{
		EarlyOpen:  7 * time.Hour,
		Open:       9*time.Hour + 30*time.Minute,
		Close:      16 * time.Hour,
		EarlyClose: 13 * time.Hour,
		LateClose:  20 * time.Hour,
	})
	assert.True(c.IsOpen(time.Date(2014, 12, 31, 10, 30, 0, 0, Chicago)))
	assert.False(c.IsOpen(time.Date(2015, 1, 1, 10, 30, 0, 0, Chicago)))
	assert.True(c.IsOpen(time.Date(2015, 1, 2, 10, 30, 0, 0, Chicago)))
	assert.False(c.IsOpen(time.Date(2015, 1, 3, 10, 30, 0, 0, Chicago)))
	assert.False(c.IsOpen(time.Date(2015, 1, 4, 10, 30, 0, 0, Chicago)))
	assert.False(c.IsOpen(time.Date(2015, 1, 2, 9, 0, 0, 0, Chicago)))
	assert.True(c.IsOpen(time.Date(2015, 1, 2, 10, 30, 0, 0, Chicago)))
	assert.True(c.IsOpen(time.Date(2015, 1, 2, 12, 30, 0, 0, Chicago)))
	assert.False(c.IsOpen(time.Date(2015, 1, 2, 16, 30, 0, 0, Chicago)))
	c.SetSession(&Session{
		EarlyOpen:  7 * time.Hour,
		Open:       9*time.Hour + 30*time.Minute,
		Close:      16 * time.Hour,
		BreakStart: 11 * time.Hour,
		BreakStop:  13 * time.Hour,
		EarlyClose: 13 * time.Hour,
		LateClose:  20 * time.Hour,
	})
	assert.False(c.IsOpen(time.Date(2015, 1, 2, 9, 0, 0, 0, Chicago)))
	assert.True(c.IsOpen(time.Date(2015, 1, 2, 10, 30, 0, 0, Chicago)))
	assert.False(c.IsOpen(time.Date(2015, 1, 2, 12, 30, 0, 0, Chicago)))
	assert.False(c.IsOpen(time.Date(2015, 1, 2, 16, 30, 0, 0, Chicago)))
}

func TestTimestamps(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago, 2010, 2012)
	c.AddHolidays(NewYear, Epiphany)
	var prev int64
	for _, t := range c.hts {
		assert.True(prev < t)
		assert.True(c.IsHoliday(BOD(time.Unix(t, 0).In(c.Loc))))
	}
}
func TestNextClose(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago, 2014, 2015)
	c.AddHolidays(NewYear)
	assert.PanicsWithError(fmt.Sprint(errNoSession), func() { c.NextClose(time.Date(2014, 12, 31, 10, 30, 0, 0, Chicago)) })
	c.SetSession(&Session{
		Open:  9 * time.Hour,
		Close: 16 * time.Hour,
	})
	assert.Equal(time.Date(2014, 12, 31, 16, 00, 0, 0, Chicago), c.NextClose(time.Date(2014, 12, 31, 10, 30, 0, 0, Chicago)))
	assert.Equal(time.Date(2015, 1, 2, 16, 00, 0, 0, Chicago), c.NextClose(time.Date(2015, 1, 1, 10, 30, 0, 0, Chicago)))
	assert.Equal(time.Date(2015, 1, 2, 16, 00, 0, 0, Chicago), c.NextClose(time.Date(2015, 1, 2, 10, 30, 0, 0, Chicago)))
	assert.Equal(time.Date(2015, 1, 5, 16, 00, 0, 0, Chicago), c.NextClose(time.Date(2015, 1, 3, 10, 30, 0, 0, Chicago)))
	assert.Equal(time.Date(2015, 1, 5, 16, 00, 0, 0, Chicago), c.NextClose(time.Date(2015, 1, 4, 10, 30, 0, 0, Chicago)))
	assert.Equal(time.Date(2015, 1, 5, 16, 00, 0, 0, Chicago), c.NextClose(time.Date(2015, 1, 5, 10, 30, 0, 0, Chicago)))
}

func TestCalendarString(t *testing.T) {
	assert := assert.New(t)
	c := NewCalendar("Calendar", Chicago, 2011, 2015)
	assert.Equal("Calendar Calendar:\n", c.String())
	c.AddHolidays(NewYear)
	assert.Equal("Calendar Calendar:\n\t2013/01/01 New Year's Day\n\t2014/01/01 New Year's Day\n\t2015/01/01 New Year's Day\n", c.String())
}
