package calendar

import (
	"time"
)

type holidayCalc func(*Holiday, int, *time.Location) time.Time

type Holiday struct {
	Name       string
	OnYear     int        // Used if holiday occures only on specific year
	BeforeYear int        // Used if holiday occures before a specific year
	AfterYear  int        // Used if holiday occures after or on a specific year
	Month      time.Month // Gregorian month also used for lunar month
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
	if h.OnYear > 0 && year != h.OnYear {
		return time.Time{}
	}

	if h.BeforeYear != 0 && year >= h.BeforeYear {
		return time.Time{}
	}

	if h.AfterYear != 0 && year < h.AfterYear {
		return time.Time{}
	}
	if h.OnYear < 0 {
		year += h.OnYear
	}
	t := h.calc(h, year, loc)

	if h.Offset != 0 {
		t = t.AddDate(0, 0, h.Offset)
	}

	if h.Observance != nil {
		t = h.Observance(t)
	}

	if IsWeekend(t) {
		return time.Time{}
	}

	return t
}

// Day of month in Gregorian Calendar
func CalcDayOfMonth(h *Holiday, year int, loc *time.Location) time.Time {
	return time.Date(year, h.Month, h.Day, 0, 0, 0, 0, loc)
}

// Day of month in Lunar Calendar
func CalcLunarDayOfMonth(h *Holiday, year int, loc *time.Location) time.Time {
	return FromLunar(time.Date(year, h.Month, h.Day, 0, 0, 0, 0, loc), false)
}

// Nth occurence of a weekday like 3rd monday, -1 means last monday of the month
func CalcNthWeekday(h *Holiday, year int, loc *time.Location) time.Time {
	month := h.Month
	if h.NthWeekday < 0 {
		month++
	}
	return NthWeekday(year, month, h.Weekday, h.NthWeekday, loc)
}

// Easter sunday
func CalcEasterOffset(h *Holiday, year int, loc *time.Location) time.Time {
	day, month := Easter(year)
	return time.Date(year, month, day, 0, 0, 0, 0, loc)
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
