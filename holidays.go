package calendar

import "time"

var (
	NewYear = &Holiday{
		Name:  "New Year's Day",
		Month: time.January,
		Day:   1,
		calc:  CalcDayOfMonth,
	}
	// Epiphany represents Epiphany on 6-Jan
	Epiphany = &Holiday{
		Name:  "Epiphany",
		Month: time.January,
		Day:   6,
		calc:  CalcDayOfMonth,
	}

	// MaundyThursday represents Maundy Thursday - three days before Easter
	MaundyThursday = &Holiday{
		Name:   "Maundy Thursday",
		Offset: -3,
		calc:   CalcEasterOffset,
	}

	// GoodFriday represents Good Friday - two days before Easter
	GoodFriday = &Holiday{
		Name:   "Good Friday",
		Offset: -2,
		calc:   CalcEasterOffset,
	}

	// EasterMonday represents Easter Monday - the day after Easter
	EasterMonday = &Holiday{
		Name:   "Easter Monday",
		Offset: 1,
		calc:   CalcEasterOffset,
	}

	// WorkersDay represents International Workers' Day on 1-May
	WorkersDay = &Holiday{
		Name:  "International Workers' Day",
		Month: time.May,
		Day:   1,
		calc:  CalcDayOfMonth,
	}

	// AscensionDay represents Ascension Day on the 39th day after Easter
	AscensionDay = &Holiday{
		Name:   "Ascension Day",
		Offset: 39,
		calc:   CalcEasterOffset,
	}

	// PentecostMonday represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	PentecostMonday = &Holiday{
		Name:   "Pentecost Monday",
		Offset: 50,
		calc:   CalcEasterOffset,
	}

	// AssumptionOfMary represents Assumption of Mary on 15-Aug
	AssumptionOfMary = &Holiday{
		Name:  "Assumption of Mary",
		Month: time.August,
		Day:   15,
		calc:  CalcDayOfMonth,
	}

	// AllSaintsDay represents All Saints' Day on 1-Nov
	AllSaintsDay = &Holiday{
		Name:  "All Saints' Day",
		Month: time.November,
		Day:   1,
		calc:  CalcDayOfMonth,
	}

	// ArmisticeDay represents Armistice Day on 11-Nov
	ArmisticeDay = &Holiday{
		Name:  "Armistice Day",
		Month: time.November,
		Day:   11,
		calc:  CalcDayOfMonth,
	}

	// ImmaculateConception represents Immaculate Conception on 8-Dec
	ImmaculateConception = &Holiday{
		Name:  "Immaculate Conception",
		Month: time.December,
		Day:   8,
		calc:  CalcDayOfMonth,
	}

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = &Holiday{
		Name:  "Christmas Day",
		Month: time.December,
		Day:   25,
		calc:  CalcDayOfMonth,
	}

	// BoxingDay represents the day after Christmas (Boxing Day / St. Stephen's Day) on 26-Dec
	BoxingDay = &Holiday{
		Name:  "Boxing Day",
		Month: time.December,
		Day:   26,
		calc:  CalcDayOfMonth,
	}
)
