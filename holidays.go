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
	// Lunar New Year's day
	LunarNewYear = &Holiday{
		Name:  "Lunar New Year's Day",
		Month: time.January,
		Day:   1,
		calc:  CalcLunarDayOfMonth,
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

	// Northward Equinox - March Equinox - ~20-Mar
	NorthwardEquinox = &Holiday{
		Name: "Northward Equinox",
		calc: CalcNorthwardEquinox,
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

	// Northern Solstice - June Solstice - ~20-Jun
	NorthernSolstice = &Holiday{
		Name: "Northern Solstice",
		calc: CalcNorthernSolstice,
	}

	// Assumption of Mary on 15-Aug
	AssumptionOfMary = &Holiday{
		Name:  "Assumption of Mary",
		Month: time.August,
		Day:   15,
		calc:  CalcDayOfMonth,
	}

	// Souththward Equinox - September Equinox - ~20-Sep
	SouthwardEquinox = &Holiday{
		Name: "Souththward Equinox",
		calc: CalcSouthwardEquinox,
	}

	// Reformation Day - 31-Oct-2017
	ReformationDay = &Holiday{
		Name:  "Reformation Day",
		Month: time.October,
		Day:   31,
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

	// Southern Solstice - December Solstice - ~20-Dec
	SouthernSolstice = &Holiday{
		Name: "Southern Solstice",
		calc: CalcSouthernSolstice,
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

	// Special Non Working days for the United States of America.

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

	// Special Non Working days for the UK

	// Queen Elizabeth II Golden Jubilee - 3,4 Jun 2002
	GoldenJubilee = &Holiday{
		Name:   "Golden Jubilee",
		Month:  time.June,
		Day:    3,
		OnYear: 2002,
		calc:   CalcDayOfMonth,
	}
	GoldenJubileeDays = []*Holiday{
		GoldenJubilee,
		GoldenJubilee.Copy("Golden Jubilee day 2").SetOffset(1),
	}

	// Wedding Day of Prince William and Catherine Middleton
	WilliamWedding = &Holiday{
		Name:   "Wedding Day of Prince William and Catherine Middleton",
		Month:  time.April,
		Day:    29,
		OnYear: 2011,
		calc:   CalcDayOfMonth,
	}

	// Queen Elizabeth II Diamond Jubilee - 4,5 Jun 2012
	DiamondJubilee = &Holiday{
		Name:   "Diamond Jubilee",
		Month:  time.June,
		Day:    4,
		OnYear: 2012,
		calc:   CalcDayOfMonth,
	}
	DiamondJubileeDays = []*Holiday{
		DiamondJubilee,
		DiamondJubilee.Copy("Diamond Jubilee day 2").SetOffset(1),
	}

	// VE-Day 75th Anniversary
	VEAnniversary = &Holiday{
		Name:   "VE-Day 75th Anniversary",
		Month:  time.May,
		Day:    8,
		OnYear: 2020,
		calc:   CalcDayOfMonth,
	}
)

// Netherlands Holidays

var (
	// Koninginnedag (Koningsdag King's day since 2013) - 30-Apr
	QueensDay = &Holiday{
		Name:  "Queen's Day",
		Month: time.April,
		Day:   30,
		calc:  CalcDayOfMonth,
	}
)

// Belgium Holidays

var (
	// Belgium Independence Day - 21-Jul
	BelgiumIndependenceDay = &Holiday{
		Name:  "Belgium Independence Day",
		Month: time.July,
		Day:   21,
		calc:  CalcDayOfMonth,
	}
)

// Portugal Holidays

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

// France Holidays

var (
	// Bastille Day - 14-Jul
	BastilleDay = &Holiday{
		Name:  "Bastille Day",
		Month: time.July,
		Day:   14,
		calc:  CalcDayOfMonth,
	}
)

// Spain Holidays

var (
	// National Day of Spain - 5-Oct
	SpainNationalDay = &Holiday{
		Name:  "National Day of Spain",
		Month: time.October,
		Day:   12,
		calc:  CalcDayOfMonth,
	}
	// Constitution Day - 6-Dec
	SpainConstitutionDay = &Holiday{
		Name:  "Constitution Day",
		Month: time.December,
		Day:   6,
		calc:  CalcDayOfMonth,
	}
)

// Germany Holidays

var (
	// German Unity Day - 3-Oct
	GermanUnityDay = &Holiday{
		Name:  "German Unity Day",
		Month: time.October,
		Day:   3,
		calc:  CalcDayOfMonth,
	}
)

// Swiss Holidays

var (
	// Berchtold's Day
	BerchtoldsDay = &Holiday{
		Name:  "Berchtold's Day",
		Month: time.January,
		Day:   2,
		calc:  CalcDayOfMonth,
	}
	// Swiss National Day
	SwissNationalDay = &Holiday{
		Name:  "Swiss National Day",
		Month: time.August,
		Day:   1,
		calc:  CalcDayOfMonth,
	}
)

// China Holidays
var (
	// Quingming Festival
	QingmingJie = &Holiday{
		Name:  "Quingming Festival",
		Month: time.April,
		Day:   5,
		calc:  CalcDayOfMonth,
	}
	// Dragon Boat Festival day - 5th day of 5th lunar month
	DragonBoatFestival = &Holiday{
		Name:  "Dragon Boat Festival",
		Month: time.Month(5),
		Day:   5,
		calc:  CalcLunarDayOfMonth,
	}
	// People's Republic of China National day - 1-Oct
	ChinaNationalDay = &Holiday{
		Name:  "China National Day",
		Month: time.October,
		Day:   1,
		calc:  CalcDayOfMonth,
	}
	// Mid-Autumn Festival 中秋節 - 15yh day of the 8th lunar month
	MidAutumnFestival = &Holiday{
		Name:  "Dragon Boat Festival",
		Month: time.Month(8),
		Day:   15,
		calc:  CalcLunarDayOfMonth,
	}
)

// Japan Holidays

var (
	// Coming of Age Day - Seijin no hi - 成人の日 - 2nd monday of January
	ComingOfAgeDay = &Holiday{
		Name:       "Coming of Age Day",
		Month:      time.January,
		Weekday:    time.Monday,
		NthWeekday: 2,
		calc:       CalcNthWeekday,
	}
	// National Foundation Day - Kenkoku kinen no hi - 建国記念の日 - 11-Feb
	NationalFoundationDay = &Holiday{
		Name:  "National Foundation Day",
		Month: time.February,
		Day:   11,
		calc:  CalcDayOfMonth,
	}
	// Greenery Day Before 2007 - Midori no hi - みどりの日 - 29-Apr before 2007
	GreeneryDayBefore2007 = &Holiday{
		Name:       "Greenery Day",
		Month:      time.April,
		Day:        29,
		BeforeYear: 2007,
		calc:       CalcDayOfMonth,
	}
	// Showa Day - Shōwa no hi - 昭和の日 - 29-Apr after 2007
	ShowaDay = &Holiday{
		Name:      "Showa Day",
		Month:     time.April,
		Day:       29,
		AfterYear: 2007,
		calc:      CalcDayOfMonth,
	}
	// Constitution Memorial Day - Kenpō kinen bi - 憲法記念日 - 3-May
	ConstitutionMemorialDay = &Holiday{
		Name:  "Constitution Memorial Day",
		Month: time.May,
		Day:   3,
		calc:  CalcDayOfMonth,
	}
	// Citizen's Holiday ot the Golden Week before 2007 - Kokumin no kyūjitsu - 国民の休日 - 4-May before 2007
	CitizensHolidayGoldenWeek = &Holiday{
		Name:       "Citizen's Holiday of the Golden Week",
		Month:      time.May,
		Day:        4,
		BeforeYear: 2007,
		calc:       CalcDayOfMonth,
	}
	// Greenery Day After 2007 - Midori no hi - みどりの日 - 4-May after 2007
	GreeneryDayAfter2007 = &Holiday{
		Name:      "Greenery Day",
		Month:     time.May,
		Day:       4,
		AfterYear: 2007,
		calc:      CalcDayOfMonth,
	}
	// Children's Day - Kodomo no hi - こどもの日 - 5-May
	ChildrensDay = &Holiday{
		Name:  "Children's Day",
		Month: time.May,
		Day:   5,
		calc:  CalcDayOfMonth,
	}
	// Marine Day - Umi no hi - 海の日 - 3rd monday of July
	MarineDay = &Holiday{
		Name:       "Marine Day",
		Month:      time.July,
		Weekday:    time.Monday,
		NthWeekday: 3,
		AfterYear:  2003,
		calc:       CalcNthWeekday,
	}
	// Marine Day before 2003 - Umi no hi - 海の日 - 20-July
	MarineDayBefore2003 = &Holiday{
		Name:       "Marine Day before 2003",
		Month:      time.July,
		Day:        23,
		BeforeYear: 2003,
		calc:       CalcDayOfMonth,
	}
	// Marine Day 2020- Umi no hi - 海の日 - Changed for Tokyo Olympics - 23-Jul-2020
	MarineDay2020 = &Holiday{
		Name:   "Marine Day 2020",
		Month:  time.July,
		Day:    23,
		OnYear: 2020,
		calc:   CalcDayOfMonth,
	}
	// Marine Day 2021- Umi no hi - 海の日 - Changed for Tokyo Olympics (2020 postponed) - 22-Jul-2021
	MarineDay2021 = &Holiday{
		Name:   "Marine Day 2021",
		Month:  time.July,
		Day:    22,
		OnYear: 2021,
		calc:   CalcDayOfMonth,
	}
	// Mountain Day - Yama no hi - 山の日 - 3rd monday of July
	MountainDay = &Holiday{
		Name:      "Mountain Day",
		Month:     time.August,
		Day:       11,
		AfterYear: 2016,
		calc:      CalcDayOfMonth,
	}
	// Mountain Day 2020 - Yama no hi - 山の日 - Changed for Tokyo Olympics - 10-Aug-2020
	MountainDay2020 = &Holiday{
		Name:   "Mountain Day 2020",
		Month:  time.August,
		Day:    10,
		OnYear: 2020,
		calc:   CalcDayOfMonth,
	}
	// Mountain Day 2021 - Yama no hi - 山の日 - Changed for Tokyo Olympics (2020 postponed) - 9-Aug-2021
	MountainDay2021 = &Holiday{
		Name:   "Mountain Day 2021",
		Month:  time.August,
		Day:    9,
		OnYear: 2021,
		calc:   CalcDayOfMonth,
	}
	// Respect for the Aged Day - Keirō no hi - 敬老の日 - 3rd monday of September after 2003
	RespectForTheAgedDay = &Holiday{
		Name:       "Respect for the Aged Day",
		Month:      time.September,
		Weekday:    time.Monday,
		NthWeekday: 3,
		AfterYear:  2003,
		calc:       CalcNthWeekday,
	}
	// Respect for the Aged Day before 2003 - Keirō no hi - 敬老の日 - 15-Sep before 2003
	RespectForTheAgedDayBefore2003 = &Holiday{
		Name:       "Respect for the Aged Day before 2003",
		Month:      time.September,
		Day:        15,
		NthWeekday: 3,
		BeforeYear: 2003,
		calc:       CalcDayOfMonth,
	}
	// Health and Sports Day - Supōtsu no hi - スポーツの日 - 2nd monday of October
	HealthAndSportsDay = &Holiday{
		Name:       "Health and Sports Day",
		Month:      time.October,
		Weekday:    time.Monday,
		NthWeekday: 2,
		calc:       CalcNthWeekday,
	}
	// Health and Sports Day 2020 - Changed for Tokyo Olympics - 24-Jul-2020
	HealthAndSportsDay2020 = &Holiday{
		Name:   "Health and Sports Day",
		Month:  time.July,
		Day:    24,
		OnYear: 2020,
		calc:   CalcDayOfMonth,
	}
	// Health and Sports Day 2021 - Changed for Tokyo Olympics (2020 postponed) - 23-Jul-2021
	HealthAndSportsDay2021 = &Holiday{
		Name:   "Health and Sports Day",
		Month:  time.July,
		Day:    23,
		OnYear: 2021,
		calc:   CalcDayOfMonth,
	}
	// Culture Day - Bunka no hi - 文化の日 - 3-Nov
	CultureDay = &Holiday{
		Name:  "Culture Day",
		Month: time.November,
		Day:   3,
		calc:  CalcDayOfMonth,
	}
	// Labor Thanksgiving Day - Kinrō Kansha no hi - 勤労感謝の日 - 23-Nov
	LaborThanksgivingDay = &Holiday{
		Name:  "Labor Thanksgiving Day",
		Month: time.November,
		Day:   23,
		calc:  CalcDayOfMonth,
	}
	// Emperor Akihito's Birthday - Tennō tanjōbi - 天皇誕生日 - 23-Dec before 30-Apr-2019
	EmperorAkihitoBirthday = &Holiday{
		Name:       "Emperor Akihito's Birthday",
		Month:      time.December,
		Day:        23,
		BeforeYear: 2019,
		calc:       CalcDayOfMonth,
	}
	// Emperor Naruhito's Birthday - Tennō tanjōbi - 天皇誕生日 - 23-Feb before 30-Apr-2019
	EmperorNaruhitoBirthday = &Holiday{
		Name:      "Emperor Naruhito's Birthday",
		Month:     time.February,
		Day:       23,
		AfterYear: 2020,
		calc:      CalcDayOfMonth,
	}

	// Special Non Working days for Japan

	// Emperor Akihito's Abdication Day - Taiirei-Seiden-no-gi - 退位礼正殿の儀 - 30-Apr-2019
	AbdicationDay = &Holiday{
		Name:   "Emperor Akihito's Abdication Day",
		Month:  time.April,
		Day:    30,
		OnYear: 2019,
		calc:   CalcDayOfMonth,
	}
	// Beginning of the Reiwa era and accession date of Emperor Naruhito - 1-May-2019
	AccessionDay = &Holiday{
		Name:   "Naruhito's Accession Day",
		Month:  time.May,
		Day:    1,
		OnYear: 2019,
		calc:   CalcDayOfMonth,
	}
	// Accession Citizen's Holiday - 2-May-2019
	AccessionCitizensHoliday = &Holiday{
		Name:   "Accession Citizen's Holiday",
		Month:  time.May,
		Day:    2,
		OnYear: 2019,
		calc:   CalcDayOfMonth,
	}
	// Enthronement Ceremony - Sokuirei-Seiden-no-gi - 即位礼正殿の儀 - 22-Oct-2019
	EnthronementCeremony = &Holiday{
		Name:   "Enthronement Ceremony",
		Month:  time.October,
		Day:    22,
		OnYear: 2019,
		calc:   CalcDayOfMonth,
	}
	// Silver Week Citizen's Holidays on 22nd of September
	SilverWeekCitizensHoliday = &Holiday{
		Name:   "Silver Week Citizen's Holidays",
		Month:  time.September,
		Day:    22,
		OnYear: 2009,
		calc:   CalcDayOfMonth,
	}
	// Silver Week Citizen's Holidays on 21st of September
	SilverWeekCitizensHoliday21 = &Holiday{
		Name:   "Silver Week Citizen's Holidays on the 21st",
		Month:  time.September,
		Day:    21,
		OnYear: 2032,
		calc:   CalcDayOfMonth,
	}
	SilverWeekCitizensHolidays = []*Holiday{
		SilverWeekCitizensHoliday,
		SilverWeekCitizensHoliday.Copy().SetOnYear(2015),
		SilverWeekCitizensHoliday.Copy().SetOnYear(2020),
		SilverWeekCitizensHoliday.Copy().SetOnYear(2026),
		SilverWeekCitizensHoliday21,
		SilverWeekCitizensHoliday.Copy().SetOnYear(2037),
		SilverWeekCitizensHoliday.Copy().SetOnYear(2043),
		SilverWeekCitizensHoliday21.Copy().SetOnYear(2049),
		SilverWeekCitizensHoliday.Copy().SetOnYear(2054),
		SilverWeekCitizensHoliday21.Copy().SetOnYear(2060),
		SilverWeekCitizensHoliday.Copy().SetOnYear(2071),
		SilverWeekCitizensHoliday21.Copy().SetOnYear(2077),
		SilverWeekCitizensHoliday21.Copy().SetOnYear(2088),
		SilverWeekCitizensHoliday21.Copy().SetOnYear(2094),
		SilverWeekCitizensHoliday.Copy().SetOnYear(2099),
	}
)
