package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestYearRange(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int{2010, 2011, 2012}, YearRange(2010, 2012))
}

func TestYearInRange(t *testing.T) {
	assert := assert.New(t)
	var yr []int
	assert.False(YearInRange(2015, yr))
	yr = YearRange(2010, 2012)
	assert.False(YearInRange(2015, yr))
	assert.True(YearInRange(2010, yr))
	assert.True(YearInRange(2011, yr))
	assert.True(YearInRange(2011, yr))
}

func TestIsWeekend(t *testing.T) {
	assert := assert.New(t)
	assert.False(IsWeekend(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)))
	assert.True(IsWeekend(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)))
	assert.True(IsWeekend(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)))
	assert.False(IsWeekend(time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)))
}

func TestBOD(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(
		time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC),
		BOD(time.Date(2006, 1, 2, 15, 4, 5, 789, time.UTC)),
	)
}

func TestEOD(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(
		time.Date(2006, 1, 2, 23, 59, 59, 999999999, time.UTC),
		EOD(time.Date(2006, 1, 2, 15, 4, 5, 789, time.UTC)),
	)
}

func TestNthWeekday(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Time{}, NthWeekday(2020, 1, time.Friday, 0, Paris))

	assert.Equal(time.Date(2020, 1, 4, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Saturday, 1, Paris))
	assert.Equal(time.Date(2020, 1, 5, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Sunday, 1, Paris))
	// weekday < first day of the month
	assert.Equal(time.Date(2020, 1, 6, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Monday, 1, Paris))
	assert.Equal(time.Date(2020, 1, 13, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Monday, 2, Paris))
	assert.Equal(time.Date(2020, 1, 20, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Monday, 3, Paris))
	assert.Equal(time.Date(2020, 1, 27, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Monday, 4, Paris))
	assert.Equal(time.Date(2020, 2, 3, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Monday, 5, Paris))
	assert.Equal(time.Date(2020, 2, 10, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Monday, 6, Paris))
	assert.Equal(time.Date(2019, 12, 30, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Monday, -1, Paris))
	assert.Equal(time.Date(2019, 12, 23, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Monday, -2, Paris))
	assert.Equal(time.Date(2019, 12, 16, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Monday, -3, Paris))
	assert.Equal(time.Date(2019, 12, 9, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Monday, -4, Paris))
	assert.Equal(time.Date(2019, 12, 2, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Monday, -5, Paris))
	assert.Equal(time.Date(2019, 11, 25, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Monday, -6, Paris))
	// weekday == first day of the month
	assert.Equal(time.Date(2020, 1, 1, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Wednesday, 1, Paris))
	assert.Equal(time.Date(2020, 1, 8, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Wednesday, 2, Paris))
	assert.Equal(time.Date(2020, 1, 15, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Wednesday, 3, Paris))
	assert.Equal(time.Date(2020, 1, 22, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Wednesday, 4, Paris))
	assert.Equal(time.Date(2020, 1, 29, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Wednesday, 5, Paris))
	assert.Equal(time.Date(2020, 2, 5, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Wednesday, 6, Paris))
	assert.Equal(time.Date(2019, 12, 25, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Wednesday, -1, Paris))
	assert.Equal(time.Date(2019, 12, 18, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Wednesday, -2, Paris))
	assert.Equal(time.Date(2019, 12, 11, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Wednesday, -3, Paris))
	assert.Equal(time.Date(2019, 12, 4, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Wednesday, -4, Paris))
	assert.Equal(time.Date(2019, 11, 27, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Wednesday, -5, Paris))
	assert.Equal(time.Date(2019, 11, 20, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Wednesday, -6, Paris))
	// weekday > first day of the month
	assert.Equal(time.Date(2020, 1, 3, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Friday, 1, Paris))
	assert.Equal(time.Date(2020, 1, 10, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Friday, 2, Paris))
	assert.Equal(time.Date(2020, 1, 17, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Friday, 3, Paris))
	assert.Equal(time.Date(2020, 1, 24, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Friday, 4, Paris))
	assert.Equal(time.Date(2020, 1, 31, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Friday, 5, Paris))
	assert.Equal(time.Date(2020, 2, 7, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Friday, 6, Paris))
	assert.Equal(time.Date(2019, 12, 27, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Friday, -1, Paris))
	assert.Equal(time.Date(2019, 12, 20, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Friday, -2, Paris))
	assert.Equal(time.Date(2019, 12, 13, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Friday, -3, Paris))
	assert.Equal(time.Date(2019, 12, 6, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Friday, -4, Paris))
	assert.Equal(time.Date(2019, 11, 29, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Friday, -5, Paris))
	assert.Equal(time.Date(2019, 11, 22, 0, 0, 0, 0, Paris), NthWeekday(2020, 1, time.Friday, -6, Paris))

}

func TestJulianGregorian(t *testing.T) {
	assert := assert.New(t)
	gregorian := time.Date(2021, 6, 8, 12, 0, 0, 0, time.UTC)
	julian := time.Date(2021, 5, 26, 12, 0, 0, 0, time.UTC)
	hijri := time.Date(1442, 10, 27, 12, 0, 0, 0, time.UTC)
	jd := 2459374
	assert.Equal(jd, GtoJDN(gregorian))
	assert.Equal(jd, JtoJDN(julian))
	assert.Equal(jd, HtoJDN(hijri))
	assert.Equal(julian, JDNtoJ(jd))
	assert.Equal(gregorian, JDNtoG(jd))
	assert.Equal(hijri, JDNtoH(jd))
	assert.Equal(gregorian, JulianToGegorian(julian))
	assert.Equal(julian, GregorianToJulian(gregorian))
	assert.Equal(gregorian, HijriToGregorian(hijri))
	assert.Equal(hijri, GregorianToHijri(gregorian))
	assert.Equal(1442, HijriYear(2021))
}

func TestLeapYear(t *testing.T) {
	assert := assert.New(t)
	assert.False(IsJulianLeapYear(2021))
	assert.False(IsLeapYear(2021))
	assert.True(IsJulianLeapYear(2008))
	assert.True(IsLeapYear(2008))
	assert.True(IsJulianLeapYear(1900))
	assert.False(IsLeapYear(1900))
	assert.True(IsJulianLeapYear(2000))
	assert.True(IsLeapYear(2000))
}

func TestGetMonthDays(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(31, GetMonthDays(2020, time.January))
	assert.Equal(29, GetMonthDays(2020, time.February)) // Leap year
	assert.Equal(28, GetMonthDays(2021, time.February)) // Non leap year
	assert.Equal(31, GetMonthDays(2020, time.March))
	assert.Equal(30, GetMonthDays(2020, time.April))
	assert.Equal(31, GetMonthDays(2020, time.May))
	assert.Equal(30, GetMonthDays(2020, time.June))
	assert.Equal(31, GetMonthDays(2020, time.July))
	assert.Equal(31, GetMonthDays(2020, time.August))
	assert.Equal(30, GetMonthDays(2020, time.September))
	assert.Equal(31, GetMonthDays(2020, time.October))
	assert.Equal(30, GetMonthDays(2020, time.November))
	assert.Equal(31, GetMonthDays(2020, time.December))

}
