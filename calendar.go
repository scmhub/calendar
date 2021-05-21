package calendar

import (
	"fmt"
	"time"
)

const YearsAhead = 5
const YearsPast = 5

var (
	NewYork, _ = time.LoadLocation("America/New_York")
	Chicago, _ = time.LoadLocation("America/Chicago")
	Paris, _   = time.LoadLocation("Europe/Paris")
)

type Calendar struct {
	Name         string
	Loc          *time.Location
	startYear    int
	endYear      int
	opening      time.Duration
	closing      time.Duration
	earlyClosing time.Duration
	holidays     []*Holiday
	calendar     map[int64]*Holiday
}

func newCalendar(name string, loc *time.Location, start, end int) *Calendar {
	return &Calendar{
		Name:         name,
		Loc:          loc,
		startYear:    start,
		endYear:      end,
		opening:      9 * time.Hour,
		closing:      17*time.Hour + 30*time.Minute,
		earlyClosing: 12*time.Hour + 30*time.Minute,
		calendar:     make(map[int64]*Holiday),
	}
}

func NewCalendar(name string, loc string, years ...int) *Calendar {
	l, err := time.LoadLocation(loc)
	if err != nil {
		panic(err)
	}
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
	return newCalendar(name, l, start, end)
}

func (c *Calendar) Years() (start, end int) {
	return c.startYear, c.endYear
}

func (c *Calendar) SetYears(start, end int) {
	c.startYear, c.endYear = start, end
}

func (c *Calendar) addHoliday(h *Holiday) {
	for y := c.startYear; y <= c.endYear; y++ {
		t := h.Calc(y, c.Loc)
		if !t.IsZero() {
			c.calendar[t.Unix()] = h
		}
	}
}

func (c *Calendar) AddHoliday(h *Holiday) {
	c.holidays = append(c.holidays, h)
	c.addHoliday(h)
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

func (c *Calendar) String() string {
	str := fmt.Sprintf("Calendar %v:\n", c.Name)
	for t, h := range c.calendar {
		str += fmt.Sprintf("\t%v %v\n", time.Unix(t, 0).Format("2006/01/02"), h.Name)
	}
	return str
}
