package calendar

import "time"

var (
	USNewYear = NewYear.Copy("US New Year's day").SetObservance(nearestWorkday)

	// MLKDay represents Martin Luther King Jr. Day on the 3rd Monday in January
	MLKDay = &Holiday{
		Name:       "Martin Luther King Jr. Day",
		Month:      time.January,
		Weekday:    time.Monday,
		NthWeekday: 3,
		calc:       CalcNthWeekday,
	}
	// PresidentsDay represents Presidents' Day on the 3rd Monday in February
	PresidentsDay = &Holiday{
		Name:       "Presidents' Day",
		Month:      time.February,
		Weekday:    time.Monday,
		NthWeekday: 3,
		calc:       CalcNthWeekday,
	}
	// MemorialDay represents Memorial Day on the last Monday in May
	MemorialDay = &Holiday{
		Name:       "Memorial Day",
		Month:      time.May,
		Weekday:    time.Monday,
		NthWeekday: -1,
		calc:       CalcNthWeekday,
	}
	// IndependenceDay represents Independence Day on 4-Jul
	IndependenceDay = &Holiday{
		Name:       "Independence Day",
		Month:      time.July,
		Day:        4,
		observance: nearestWorkday,
		calc:       CalcDayOfMonth,
	}

	// LaborDay represents Labor Day on the first Monday in September
	LaborDay = &Holiday{
		Name:       "Labor Day",
		Month:      time.September,
		Weekday:    time.Monday,
		NthWeekday: 1,
		calc:       CalcNthWeekday,
	}

	// ColumbusDay represents Columbus Day on the second Monday in October
	ColumbusDay = &Holiday{
		Name:       "Columbus Day",
		Month:      time.October,
		Weekday:    time.Monday,
		NthWeekday: 2,
		calc:       CalcNthWeekday,
	}

	// VeteransDay represents Veterans Day on 11-Nov
	VeteransDay = &Holiday{
		Name:  "Veterans Day",
		Month: time.November,
		Day:   11,
		calc:  CalcDayOfMonth,
	}

	// ThanksgivingDay represents Thanksgiving Day on the fourth Thursday in November
	ThanksgivingDay = &Holiday{
		Name:       "Thanksgiving Day",
		Month:      time.November,
		Weekday:    time.Thursday,
		NthWeekday: 4,
		calc:       CalcNthWeekday,
	}

	// Black Friday is the day following Thanksgiving
	BlackFriday = &Holiday{
		Name:       "Black Friday",
		Month:      time.November,
		Weekday:    time.Thursday,
		NthWeekday: 4,
		Offset:     1,
		calc:       CalcNthWeekday,
	}
	// ChristmasDay represents Christmas Day on 25-Dec
	USChristmasDay = ChristmasDay.Copy("US Christmas day").SetObservance(nearestWorkday)
)
