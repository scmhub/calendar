package calendar

import "time"

/*
	Market Identifier Code (MIC) (ISO 10383)
	https://www.iso20022.org/market-identifier-codes

*/

// America

// Us Equities Markets
func usEquities(name string, loc *time.Location, years ...int) *Calendar {
	c := NewCalendar(name, loc, years...)
	// Session
	c.SetSession(&Session{
		EarlyOpen:  7 * time.Hour,
		Open:       9*time.Hour + 30*time.Minute,
		Close:      16 * time.Hour,
		EarlyClose: 13 * time.Hour,
		LateClose:  20 * time.Hour,
	})
	// Recurring Holidays
	c.AddHolidays(
		NewYear.Copy().SetObservance(sundayToMonday),
		MLKDay,
		PresidentsDay,
		GoodFriday,
		MemorialDay,
		IndependenceDay,
		LaborDay,
		ThanksgivingDay,
		ChristmasDay.Copy().SetObservance(nearestWorkday),
	)
	// Non Recurring Holidays
	c.AddHolidays(USNationalDaysOfMourning...)
	c.AddHolidays(SeptemberElevenDays...)
	c.AddHolidays(HurricaneSandyDays...)
	// Early Closing
	c.AddEarlyClosingDays(
		BeforeIndependenceDay.Copy().SetObservance(onlyOnWeekdays(time.Monday, time.Tuesday, time.Thursday)),
		AfterIndependenceDay.Copy().SetBeforeYear(2013).SetObservance(onlyOnWeekdays(time.Friday)),
		BeforeIndependenceDay.Copy().SetAfterYear(2013).SetObservance(onlyOnWeekdays(time.Wednesday)),
		BlackFriday,
		ChristmasEve.Copy().SetObservance(exeptOnWeekdays(time.Friday)), // would overlapp Christmas day observance
	)
	return c
}

// New York Stock Exchange
func XNYS(years ...int) *Calendar {
	c := usEquities("New York Stock Exchange", NewYork, years...)
	return c
}

// NASDAQ
func XNAS(years ...int) *Calendar {
	c := usEquities("NASDAQ", NewYork, years...)
	return c
}

// Chicago Board Options Exchange
func XCBO(years ...int) *Calendar {
	c := usEquities("Chicago Board Options Exchange", Chicago, years...)
	return c
}

// Cboe Futures Exchange
func XCBF(years ...int) *Calendar {
	c := usEquities("Cboe Futures Exchange", Chicago, years...)
	return c
}

// Toronto Stock Exchange
func XTSE(years ...int) *Calendar {
	c := NewCalendar("Toronto Stock Exchange", Chicago, years...)
	// Session
	c.SetSession(&Session{
		Open:  9*time.Hour + 30*time.Minute,
		Close: 16 * time.Hour,
	})
	//TODO: add holidays
	return c
}

// Mexican Stock Exchange
func XMEX(years ...int) *Calendar {
	c := NewCalendar("Mexican Stock Exchange", Mexico, years...)
	// Session
	c.SetSession(&Session{
		Open:  8*time.Hour + 30*time.Minute,
		Close: 15 * time.Hour,
	})
	//TODO: add holidays
	return c
}

// Brasilian Stock Exchange - Bolsa de Valores, Mercados e Futuros
func BVMF(years ...int) *Calendar {
	c := NewCalendar("Brasilian Stock Exchange", SaoPaulo, years...)
	// Session
	c.SetSession(&Session{
		Open:  10 * time.Hour,
		Close: 17 * time.Hour,
	})
	//TODO: add holidays
	return c
}

// Europe

