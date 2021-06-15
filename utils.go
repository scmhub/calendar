package calendar

import (
	"time"
)

// Gregorian calendar start date
var LilianDate = time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC)

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

// Beginning Of Day is the first time of that day
func BOD(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// End Of Day is the last time of that day
func EOD(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, 999999999, t.Location())
}

// Nth occurence of a weekday like 3rd monday
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

// year -> day and month of Easter
func Easter(year int) (day int, month time.Month) {
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
	month = time.Month(n / 31)
	day = (n % 31) + 1
	return
}

// Gregorian calendar time -> Julian Day Number
func GtoJDN(t time.Time) int {
	t = t.UTC()
	a := (14 - int(t.Month())) / 12
	y := t.Year() + 4800 - a
	m := int(t.Month()) + 12*a - 3
	return t.Day() + (153*m+2)/5 + 365*y + y/4 - y/100 + y/400 - 32045
}

// Julian calendar time -> Julian Day Number
func JtoJDN(t time.Time) int {
	t = t.UTC()
	a := (14 - int(t.Month())) / 12
	y := t.Year() + 4800 - a
	m := int(t.Month()) + 12*a - 3
	return t.Day() + (153*m+2)/5 + 365*y + y/4 - 32083
}

// Julian Day Number -> Gregorian calendar time (UTC default loc if not provided)
func JDNtoG(jd int, loc ...*time.Location) time.Time {
	a := jd + 32044
	b := (4*a + 3) / 146097
	c := a - 146097*b/4
	d := (4*c + 3) / 1461
	e := c - (1461*d)/4
	m := (5*e + 2) / 153
	day := e - (153*m+2)/5 + 1
	month := time.Month(m + 3 - 12*(m/10))
	year := 100*b + d - 4800 + m/10
	t := time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
	if len(loc) != 0 {
		t = t.In(loc[0])
	}
	return t
}

// Julian Day Number -> Julian calendar time (UTC default loc if not provided)
func JDNtoJ(jd int, loc ...*time.Location) time.Time {
	b := 0
	c := jd + 32082
	d := (4*c + 3) / 1461
	e := c - (1461*d)/4
	m := (5*e + 2) / 153
	day := e - (153*m+2)/5 + 1
	month := time.Month(m + 3 - 12*(m/10))
	year := 100*b + d - 4800 + m/10
	t := time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
	if len(loc) != 0 {
		t = t.In(loc[0])
	}
	return t
}

// Julian calendar Time -> Gregorian calendar Time
func JulianToGegorian(t time.Time) time.Time {
	g := JDNtoG(JtoJDN(t), t.Location())
	g = time.Date(g.Year(), g.Month(), g.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), g.Location())
	return g
}

// Gregorian calendar Time -> Julian calendar Time
func GregorianToJulian(t time.Time) time.Time {
	j := JDNtoJ(GtoJDN(t), t.Location())
	j = time.Date(j.Year(), j.Month(), j.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), j.Location())
	return j
}

// Is a Leap Year in the Julian calendar
func IsJulianLeapYear(year int) bool {
	return year%4 == 0
}

// Is a Leap Year in the Gregorian calendar
func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

// Year, Month -> number of days in the month
func GetMonthDays(year int, month time.Month) int {
	monthDays := map[time.Month]int{1: 31, 2: 28, 3: 31, 4: 30, 5: 31, 6: 30, 7: 31, 8: 31, 9: 30, 10: 31, 11: 30, 12: 31}
	if month == time.February && IsLeapYear(year) {
		return 29
	}
	return monthDays[month]
}

// Get the number of days in the Gregorian calendar year
func GetYearDays(year int) int {
	if IsLeapYear(year) {
		return 366
	}
	return 365
}
