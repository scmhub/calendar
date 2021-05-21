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