// London Stock Exchange
func XLON(years ...int) *Calendar {
	c := NewCalendar("London Stock Exchange", London, years...)
	// Session
	c.SetSession(&Session{
		Open:       8 * time.Hour,
		BreakStart: 12 * time.Hour,
		BreakStop:  12*time.Hour + 2*time.Minute,
		Close:      16*time.Hour + 30*time.Minute,
		EarlyClose: 12*time.Hour + 30*time.Minute,
	})
	// Recurring Holidays
	c.AddHolidays(
		NewYear.Copy().SetObservance(nextMonday),
		GoodFriday,
		EasterMonday,
		EarlyMay,
		LateMay,
		SummerHoliday,
		ChristmasDay.Copy().SetObservance(nextMonday),
		BoxingDay.Copy().SetObservance(nextMonday),
	)
	// Non Recurring Holidays
	c.AddHolidays(GoldenJubileeDays...)
	c.AddHolidays(WilliamWedding)
	c.AddHolidays(DiamondJubileeDays...)
	c.AddHolidays(VEAnniversary)
	// Early Closing
	c.AddEarlyClosingDays(
		ChristmasEve,
		NewYearsEve,
	)
	return c
}

// euronext
func euronext(name string, loc *time.Location, years ...int) *Calendar {
	c := NewCalendar(name, loc, years...)
	// Session
	c.SetSession(&Session{
		Open:       9 * time.Hour,
		Close:      17*time.Hour + 30*time.Minute,
		EarlyClose: 14*time.Hour + 5*time.Minute,
	})
	// Recurring Holidays
	c.AddHolidays(
		NewYear,
		GoodFriday,
		EasterMonday,
		WorkersDay,
		PentecostMonday.Copy().SetBeforeYear(2002),
		ChristmasDay.Copy().SetObservance(nextMonday),
		BoxingDay,
		NewYearsEve.Copy().SetBeforeYear(2002),
	)
	// Early Closing
	c.AddEarlyClosingDays(
		ChristmasEve,
		NewYearsEve.Copy().SetAfterYear(2002),
	)
	return c
}

// Euronext Amsterdam
func XAMS(years ...int) *Calendar {
	c := euronext("Euronext Amsterdam", Amsterdam, years...)
	c.session.Close = 17*time.Hour + 40*time.Minute
	c.AddHolidays(
		QueensDay.Copy().SetBeforeYear(2002),
	)
	return c
}

// Euronext Brussels
func XBRU(years ...int) *Calendar {
	c := euronext("Euronext Brussels", Brussels, years...)
	c.AddHolidays(
		BelgiumIndependenceDay.Copy().SetBeforeYear(2002),
	)
	return c
}

// Euronext Lisbon
func XLIS(years ...int) *Calendar {
	c := euronext("Euronext Lisbon", Lisbon, years...)
	c.AddHolidays(
		Carnival.Copy().SetBeforeYear(2003),
		LibertyDay.Copy().SetBeforeYear(2003),
		CorpusChristi.Copy().SetBeforeYear(2003),
		PortugalDay.Copy().SetBeforeYear(2003),
		SaintAnthonysDay.Copy().SetBeforeYear(2002),
		PortugalRepublicDay.Copy().SetBeforeYear(2003),
	)
	return c
}

// Euronext Paris
func XPAR(years ...int) *Calendar {
	c := euronext("Euronext Paris", Paris, years...)
	c.AddHolidays(
		BastilleDay.Copy().SetBeforeYear(2002),
	)

	return c
}

// Euronext Milan - Borsa Italiana S.P.A
func XMIL(years ...int) *Calendar {
	c := euronext("Euronext Milan", Milan, years...)
	// Session
	c.SetSession(&Session{
		Open:  8 * time.Hour,
		Close: 17*time.Hour + 30*time.Minute,
	})
	c.AddHolidays(
		AssumptionOfMary.Copy("Ferragosto"),
	)
	return c
}

