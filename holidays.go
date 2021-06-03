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

	// Carnival, Mardi Gras or Fat Tuesday  - 47 days before Easter
	Carnival = &Holiday{
		Name:   "Carnival",
		Offset: -47,
		calc:   CalcEasterOffset,
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

	// Pentecost Monday (or Whit Monday) on the day after Pentecost (50 days after Easter)
	PentecostMonday = &Holiday{
		Name:   "Pentecost Monday",
		Offset: 50,
		calc:   CalcEasterOffset,
	}

	// Corpus Christi 60 days after Easter
	CorpusChristi = &Holiday{
		Name:   "Corpus Christi",
		Offset: 60,
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
		Observance: nearestWorkday,
		calc:       CalcDayOfMonth,
	}

	// Day Before Independence Day
	BeforeIndependenceDay = &Holiday{
		Name:  "Day before Independence Day",
		Month: time.July,
		Day:   3,
		calc:  CalcDayOfMonth,
	}
	// Day Before Independence Day
	AfterIndependenceDay = &Holiday{
		Name:  "Day after Independence Day",
		Month: time.July,
		Day:   5,
		calc:  CalcDayOfMonth,
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

// Special Non Working days for the United States of America.
var (
	// President Richard Nixon - April 27, 1994
	NixonMourningDay = &Holiday{
		Name:   "President Richard Nixon Mourning Day",
		Month:  time.April,
		Day:    27,
		OnYear: 1994,
		calc:   CalcDayOfMonth,
	}
	// President Ronald W. Reagan - June 11, 2004
	ReaganMourningDay = &Holiday{
		Name:   "President Ronald W. Reagan Mourning Day",
		Month:  time.June,
		Day:    11,
		OnYear: 2004,
		calc:   CalcDayOfMonth,
	}
	// President Gerald R. Ford - Jan 2, 2007
	FordMourningDay = &Holiday{
		Name:   "President Gerald R. Ford Mourning Day",
		Month:  time.January,
		Day:    2,
		OnYear: 2007,
		calc:   CalcDayOfMonth,
	}
	// President George Bush Senior - Nov 30, 2018
	BushSeniorMourningDay = &Holiday{
		Name:   "President George Bush Senior Mourning Day",
		Month:  time.November,
		Day:    30,
		OnYear: 2018,
		calc:   CalcDayOfMonth,
	}
	// National days of Mourning for the United States of America
	USNationalDaysOfMourning = []*Holiday{
		NixonMourningDay,
		ReaganMourningDay,
		FordMourningDay,
		BushSeniorMourningDay,
	}

	// September 11 - september 11, 2001
	SeptemberEleven = &Holiday{
		Name:   "Sepember 11",
		Month:  time.September,
		Day:    11,
		OnYear: 2001,
		calc:   CalcDayOfMonth,
	}

	// September 11 -14 range
	SeptemberElevenDays = []*Holiday{
		SeptemberEleven,
		SeptemberEleven.Copy("Sepember 11 day 2").SetOffset(1),
		SeptemberEleven.Copy("Sepember 11 day 3").SetOffset(2),
		SeptemberEleven.Copy("Sepember 11 day 4").SetOffset(3),
	}

	// Hurricane Sandy - october 29, 2012
	HurricaneSandy = &Holiday{
		Name:   "Hurricane Sandy",
		Month:  time.October,
		Day:    29,
		OnYear: 2012,
		calc:   CalcDayOfMonth,
	}
	HurricaneSandyDays = []*Holiday{
		HurricaneSandy,
		HurricaneSandy.Copy("Hurricane Sandy day 2").SetOffset(1),
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

// Amsterdam Holiday

var (
	// Koninginnedag (Koningsdag King's day since 2013) - 30-Apr
	QueensDay = &Holiday{
		Name:  "Queen's Day",
		Month: time.April,
		Day:   30,
		calc:  CalcDayOfMonth,
	}
)

// Brussels Holiday

var (
	// Belgium Independence Day - 21-Jul
	BelgiumIndependenceDay = &Holiday{
		Name:  "Belgium Independence Day",
		Month: time.July,
		Day:   21,
		calc:  CalcDayOfMonth,
	}
)

// Lisbon Holiday

var (
	// Liberty Day - 25-Apr
	LibertyDay = &Holiday{
		Name:  "Liberty Day",
		Month: time.April,
		Day:   25,
		calc:  CalcDayOfMonth,
	}
	// Portugal Day - 10-Jun
	PortugalDay = &Holiday{
		Name:  "Portugal Republic Day",
		Month: time.October,
		Day:   5,
		calc:  CalcDayOfMonth,
	}
	// Portugal Independence Day - 13-Jun
	SaintAnthonysDay = &Holiday{
		Name:  "Saint Anthony's Day",
		Month: time.December,
		Day:   1,
		calc:  CalcDayOfMonth,
	}
	// Portugal Republic Day - 5-Oct
	PortugalRepublicDay = &Holiday{
		Name:  "Portugal REpublic Day",
		Month: time.October,
		Day:   5,
		calc:  CalcDayOfMonth,
	}
	// Portugal Independence Day - 1-Dec
	PortugalIndependenceDay = &Holiday{
		Name:  "Portugal Independence Day",
		Month: time.December,
		Day:   1,
		calc:  CalcDayOfMonth,
	}
)

// Paris Holidays

var (
	// Bastille Day - 14-Jul
	BastilleDay = &Holiday{
		Name:  "Bastille Day",
		Month: time.July,
		Day:   14,
		calc:  CalcDayOfMonth,
	}
)
