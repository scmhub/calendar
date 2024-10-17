package calendar

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

const YearsAhead = 5
const YearsPast = 5

// Predefined time locations for various cities.
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

// Session defines the operating hours and breaks for the calendar.
type Session struct {
	EarlyOpen  time.Duration
	Open       time.Duration
	BreakStart time.Duration
	BreakStop  time.Duration
	Close      time.Duration
	EarlyClose time.Duration
	LateClose  time.Duration
}

// HasBreak checks if the session has a break defined.
func (s *Session) HasBreak() bool {
	return s.BreakStart != 0
}

// IsZero checks if the Session is empty.
func (s Session) IsZero() bool {
	return s == Session{}
}

// Calendar represents a calendar with holidays and sessions.
type Calendar struct {
	Name      string
	Loc       *time.Location     // NewYork or time.LoadLocation("America/New_York")
	startYear int                // default is time.Now().Year() - YearsPast
	endYear   int                // default is time.Now().Year() + YearsAhead
	session   *Session           // Session
	h         []*Holiday         // Holidays list
	hts       []int64            // Sorted holidays timestamps (Unix time)
	hmap      map[int64]*Holiday // {timestamps: *Holiday} map
	ects      []int64            // Sorted early close timestamps (Unix time)
	ecmap     map[int64]*Holiday // {timestamps: *Holiday} map for early close days
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

// NewCalendar creates a new Calendar instance based on the provided parameters.
// It accepts a name for the calendar, a location, and an optional list of years.
// If no years are specified, it defaults to a range of 5 years before and after the current year.
// If one year is provided, it sets that year as the start and calculates the end as 10 years later.
// If two years are provided, it calculates the end year based on the first year and checks if the second is less than 100.
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

// Session returns the current session details.
func (c *Calendar) Session() *Session {
	return c.session
}

// SetSession updates the session details for the calendar.
func (c *Calendar) SetSession(s *Session) {
	c.session = s
}

// Years returns the start and end years of the calendar.
func (c *Calendar) Years() (start, end int) {
	return c.startYear, c.endYear
}

// SetYears updates the start and end years for the calendar and resets holidays.
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

// AddHolidays appends holidays to the calendar and adds them to the holiday list.
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

// AddEarlyClosingDays appends early closing holidays to the calendar.
func (c *Calendar) AddEarlyClosingDays(h ...*Holiday) {
	for _, ho := range h {
		c.addEarlyClosingDay(ho)
		c.h = append(c.h, ho)
	}
}

// HasHoliday checks if a specific holiday is present in the calendar.
func (c *Calendar) HasHoliday(h *Holiday) bool {
	for _, ho := range c.h {
		if h == ho {
			return true
		}
	}
	return false
}

func (c *Calendar) ensureInRange(t time.Time) {
	year := t.Year()
	if year < c.startYear || year > c.endYear {
		panic(fmt.Sprintf("provided time %v is outside the calendar range (%d - %d)", t, c.startYear, c.endYear))
	}
}

// IsBusinessDay checks if a specific Tims is a business day for this calendar.
func (c *Calendar) IsBusinessDay(t time.Time) bool {
	c.ensureInRange(t)
	if IsWeekend(t) {
		return false
	}
	if c.IsHoliday(t) {
		return false
	}
	return true
}

// IsHoliday checks if a specific Tims is a holiday day for this calendar.
func (c *Calendar) IsHoliday(t time.Time) bool {
	c.ensureInRange(t)
	_, ok := c.hmap[BOD(t).Unix()]
	return ok
}

// IsHoliday checks if a specific Tims is a early close for this calendar.
func (c *Calendar) IsEarlyClose(t time.Time) bool {
	c.ensureInRange(t)
	_, ok := c.ecmap[BOD(t).Unix()]
	return ok
}

// IsOpen checks if a specific Tims is in a business session for this calendar.
func (c *Calendar) IsOpen(t time.Time) bool {
	c.ensureInRange(t)
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

// NextBusinessDay returns the business day following the provided Time.
func (c *Calendar) NextBusinessDay(t time.Time) time.Time {
	c.ensureInRange(t)
	t = t.AddDate(0, 0, 1)
	for !c.IsBusinessDay(t) {
		t = t.AddDate(0, 0, 1)
	}
	return t
}

// NextHoliday returns the following holiday time and Holiday for the provided Time.
func (c *Calendar) NextHoliday(t time.Time) (time.Time, *Holiday) {
	c.ensureInRange(t)
	for _, ts := range c.hts {
		if t.Unix() < ts {
			return time.Unix(ts, 0).In(c.Loc), c.hmap[ts]
		}
	}
	return time.Time{}, nil
}

// NextClose return the next closing time
func (c *Calendar) NextClose(t time.Time) time.Time {
	c.ensureInRange(t)
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
	var allts []int
	for _, ts := range c.hts {
		allts = append(allts, int(ts))
	}
	for _, ts := range c.ects {
		allts = append(allts, int(ts))
	}
	sort.Ints(allts)
	for _, ts := range allts {
		t := int64(ts)
		h, ok := c.hmap[t]
		if ok {
			str += fmt.Sprintf("\t%-15v    %v\n", time.Unix(t, 0).In(c.Loc).Format("2006-Jan-02 Mon"), h.Name)
		}
		ec, ok := c.ecmap[t]
		if ok {
			str += fmt.Sprintf("\t%-15v ec %v\n", time.Unix(t, 0).In(c.Loc).Format("2006-Jan-02 Mon"), ec.Name)
		}
	}

	return str
}

var errNoSession = errors.New("no Session defined")
var errNoEarlyClose = errors.New("no EarlyClose defined")
