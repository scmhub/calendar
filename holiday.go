package calendar

import "time"

type holidayCalc func(*Holiday, int, *time.Location) time.Time

type Holiday struct {
	Name        string
	Years       []int // Used if holiday occures only on certain years
	Month       time.Month
	Day         int
	Weekday     time.Weekday
	NthWeekday  int
	Offset      int
	calc        holidayCalc
	observances []observance
}

func (h *Holiday) Calc(year int, loc *time.Location) time.Time {
	if len(h.Years) != 0 && !YearInRange(year, h.Years) {
		return time.Time{}
	}

	t := h.calc(h, year, loc)

	if h.Offset != 0 {
		t = t.AddDate(0, 0, h.Offset)
	}

	if h.observances == nil {
		return t
	}
	for _, o := range h.observances {
		t = o(t)
	}

	return t
}

func CalcDayOfMonth(h *Holiday, year int, loc *time.Location) time.Time {
	return time.Date(year, h.Month, h.Day, 0, 0, 0, 0, loc)
}

func CalcNthWeekday(h *Holiday, year int, loc *time.Location) time.Time {
	return time.Date(year, h.Month, h.Day, 0, 0, 0, 0, loc)
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

	return time.Date(year, time.Month(month), day+ho.Offset, 0, 0, 0, 0, loc)

}
