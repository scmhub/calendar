package calendar

import "time"

// Common Holidays

var (
	// New Year's day on 1-Jan
	NewYear = &Holiday{
		Name:  "New Year's Day",
		Month: time.January,
		Day:   1,
		calc:  CalcDayOfMonth,
	}

	// Epiphany on 6-Jan
	Epiphany = &Holiday{
		Name:  "Epiphany",
		Month: time.January,
		Day:   6,
		calc:  CalcDayOfMonth,
	}

	// Maundy Thursday - three days before Easter
	MaundyThursday = &Holiday{
		Name:   "Maundy Thursday",
		Offset: -3,
		calc:   CalcEasterOffset,
	}

	// Good Friday - two days before Easter
	GoodFriday = &Holiday{
		Name:   "Good Friday",
		Offset: -2,
		calc:   CalcEasterOffset,
	}

	// Easter Monday - the day after Easter
	EasterMonday = &Holiday{
		Name:   "Easter Monday",
		Offset: 1,
		calc:   CalcEasterOffset,
	}

	// WorkersDay is the International Workers' Day on 1-May
	WorkersDay = &Holiday{
		Name:  "International Workers' Day",
		Month: time.May,
		Day:   1,
		calc:  CalcDayOfMonth,
	}

	// Ascension Day on the 39th day after Easter
	AscensionDay = &Holiday{
		Name:   "Ascension Day",
		Offset: 39,
		calc:   CalcEasterOffset,
	}

	// Pentecost Monday on the day after Pentecost (50 days after Easter)
	PentecostMonday = &Holiday{
		Name:   "Pentecost Monday",
		Offset: 50,
		calc:   CalcEasterOffset,
	}

	// Assumption of Mary on 15-Aug
	AssumptionOfMary = &Holiday{
		Name:  "Assumption of Mary",
		Month: time.August,
		Day:   15,
		calc:  CalcDayOfMonth,
	}

	// All Saints' Day on 1-Nov
	AllSaintsDay = &Holiday{
		Name:  "All Saints' Day",
		Month: time.November,
		Day:   1,
		calc:  CalcDayOfMonth,
	}

	// Armistice Day on 11-Nov
	ArmisticeDay = &Holiday{
		Name:  "Armistice Day",
		Month: time.November,
		Day:   11,
		calc:  CalcDayOfMonth,
	}

	// Immaculate Conception on 8-Dec
	ImmaculateConception = &Holiday{
		Name:  "Immaculate Conception",
		Month: time.December,
		Day:   8,
		calc:  CalcDayOfMonth,
	}

	// Christmas Eve on 24-Dec
	ChristmasEve = &Holiday{
		Name:  "Christmas Eve",
		Month: time.December,
		Day:   24,
		calc:  CalcDayOfMonth,
	}

	// Christmas Day on 25-Dec
	ChristmasDay = &Holiday{
		Name:  "Christmas Day",
		Month: time.December,
		Day:   25,
		calc:  CalcDayOfMonth,
	}

	// BoxingDay is the day after Christmas (Boxing Day / St. Stephen's Day) on 26-Dec
	BoxingDay = &Holiday{
		Name:  "Boxing Day",
		Month: time.December,
		Day:   26,
		calc:  CalcDayOfMonth,
	}

	// New Year's Eve on 31-Dec
	NewYearsEve = &Holiday{
		Name:  "New Year's Eve",
		Month: time.December,
		Day:   31,
		calc:  CalcDayOfMonth,
	}
)

// US Holidays

var (
	// Martin Luther King Jr. Day on the 3rd Monday in January
	MLKDay = &Holiday{
		Name:       "Martin Luther King Jr. Day",
		Month:      time.January,
		Weekday:    time.Monday,
		NthWeekday: 3,
		calc:       CalcNthWeekday,
	}
	// Presidents' Day or Washington's Birthday on the 3rd Monday in February
	PresidentsDay = &Holiday{
		Name:       "Presidents' Day",
		Month:      time.February,
		Weekday:    time.Monday,
		NthWeekday: 3,
		calc:       CalcNthWeekday,
	}
	// Memorial Day on the last Monday in May
	MemorialDay = &Holiday{
		Name:       "Memorial Day",
		Month:      time.May,
		Weekday:    time.Monday,
		NthWeekday: -1,
		calc:       CalcNthWeekday,
	}
	// Independence Day on 4-Jul
	IndependenceDay = &Holiday{
		Name:       "Independence Day",
		Month:      time.July,
		Day:        4,
		observance: nearestWorkday,
		calc:       CalcDayOfMonth,
	}

	// Day Before Independence Day
	BeforeIndependenceDay = &Holiday{
		Name:       "Day before Independence Day",
		Month:      time.July,
		Day:        3,
		observance: onlyOnWeekdays(time.Monday, time.Tuesday, time.Thursday),
		calc:       CalcDayOfMonth,
	}
	// Day Before Independence Day
	AfterIndependenceDay = &Holiday{
		Name:       "Day after Independence Day",
		Month:      time.July,
		Day:        5,
		observance: onlyOnWeekdays(time.Friday),
		calc:       CalcDayOfMonth,
	}

	// Labor Day on the first Monday in September
	LaborDay = &Holiday{
		Name:       "Labor Day",
		Month:      time.September,
		Weekday:    time.Monday,
		NthWeekday: 1,
		calc:       CalcNthWeekday,
	}

	// Columbus Day on the second Monday in October
	ColumbusDay = &Holiday{
		Name:       "Columbus Day",
		Month:      time.October,
		Weekday:    time.Monday,
		NthWeekday: 2,
		calc:       CalcNthWeekday,
	}

	// Veterans Day on 11-Nov
	VeteransDay = &Holiday{
		Name:  "Veterans Day",
		Month: time.November,
		Day:   11,
		calc:  CalcDayOfMonth,
	}

	// Thanksgiving Day on the fourth Thursday in November
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
)

// UK Holidays

var (
	// Early May or MAy Day on the first monday of May
	EarlyMay = &Holiday{
		Name:       "Early May",
		Month:      time.May,
		Weekday:    time.Monday,
		NthWeekday: 1,
		calc:       CalcNthWeekday,
	}

	// Late May on the last monday of May
	LateMay = &Holiday{
		Name:       "May Day",
		Month:      time.May,
		Weekday:    time.Monday,
		NthWeekday: -1,
		calc:       CalcNthWeekday,
	}

	// Late May on the last monday of May
	SummerHoliday = &Holiday{
		Name:       "Summer Bank Holiday",
		Month:      time.August,
		Weekday:    time.Monday,
		NthWeekday: -1,
		calc:       CalcNthWeekday,
	}
)
