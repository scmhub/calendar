package calendar

import (
	"time"
)

func YearRange(start, end int) []int {
	var r []int
	for i := start; i <= end; i++ {
		r = append(r, i)
	}
	return r
}

func YearInRange(year int, yearRange []int) bool {
	for _, y := range yearRange {
		if y == year {
			return true
		}
	}
	return false
}

func IsWeekend(t time.Time) bool {
	wd := t.Weekday()
	return wd == time.Saturday || wd == time.Sunday
}

// Beginning of day
func BOD(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// End of day
func EOD(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, 999999999, t.Location())
}

func NthWeekday(year int, month time.Month, weekday time.Weekday, n int, loc *time.Location) time.Time {
	if n == 0 {
		return time.Time{}
	}
	if n < 0 {
		n++
	}
	t := time.Date(year, month, 1, 0, 0, 0, 0, loc)
	wd := t.Weekday()
	if weekday < wd {
		return t.AddDate(0, 0, (7-int(wd-weekday))+((n-1)*7))
	}
	if weekday > wd {
		return t.AddDate(0, 0, int(weekday-wd)+((n-1)*7))
	}
	return t.AddDate(0, 0, (n-1)*7)
}
