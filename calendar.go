package calendar

import (
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
	Bombay, _       = time.LoadLocation("Asia/Kolkata")
	Singapore, _    = time.LoadLocation("Asia/Singapore")
	Bangkok, _      = time.LoadLocation("Asia/Bangkok")
	HongKong, _     = time.LoadLocation("Asia/Hong_Kong")
	Shenzhen, _     = time.LoadLocation("Asia/Hong_Kong")
	Shanghai, _     = time.LoadLocation("Asia/Shanghai")
	Seoul, _        = time.LoadLocation("Asia/Seoul")
	Tokyo, _        = time.LoadLocation("Asia/Tokyo")
	Sydney, _       = time.LoadLocation("Australia/Sydney")
)

type Session [2]time.Duration

type Calendar struct {
	Name         string
	Loc          *time.Location
	startYear    int
	endYear      int
	early        *Session
	morning      *Session // If one Session morning is whole day
	afternoon    *Session
	late         *Session
	earlyClosing time.Duration
	holidays     []*Holiday // Holidays list including early closing
	timestamps   []int64    // Sorted holidays timestamps
	calendar     map[int64]*Holiday
}

func newCalendar(name string, loc *time.Location, start, end int) *Calendar {
	return &Calendar{
		Name:      name,
		Loc:       loc,
		startYear: start,
		endYear:   end,
		//opening:      9 * time.Hour,
		//closing:      17*time.Hour + 30*time.Minute,
		earlyClosing: 12*time.Hour + 30*time.Minute,
		calendar:     make(map[int64]*Holiday),
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
	c.timestamps = []int64{}
	c.calendar = make(map[int64]*Holiday)
}

func (c *Calendar) EarlySession() *Session {
	return c.early
}

func (c *Calendar) SetEarlySession(s *Session) {
	c.early = s
}

// morning, afternoon or day, nil
func (c *Calendar) CoreSessions() (*Session, *Session) {
	return c.morning, c.afternoon
}

// one *Session for the whole day, two *Sessions for morning and afernoon
func (c *Calendar) SetCoreSessions(s ...*Session) {
	if len(s) < 1 || len(s) > 2 {
		panic("Wrong Session number")
	}
	c.morning = s[0]
	c.afternoon = nil
	if len(s) > 1 {
		c.afternoon = s[1]
	}
}

func (c *Calendar) LateSession() *Session {
	return c.late
}

func (c *Calendar) SetLateSession(s *Session) {
	c.late = s
}

func (c *Calendar) Years() (start, end int) {
	return c.startYear, c.endYear
}

func (c *Calendar) SetYears(start, end int) {
	c.startYear, c.endYear = start, end
	c.reset()

	for _, h := range c.holidays {
		c.addHoliday(h)
	}
}

func (c *Calendar) addHoliday(h *Holiday) {
	for y := c.startYear; y <= c.endYear; y++ {
		t := h.Calc(y, c.Loc)
		if !t.IsZero() {
			c.calendar[t.Unix()] = h
			c.timestamps = append(c.timestamps, t.Unix())
		}
	}
	sort.Slice(c.timestamps, func(i, j int) bool {
		return c.timestamps[i] < c.timestamps[j]
	})
}

func (c *Calendar) AddHolidays(h ...*Holiday) {
	for _, holiday := range h {
		c.holidays = append(c.holidays, holiday)
		c.addHoliday(holiday)
	}

}

func (c *Calendar) HasHoliday(h *Holiday) bool {
	for _, holiday := range c.holidays {
		if h == holiday {
			return true
		}
	}
	return false
}

func (c *Calendar) IsHoliday(t time.Time) bool {
	_, ok := c.calendar[t.Unix()]
	return ok
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

func (c *Calendar) NextBusinessDay(t time.Time) time.Time {
	t = t.AddDate(0, 0, 1)
	for !c.IsBusinessDay(t) {
		t = t.AddDate(0, 0, 1)
	}
	return t
}

func (c *Calendar) NextHoliday(t time.Time) (time.Time, *Holiday) {
	for _, ts := range c.timestamps {
		if t.Unix() < ts {
			return time.Unix(ts, 0).In(c.Loc), c.calendar[ts]
		}
	}
	return time.Time{}, nil
}

func (c *Calendar) String() string {
	str := fmt.Sprintf("Calendar %v:\n", c.Name)
	for _, ts := range c.timestamps {
		str += fmt.Sprintf("\t%v %v\n", time.Unix(ts, 0).In(c.Loc).Format("2006/01/02"), c.calendar[ts].Name)
	}
	return str
}
