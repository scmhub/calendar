package calendar

import "time"

/*
	Market Identifier Code (MIC) (ISO 10383)
	https://www.iso20022.org/market-identifier-codes

*/

// America

// New York Stock Exchange
func XNYS(years ...int) *Calendar {
	c := NewCalendar("New York Stock Exchange", NewYork, years...)
	c.SetEarlySession(&Session{7 * time.Hour, 9*time.Hour + 30*time.Minute})
	c.SetCoreSessions(&Session{9*time.Hour + 30*time.Minute, 16 * time.Hour})
	c.SetLateSession(&Session{16 * time.Hour, 20 * time.Hour})
	c.SetEarlyClosing(13 * time.Hour)
	c.AddHolidays(
		NewYear.Copy("New Year's Day").SetObservance(sundayToMonday),
		MLKDay,
		PresidentsDay,
		GoodFriday,
		MemorialDay,
		IndependenceDay, // !! half monday if tuesday
		LaborDay,
		ThanksgivingDay,
		BlackFriday, // !! early-closing 1 pm
		ChristmasDay.Copy("Christmas Day").SetObservance(nearestWorkday),
	)
	// c.AddEarlyClose(
	// 	BeforeIndependenceDay,
	// 	AfterIndependenceDay,
	// )
	return c
}

// NASDAQ
func XNAS(years ...int) *Calendar {
	c := NewCalendar("NASDAQ", NewYork, years...)
	//TODO: add holidays
	return c
}

// Chicago Board Options Exchange
func XCBO(years ...int) *Calendar {
	c := NewCalendar("Chicago Board Options Exchange", Chicago, years...)
	//TODO: add holidays
	return c
}

// Cboe Futures Exchange
func XCBF(years ...int) *Calendar {
	c := NewCalendar("Cboe Futures Exchange", Chicago, years...)
	//TODO: add holidays
	return c
}

// Toronto Stock Exchange
func XTSE(years ...int) *Calendar {
	c := NewCalendar("Toronto Stock Exchange", Chicago, years...)
	//TODO: add holidays
	return c
}

// Mexican Stock Exchange
func XMEX(years ...int) *Calendar {
	c := NewCalendar("Mexican Stock Exchange", Mexico, years...)
	//TODO: add holidays
	return c
}

// Brasilian Stock Exchange - Bolsa de Valores, Mercados e Futuros
func BVMF(years ...int) *Calendar {
	c := NewCalendar("Brasilian Stock Exchange", SaoPaulo, years...)
	//TODO: add holidays
	return c
}

// Europe

// London Stock Exchange
func XLON(years ...int) *Calendar {
	c := NewCalendar("London Stock Exchange", London, years...)
	c.AddHolidays(
		NewYear.Copy("New Year's Day").SetObservance(nextMonday),
		GoodFriday,
		EasterMonday,
		EarlyMay,
		LateMay,
		SummerHoliday,
		ChristmasEve, // early-closing
		ChristmasDay.Copy("Christmas Day").SetObservance(nextMonday),
		BoxingDay.Copy("Boxing Day").SetObservance(nextMonday), // PB!!!! if christmas on saturday
		NewYearsEve, // early-closing
	)
	return c
}

// euronext
func euronext(name string, loc *time.Location, years ...int) *Calendar {
	c := NewCalendar(name, loc, years...)
	c.AddHolidays(
		NewYear,
		GoodFriday,
		EasterMonday,
		WorkersDay,
		ChristmasDay,
		ChristmasEve, // early-closing
	)
	return c
}

// Euronext Amsterdam
func XAMS(years ...int) *Calendar {
	return euronext("Euronext Amsterdam", Amsterdam, years...)
}

// Euronext Brussels
func XBRU(years ...int) *Calendar {
	return euronext("Euronext Brussels", Brussels, years...)
}

// Euronext Lisbon
func XLIS(years ...int) *Calendar {
	return euronext("Euronext Lisbon", Lisbon, years...)
}

// Euronext Paris
func XPAR(years ...int) *Calendar {
	return euronext("Euronext Paris", Paris, years...)
}

// Euronext Milan - Borsa Italiana S.P.A
func XMIL(years ...int) *Calendar {
	c := NewCalendar("Euronext Milan", Milan, years...)
	//TODO: add holidays
	return c
}

// Madrid Stock Exchange
func XMAD(years ...int) *Calendar {
	c := NewCalendar("Madrid Stock Exchange", Madrid, years...)
	//TODO: add holidays
	return c
}

// Frankfurt Stock Exchange
func XFRA(years ...int) *Calendar {
	c := NewCalendar("Deutsche Boerse", Franckfurt, years...)
	//TODO: add holidays
	return c
}

// Deutsche Borse Xetra
func XETR(years ...int) *Calendar {
	c := NewCalendar("Deutsche Borse Xetra", Franckfurt, years...)
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
	//TODO: add holidays
	return c
}

// Shenzhen Stock Exchange
func XSHE(years ...int) *Calendar {
	c := NewCalendar("Shenzhen Stock Exchange", Shenzhen, years...)
	//TODO: add holidays
	return c
}

// Shanghai Stock Exchange
func XSHG(years ...int) *Calendar {
	c := NewCalendar("Shanghai Stock Exchange", Shanghai, years...)
	//TODO: add holidays
	return c
}

// Korea Exchange
func XKRX(years ...int) *Calendar {
	c := NewCalendar("Korea Exchange", Seoul, years...)
	//TODO: add holidays
	return c
}

// Japan Exchange Group
func XJPX(years ...int) *Calendar {
	c := NewCalendar("Japan Exchange Group", Tokyo, years...)
	//TODO: add holidays
	return c
}

// Australian Securities Exchange
func XASX(years ...int) *Calendar {
	c := NewCalendar("Australian Securities Exchange", Sydney, years...)
	//TODO: add holidays
	return c
}