// Madrid Stock Exchange
func XMAD(years ...int) *Calendar {
	c := NewCalendar("Madrid Stock Exchange", Madrid, years...)
	// Session
	c.SetSession(&Session{
		Open:       9 * time.Hour,
		Close:      17*time.Hour + 30*time.Minute,
		EarlyClose: 14 * time.Hour,
	})
	// Recurring Holidays
	c.AddHolidays(
		NewYear,
		Epiphany.Copy().SetBeforeYear(2007),
		GoodFriday,
		EasterMonday,
		WorkersDay,
		AssumptionOfMary.Copy().SetObservance(nextMonday).SetBeforeYear(2005),
		SpainNationalDay.Copy().SetBeforeYear(2005),
		AllSaintsDay.Copy().SetBeforeYear(2005),
		SpainConstitutionDay.Copy().SetBeforeYear(2005),
		ImmaculateConception.Copy().SetBeforeYear(2005),
		ChristmasEve.Copy().SetBeforeYear(2012),
		ChristmasDay.Copy().SetObservance(nextMonday),
		BoxingDay,
		NewYearsEve.Copy().SetBeforeYear(2012),
	)
	// Early Closing
	c.AddEarlyClosingDays(
		ChristmasEve.Copy().SetAfterYear(2012),
		NewYearsEve.Copy().SetAfterYear(2012),
	)
	return c
}

// german
func german(name string, loc *time.Location, years ...int) *Calendar {
	c := NewCalendar(name, loc, years...)
	// Session
	c.SetSession(&Session{
		EarlyOpen:  8 * time.Hour,
		Open:       9 * time.Hour,
		Close:      17*time.Hour + 30*time.Minute,
		EarlyClose: 12*time.Hour + 30*time.Minute,
		LateClose:  20 * time.Hour,
	})
	// Recurring Holidays
	c.AddHolidays(
		NewYear,
		GoodFriday,
		EasterMonday,
		WorkersDay,
		PentecostMonday.Copy().SetAfterYear(2015),
		GermanUnityDay,
		ChristmasEve,
		ChristmasDay.Copy().SetObservance(nextMonday),
		BoxingDay,
		NewYearsEve,
	)
	// Non Recurring Holidays
	c.AddHolidays(
		PentecostMonday.Copy().SetOnYear(2007),
		ReformationDay.Copy("Reformation 500th Anniversary").SetOnYear(2017),
	)
	// Early Closing
	c.AddEarlyClosingDays(
		NewYearsEve.Copy("New Year's Eve Eve").SetObservance(previousWorkday).SetAfterYear(2002),
	)
	return c
}

// Frankfurt Stock Exchange
func XFRA(years ...int) *Calendar {
	c := german("Deutsche Boerse", Franckfurt, years...)
	return c
}

// Deutsche Borse Xetra
func XETR(years ...int) *Calendar {
	c := german("Deutsche Borse Xetra", Franckfurt, years...)
	return c
}

// SIX Group (SWX Swiss Exchange)
func XSWX(years ...int) *Calendar {
	c := NewCalendar("SIX Group", Zurich, years...)
	// Session
	c.SetSession(&Session{
		Open:  9 * time.Hour,
		Close: 17*time.Hour + 30*time.Minute,
	})
	// Recurring Holidays
	c.AddHolidays(
		NewYear,
		BerchtoldsDay,
		GoodFriday,
		EasterMonday,
		WorkersDay,
		AscensionDay,
		PentecostMonday,
		SwissNationalDay,
		ChristmasEve,
		ChristmasDay,
		BoxingDay,
		NewYearsEve,
	)
	return c
}

// Asia/Pacific

// Bombay Stock Exchange
func XBOM(years ...int) *Calendar {
	c := NewCalendar("Bombay Stock Exchange", Mumbai, years...)
	// Session
	c.SetSession(&Session{
		Open:  9 * time.Hour,
		Close: 15*time.Hour + 30*time.Minute,
	})
	//TODO: add holidays
	return c
}

// Stock Exchange of Thailand
func XBKK(years ...int) *Calendar {
	c := NewCalendar("Stock Exchange of Thailand", Bangkok, years...)
	//TODO: add holidays
	return c
}

