package calendar

import (
	"time"
)

type holidayCalc func(*Holiday, int, *time.Location) time.Time

type Holiday struct {
	Name       string
	OnYear     int // Used if holiday occures only on specific year
	BeforeYear int // Used if holiday occures before a specific year
	AfterYear  int // Used if holiday occures after or on a specific year
	Month      time.Month
	Day        int
	Weekday    time.Weekday
	NthWeekday int
	Offset     int
	calc       holidayCalc
	Observance observance
}

func (h Holiday) Copy(name ...string) *Holiday {
	if len(name) > 0 {
		h.Name = name[0]
	}
	return &h
}

func (h *Holiday) SetOnYear(year int) *Holiday {
	h.OnYear = year
	return h
}

func (h *Holiday) SetBeforeYear(year int) *Holiday {
	h.BeforeYear = year
	return h
}

func (h *Holiday) SetAfterYear(year int) *Holiday {
	h.AfterYear = year
	return h
}

func (h *Holiday) SetOffset(o int) *Holiday {
	h.Offset = o
	return h
}

func (h *Holiday) SetObservance(o observance) *Holiday {
	h.Observance = o
	return h
}

func (h *Holiday) Calc(year int, loc *time.Location) time.Time {
	if h.OnYear != 0 && year != h.OnYear {
		return time.Time{}
	}

	if h.BeforeYear != 0 && year >= h.BeforeYear {
		return time.Time{}
	}

	if h.AfterYear != 0 && year < h.AfterYear {
		return time.Time{}
	}

	t := h.calc(h, year, loc)

	if h.Offset != 0 {
		t = t.AddDate(0, 0, h.Offset)
	}

	if h.Observance == nil {
		if IsWeekend(t) {
			return time.Time{}
		}
		return t
	}

	return h.Observance(t)
}

func CalcDayOfMonth(h *Holiday, year int, loc *time.Location) time.Time {
	return time.Date(year, h.Month, h.Day, 0, 0, 0, 0, loc)
}

func CalcNthWeekday(h *Holiday, year int, loc *time.Location) time.Time {
	month := h.Month
	if h.NthWeekday < 0 {
		month++
	}
	return NthWeekday(year, month, h.Weekday, h.NthWeekday, loc)
}

func CalcEasterOffset(ho *Holiday, year int, loc *time.Location) time.Time {
	var month, day int
	// Meeus/Jones/Butcher algorithm
	y := year
	a := y % 19
	b := y / 100
	c := y % 100
	d := b / 4
	e := b % 4
	f := (b + 8) / 25
	g := (b - f + 1) / 3
	h := (19*a + b - d - g + 15) % 30
	i := c / 4
	k := c % 4
	l := (32 + 2*e + 2*i - h - k) % 7
	m := (a + 11*h + 22*l) / 451
	n := h + l - 7*m + 114
	month = n / 31
	day = (n % 31) + 1

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)

}

// March Equinox
func CalcNorthwardEquinox(h *Holiday, year int, loc *time.Location) time.Time {
	c := northwardEquinox(year).In(loc)
	return time.Date(year, time.March, c.Day(), 0, 0, 0, 0, loc)
}

// June Solstice
func CalcNorthernSolstice(h *Holiday, year int, loc *time.Location) time.Time {
	c := northernSolstice(year).In(loc)
	return time.Date(year, time.June, c.Day(), 0, 0, 0, 0, loc)
}

// September Equinox
func CalcSouthwardEquinox(h *Holiday, year int, loc *time.Location) time.Time {
	c := southwardEquinox(year).In(loc)
	return time.Date(year, time.September, c.Day(), 0, 0, 0, 0, loc)
}

// December Solstice
func CalcSouthernSolstice(h *Holiday, year int, loc *time.Location) time.Time {
	c := southernSolstice(year).In(loc)
	return time.Date(year, time.December, c.Day(), 0, 0, 0, 0, loc)
}
