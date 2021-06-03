package calendar

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

const YearsAhead = 5
const YearsPast = 5

var (
	Mexico, _       = time.LoadLocation("America/Mexico_City")
	Chicago, _      = time.LoadLocation("America/Chicago")
	Toronto, _      = time.LoadLocation("America/Toronto")
	NewYork, _      = time.LoadLocation("America/New_York")
	SaoPaulo, _     = time.LoadLocation("America/Sao_Paulo")
	London, _       = time.LoadLocation("Europe/London")
	Lisbon, _       = time.LoadLocation("Europe/Lisbon")
	Madrid, _       = time.LoadLocation("Europe/Madrid")
	Amsterdam, _    = time.LoadLocation("Europe/Amsterdam")
	Brussels, _     = time.LoadLocation("Europe/Brussels")
	Paris, _        = time.LoadLocation("Europe/Paris")
	Zurich, _       = time.LoadLocation("Europe/Zurich")
	Milan, _        = time.LoadLocation("Europe/Rome")
	Franckfurt, _   = time.LoadLocation("Europe/Berlin")
	Moscow, _       = time.LoadLocation("Europe/Moscow")
	Johannesburg, _ = time.LoadLocation("Africa/Johannesburg")
	Dubai, _        = time.LoadLocation("Asia/Dubai")
	Mumbai, _       = time.LoadLocation("Asia/Kolkata")
	Singapore, _    = time.LoadLocation("Asia/Singapore")
	Bangkok, _      = time.LoadLocation("Asia/Bangkok")
	HongKong, _     = time.LoadLocation("Asia/Hong_Kong")
	Shenzhen, _     = time.LoadLocation("Asia/Hong_Kong")
	Shanghai, _     = time.LoadLocation("Asia/Shanghai")
	Seoul, _        = time.LoadLocation("Asia/Seoul")
	Tokyo, _        = time.LoadLocation("Asia/Tokyo")
	Sydney, _       = time.LoadLocation("Australia/Sydney")
)

type Session struct {
	EarlyOpen  time.Duration
	Open       time.Duration
	BreakStart time.Duration
	BreakStop  time.Duration
	Close      time.Duration
	EarlyClose time.Duration
	LateClose  time.Duration
}

func (s *Session) HasBreak() bool {
	return s.BreakStart != 0
}

func (s Session) IsZero() bool {
	return s == Session{}
}

type Calendar struct {
	Name      string
	Loc       *time.Location     // NewYork or time.LoadLocation("America/New_York")
	startYear int                // default is time.Now().Year() - YearsPast
	endYear   int                //default is time.Now().Year() + YearsAhead
	session   *Session           // regular session
	h         []*Holiday         // Holidays list
	hts       []int64            // Sorted holidays timestamps (Unix time)
	hmap      map[int64]*Holiday // timestamps: *Holiday map
	ects      []int64            // Sorted early close timestamps (Unix time)
	ecmap     map[int64]*Holiday // timestamps: *Holiday map for early close days
}

func newCalendar(name string, loc *time.Location, start, end int) *Calendar {
	return &Calendar{
		Name:      name,
		Loc:       loc,
		startYear: start,
		endYear:   end,
		session:   &Session{},
		hmap:      make(map[int64]*Holiday),
		ecmap:     make(map[int64]*Holiday),
	}
}

func NewCalendar(name string, loc *time.Location, years ...int) *Calendar {
	var start, end int
	switch len(years) {
	default:
		start = time.Now().Year() - YearsPast
		end = time.Now().Year() + YearsAhead
	case 1:
		start = years[0]
		end = years[0] + YearsPast + YearsAhead
	case 2:
		start = years[0]
		if years[1] < 100 {
			end = years[0] + years[1]
		} else {
			end = years[1]
		}
	}
	return newCalendar(name, loc, start, end)
}

func (c *Calendar) reset() {
	c.hts = []int64{}
	c.hmap = make(map[int64]*Holiday)
	c.ects = []int64{}
	c.ecmap = make(map[int64]*Holiday)
}

// early, core and late Sessions
func (c *Calendar) Session() *Session {
	return c.session
}

// early, core and late Sessions
func (c *Calendar) SetSession(s *Session) {
	c.session = s
}

func (c *Calendar) Years() (start, end int) {
	return c.startYear, c.endYear
}