// Singapore Exchange (SGX)
func XSES(years ...int) *Calendar {
	c := NewCalendar("Singapore Exchange", Singapore, years...)
	// Session
	c.SetSession(&Session{
		EarlyOpen:  8*time.Hour + 30*time.Minute,
		Open:       9 * time.Hour,
		Close:      17 * time.Hour,
		EarlyClose: 12 * time.Hour,
	})
	// Recurring Holidays
	c.AddHolidays(
		NewYear.Copy().SetObservance(sundayToMonday),
		LunarNewYear,
		LunarNewYear.Copy("Lunar New Year's 2nd Day").SetOffset(1),
		GoodFriday,
		WorkersDay,
		EidAlFitr.Copy("Hari Raya Puasa").SetOffset(1),
		VesakDay,
		EidAlAdha.Copy("Hari Raya Haji").SetOffset(1),
		SingaporeNationalDay,
		Deepavali,
		ChristmasDay.Copy().SetObservance(sundayToMonday),
	)
	// Early Closing
	c.AddEarlyClosingDays(
		LunarNewYearEve,
		ChristmasEve,
		NewYearsEve,
	)
	return c
}

// Stock Exchange of Hong Kong
func XHKG(years ...int) *Calendar {
	c := NewCalendar("Stock Exchange of Hong Kong", HongKong, years...)
	// Session
	c.SetSession(&Session{
		EarlyOpen:  9 * time.Hour,
		Open:       9*time.Hour + 30*time.Minute,
		BreakStart: 12 * time.Hour,
		BreakStop:  13 * time.Hour,
		Close:      16 * time.Hour,
		EarlyClose: 12 * time.Hour,
	})
	// Recurring Holidays
	c.AddHolidays(
		NewYear.Copy().SetObservance(sundayToMonday),
		LunarNewYear,
		LunarNewYear.Copy("Lunar New Year's 2nd Day").SetOffset(1),
		LunarNewYear.Copy("Lunar New Year's 3rd Day").SetOffset(2).SetAfterYear(2013),
		LunarNewYear.Copy("Lunar New Year's 4th Day").SetOffset(3).SetAfterYear(2017).SetObservance(onlyOnWeekdays(time.Monday, time.Tuesday)),
		QingmingJie.Copy().SetObservance(sundayToMonday),
		GoodFriday,
		EasterMonday,
		BuddhasBirthday.Copy().SetObservance(sundayToMonday),
		WorkersDay.Copy().SetObservance(sundayToMonday),
		DragonBoatFestival.Copy("Dragon Boat (Tuen Ng) Festival"),
		EstablishmentDay.Copy().SetObservance(sundayToMonday),
		MidAutumnFestival.Copy("The day following the Chinese Mid-Autumn Festival").SetOffset(1),
		ChinaNationalDay.Copy().SetObservance(sundayToMonday),
		DoubleNinthFestival.Copy().SetObservance(sundayToMonday),
		ChristmasDay.Copy().SetObservance(nextMonday),
		BoxingDay.Copy().SetObservance(sundayToTuesday),
	)
	// Early Closing
	c.AddEarlyClosingDays(
		LunarNewYearEve,
		ChristmasEve,
		NewYearsEve,
	)
	return c
}

