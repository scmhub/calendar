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
		ChristmasEve,
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
	})
	c.AddHolidays(
		NewYear.Copy().SetObservance(nextMonday),
		GoodFriday,
		EasterMonday,
		EarlyMay,
		LateMay,
		SummerHoliday,
		ChristmasEve, // early-closing
		ChristmasDay.Copy().SetObservance(nextMonday),
		BoxingDay.Copy().SetObservance(nextMonday), // PB!!!! if christmas on saturday
		NewYearsEve, // early-closing
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
	c.AddHolidays(
		NewYear,
		GoodFriday,
		EasterMonday,
		PentecostMonday.Copy().SetBeforeYear(2002),
		WorkersDay,
		ChristmasDay.Copy().SetObservance(nextMonday),
		BoxingDay,
		NewYearsEve.Copy().SetBeforeYear(2002),
	)
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
	c := NewCalendar("Euronext Milan", Milan, years...)
	// Session
	c.SetSession(&Session{
		Open:  8 * time.Hour,
		Close: 17*time.Hour + 30*time.Minute,
	})
	//TODO: add holidays
	return c
}

// Madrid Stock Exchange
func XMAD(years ...int) *Calendar {
	c := NewCalendar("Madrid Stock Exchange", Madrid, years...)
	// Session
	c.SetSession(&Session{
		Open:  9 * time.Hour,
		Close: 17*time.Hour + 30*time.Minute,
	})
	//TODO: add holidays
	return c
}

// Frankfurt Stock Exchange
func XFRA(years ...int) *Calendar {
	c := NewCalendar("Deutsche Boerse", Franckfurt, years...)
	// Session
	c.SetSession(&Session{
		EarlyOpen: 8 * time.Hour,
		Open:      9 * time.Hour,
		Close:     17*time.Hour + 30*time.Minute,
		LateClose: 20 * time.Hour,
	})
	//TODO: add holidays
	return c
}

// Deutsche Borse Xetra
func XETR(years ...int) *Calendar {
	c := NewCalendar("Deutsche Borse Xetra", Franckfurt, years...)
	// Session
	c.SetSession(&Session{
		EarlyOpen: 8 * time.Hour,
		Open:      9 * time.Hour,
		Close:     17*time.Hour + 30*time.Minute,
		LateClose: 20 * time.Hour,
	})
	//TODO: add holidays
	return c
}

// SIX Group (SWX Swiss Exchange)
func XSWX(years ...int) *Calendar {
	c := NewCalendar("SIX Group", Zurich, years...)
	//TODO: add holidays
	return c
}

// Asia/Pacific

// Bombay Stock Exchange
func XBOM(years ...int) *Calendar {
	c := NewCalendar("Bombay Stock Exchange", Mumbai, years...)
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
	//TODO: add holidays
	return c
}

// Stock Exchange of Hong Kong
func XHKG(years ...int) *Calendar {
	c := NewCalendar("Stock Exchange of Hong Kong", HongKong, years...)
	// Session
	c.SetSession(&Session{
		Open:       9*time.Hour + 30*time.Minute,
		BreakStart: 12 * time.Hour,
		BreakStop:  13 * time.Hour,
		Close:      16 * time.Hour,
	})
	//TODO: add holidays
	return c
}

// Shenzhen Stock Exchange
func XSHE(years ...int) *Calendar {
	c := NewCalendar("Shenzhen Stock Exchange", Shenzhen, years...)
	// Session
	c.SetSession(&Session{
		Open:       9*time.Hour + 30*time.Minute,
		BreakStart: 11*time.Hour + 30*time.Minute,
		BreakStop:  13 * time.Hour,
		Close:      16 * time.Hour,
	})
	//TODO: add holidays
	return c
}

// Shanghai Stock Exchange
func XSHG(years ...int) *Calendar {
	c := NewCalendar("Shanghai Stock Exchange", Shanghai, years...)
	// Session
	c.SetSession(&Session{
		Open:       9*time.Hour + 30*time.Minute,
		BreakStart: 11*time.Hour + 30*time.Minute,
		BreakStop:  13 * time.Hour,
		Close:      16 * time.Hour,
	})
	//TODO: add holidays
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
	//TODO: add holidays
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