func (c *Calendar) SetYears(start, end int) {
	c.startYear, c.endYear = start, end
	c.reset()

	for _, h := range c.h {
		c.addHoliday(h)
	}
}
func (c *Calendar) addHoliday(h *Holiday) {
	for y := c.startYear; y <= c.endYear; y++ {
		t := h.Calc(y, c.Loc)
		if !t.IsZero() {
			c.hmap[t.Unix()] = h
			c.hts = append(c.hts, t.Unix())
		}
	}
	sort.Slice(c.hts, func(i, j int) bool {
		return c.hts[i] < c.hts[j]
	})
}

func (c *Calendar) AddHolidays(h ...*Holiday) {
	for _, ho := range h {
		c.h = append(c.h, ho)
		c.addHoliday(ho)
	}

}

func (c *Calendar) addEarlyClosingDay(h *Holiday) {
	for y := c.startYear; y <= c.endYear; y++ {
		t := h.Calc(y, c.Loc)
		if !t.IsZero() {
			c.ecmap[t.Unix()] = h
			c.ects = append(c.ects, t.Unix())
		}
	}
	sort.Slice(c.ects, func(i, j int) bool {
		return c.ects[i] < c.ects[j]
	})
}

func (c *Calendar) AddEarlyClosingDays(h ...*Holiday) {
	for _, ho := range h {
		c.h = append(c.h, ho)
		c.addEarlyClosingDay(ho)
	}

}

func (c *Calendar) HasHoliday(h *Holiday) bool {
	for _, ho := range c.h {
		if h == ho {
			return true
		}
	}
	return false
}

func (c *Calendar) IsBusinessDay(t time.Time) bool {
	if IsWeekend(t) {
		return false
	}
	if c.IsHoliday(t) {
		return false
	}
	return true
}

func (c *Calendar) IsHoliday(t time.Time) bool {
	_, ok := c.hmap[BOD(t).Unix()]
	return ok
}

func (c *Calendar) IsEarlyClose(t time.Time) bool {
	_, ok := c.ecmap[BOD(t).Unix()]
	return ok
}

func (c *Calendar) IsOpen(t time.Time) bool {
	if c.session.IsZero() {
		panic(errNoSession)
	}
	if !c.IsBusinessDay(t) {
		return false
	}
	if c.IsEarlyClose(t) && t.After(BOD(t).Add(c.session.EarlyClose)) {
		return false
	}
	if c.session.HasBreak() && t.After(BOD(t).Add(c.session.BreakStart)) && t.Before(BOD(t).Add(c.session.BreakStop)) {
		return false
	}
	if t.Before(BOD(t).Add(c.session.Open)) {
		return false
	}
	if t.After(BOD(t).Add(c.session.Close)) {
		return false
	}
	return true
}

func (c *Calendar) NextBusinessDay(t time.Time) time.Time {
	t = t.AddDate(0, 0, 1)
	for !c.IsBusinessDay(t) {
		t = t.AddDate(0, 0, 1)
	}
	return t
}

func (c *Calendar) NextHoliday(t time.Time) (time.Time, *Holiday) {
	for _, ts := range c.hts {
		if t.Unix() < ts {
			return time.Unix(ts, 0).In(c.Loc), c.hmap[ts]
		}
	}
	return time.Time{}, nil
}

func (c *Calendar) NextClose(t time.Time) time.Time {
	if c.session.IsZero() {
		panic(errNoSession)
	}
	if c.IsBusinessDay(t) {
		if c.IsEarlyClose(t) {
			if c.session.EarlyClose == time.Duration(0) {
				panic(errNoEarlyClose)
			}
			return BOD(t).Add(c.session.EarlyClose)
		}
		return BOD(t).Add(c.session.Close)
	}
	return c.NextClose(c.NextBusinessDay(t))
}

func (c *Calendar) String() string {
	str := fmt.Sprintf("Calendar %v:\n", c.Name)
	for _, ts := range c.hts {
		str += fmt.Sprintf("\t%v %v\n", time.Unix(ts, 0).In(c.Loc).Format("2006/01/02"), c.hmap[ts].Name)
	}
	return str
}

var errNoSession = errors.New("no Session defined")
var errNoEarlyClose = errors.New("no EarlyClose defined")