// Shenzhen Stock Exchange
func XSHE(years ...int) *Calendar {
	c := NewCalendar("Shenzhen Stock Exchange", Shenzhen, years...)
	// Session
	c.SetSession(&Session{
		EarlyOpen:  9*time.Hour + 15*time.Minute,
		Open:       9*time.Hour + 30*time.Minute,
		BreakStart: 11*time.Hour + 30*time.Minute,
		BreakStop:  13 * time.Hour,
		Close:      14*time.Hour + 57*time.Minute,
	})
	// Recurring Holidays
	c.AddHolidays(
		NewYear,
		LunarNewYearEve.Copy().SetAfterYear(2008),
		LunarNewYear,
		LunarNewYear.Copy("Lunar New Year's Day 2").SetOffset(1),
		LunarNewYear.Copy("Lunar New Year's Day 3").SetOffset(2),
		LunarNewYear.Copy("Lunar New Year's Day 4").SetOffset(3),
		LunarNewYear.Copy("Lunar New Year's Day 5").SetOffset(4),
		LunarNewYear.Copy("Lunar New Year's Day 6").SetOffset(5),
		QingmingJie.Copy().SetAfterYear(2008).AddOffset(-1).SetObservance(onlyOnWeekdays(time.Monday)),
		QingmingJie.Copy().SetAfterYear(2008).SetObservance(nextMonday),
		QingmingJie.Copy().SetAfterYear(2008).AddOffset(1).SetObservance(onlyOnWeekdays(time.Friday)),
		WorkersDay.Copy().SetObservance(nextMonday),
		WorkersDay.Copy().SetOffset(1).SetObservance(offsetOverWeekend(1)),
		WorkersDay.Copy().SetOffset(2).SetObservance(offsetOverWeekend(2)),
		DragonBoatFestival.Copy().SetAfterYear(2008).SetOffset(-1).SetObservance(onlyOnWeekdays(time.Monday)),
		DragonBoatFestival.Copy().SetAfterYear(2008).SetObservance(nextMonday),
		DragonBoatFestival.Copy().SetAfterYear(2008).SetOffset(1).SetObservance(onlyOnWeekdays(time.Friday)),
		MidAutumnFestival.Copy().SetAfterYear(2008).SetOffset(-1).SetObservance(onlyOnWeekdays(time.Monday)),
		MidAutumnFestival.Copy().SetAfterYear(2008).SetObservance(nextMonday),
		MidAutumnFestival.Copy().SetAfterYear(2008).SetOffset(1).SetObservance(onlyOnWeekdays(time.Friday)),
		ChinaNationalDay,
		ChinaNationalDay.Copy("China National Day 2").SetOffset(1),
		ChinaNationalDay.Copy("China National Day 3").SetOffset(2),
		ChinaNationalDay.Copy("China National Day 4").SetOffset(3),
		ChinaNationalDay.Copy("China National Day 5").SetOffset(4),
		ChinaNationalDay.Copy("China National Day 6").SetOffset(5),
		ChinaNationalDay.Copy("China National Day 7").SetOffset(6),
	)
	return c
}

// Shanghai Stock Exchange
func XSHG(years ...int) *Calendar {
	c := NewCalendar("Shanghai Stock Exchange", Shanghai, years...)
	// Session
	c.SetSession(&Session{
		EarlyOpen:  9*time.Hour + 15*time.Minute,
		Open:       9*time.Hour + 30*time.Minute,
		BreakStart: 11*time.Hour + 30*time.Minute,
		BreakStop:  13 * time.Hour,
		Close:      15 * time.Hour,
		LateClose:  15*time.Hour + 30*time.Minute,
	})
	// Recurring Holidays
	c.AddHolidays(
		NewYear,
		LunarNewYearEve.Copy().SetAfterYear(2008),
		LunarNewYear,
		LunarNewYear.Copy("Lunar New Year's Day 2").SetOffset(1),
		LunarNewYear.Copy("Lunar New Year's Day 3").SetOffset(2),
		LunarNewYear.Copy("Lunar New Year's Day 4").SetOffset(3),
		LunarNewYear.Copy("Lunar New Year's Day 5").SetOffset(4),
		LunarNewYear.Copy("Lunar New Year's Day 6").SetOffset(5),
		QingmingJie.Copy().SetAfterYear(2008).AddOffset(-1).SetObservance(onlyOnWeekdays(time.Monday)),
		QingmingJie.Copy().SetAfterYear(2008).SetObservance(nextMonday),
		QingmingJie.Copy().SetAfterYear(2008).AddOffset(1).SetObservance(onlyOnWeekdays(time.Friday)),
		WorkersDay.Copy().SetObservance(nextMonday),
		WorkersDay.Copy().SetOffset(1).SetObservance(offsetOverWeekend(1)),
		WorkersDay.Copy().SetOffset(2).SetObservance(offsetOverWeekend(2)),
		DragonBoatFestival.Copy().SetAfterYear(2008).SetOffset(-1).SetObservance(onlyOnWeekdays(time.Monday)),
		DragonBoatFestival.Copy().SetAfterYear(2008).SetObservance(nextMonday),
		DragonBoatFestival.Copy().SetAfterYear(2008).SetOffset(1).SetObservance(onlyOnWeekdays(time.Friday)),
		MidAutumnFestival.Copy().SetAfterYear(2008).SetOffset(-1).SetObservance(onlyOnWeekdays(time.Monday)),
		MidAutumnFestival.Copy().SetAfterYear(2008).SetObservance(nextMonday),
		MidAutumnFestival.Copy().SetAfterYear(2008).SetOffset(1).SetObservance(onlyOnWeekdays(time.Friday)),
		ChinaNationalDay,
		ChinaNationalDay.Copy("China National Day 2").SetOffset(1),
		ChinaNationalDay.Copy("China National Day 3").SetOffset(2),
		ChinaNationalDay.Copy("China National Day 4").SetOffset(3),
		ChinaNationalDay.Copy("China National Day 5").SetOffset(4),
		ChinaNationalDay.Copy("China National Day 6").SetOffset(5),
		ChinaNationalDay.Copy("China National Day 7").SetOffset(6),
	)
	return c
}

// Korea Exchange
func XKRX(years ...int) *Calendar {
	c := NewCalendar("Korea Exchange", Seoul, years...)
	// Session
	c.SetSession(&Session{
		Open:  9 * time.Hour,
		Close: 15*time.Hour + 30*time.Minute,
	})
	//TODO: add holidays
	return c
}

// Japan Exchange Group
func XJPX(years ...int) *Calendar {
	c := NewCalendar("Japan Exchange Group", Tokyo, years...)
	// Session
	c.SetSession(&Session{
		Open:       9 * time.Hour,
		BreakStart: 11*time.Hour + 30*time.Minute,
		BreakStop:  12*time.Hour + 30*time.Minute,
		Close:      15 * time.Hour,
	})
	// Recurring Holidays
	c.AddHolidays(
		NewYear,
		NewYear.Copy().SetOffset(1),
		NewYear.Copy().SetOffset(2),
		ComingOfAgeDay,
		NationalFoundationDay.Copy().SetObservance(sundayToMonday),
		NorthwardEquinox,
		GreeneryDayBefore2007,
		ShowaDay,
		ConstitutionMemorialDay.Copy().SetBeforeYear(2007).SetObservance(sundayToTuesday),
		ConstitutionMemorialDay.Copy().SetAfterYear(2007).SetObservance(sundayToWednesday),
		CitizensHolidayGoldenWeek,
		GreeneryDayAfter2007,
		ChildrensDay.SetObservance(sundayToMonday),
		MarineDayBefore2003.SetObservance(sundayToMonday),
		MarineDay.Copy().SetBeforeYear(2020),
		MarineDay2020,
		MarineDay2021,
		MarineDay.Copy().SetAfterYear(2022),
		MountainDay.SetObservance(sundayToMonday),
		MountainDay2020,
		MountainDay2021,
		MountainDay.SetAfterYear(2022).SetObservance(sundayToMonday),
		SouthwardEquinox,
		RespectForTheAgedDayBefore2003.SetObservance(sundayToMonday),
		RespectForTheAgedDay,
		HealthAndSportsDay.Copy().SetBeforeYear(2020),
		HealthAndSportsDay2020,
		HealthAndSportsDay2021,
		HealthAndSportsDay.SetAfterYear(2022),
		CultureDay,
		LaborThanksgivingDay,
		EmperorAkihitoBirthday,
		EmperorNaruhitoBirthday,
		NewYearsEve,
	)
	// Non Recurring Holidays
	c.AddHolidays(
		AbdicationDay,
		AccessionDay,
		AccessionCitizensHoliday,
		EnthronementCeremony,
	)
	c.AddHolidays(SilverWeekCitizensHolidays...)
	return c
}

// Australian Securities Exchange
func XASX(years ...int) *Calendar {
	c := NewCalendar("Australian Securities Exchange", Sydney, years...)
	// Session
	c.SetSession(&Session{
		Open:  10 * time.Hour,
		Close: 16 * time.Hour,
	})
	//TODO: add holidays
	return c
}
