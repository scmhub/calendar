package calendar

import "time"

// America

// New York Stock Exchange
func NYSE(years ...int) *Calendar {
	c := NewCalendar("New York Stock Exchange", NewYork, years...)
	//TODO: add holidays
	return c
}

// NASDAQ
func NASDAQ(years ...int) *Calendar {
	c := NewCalendar("NASDAQ", NewYork, years...)
	//TODO: add holidays
	return c
}

// Chicago Board Options Exchange
func CBOE(years ...int) *Calendar {
	c := NewCalendar("Chicago Board Options Exchange", Chicago, years...)
	//TODO: add holidays
	return c
}

// Cboe Futures Exchange
func CFE(years ...int) *Calendar {
	c := NewCalendar("Cboe Futures Exchange", Chicago, years...)
	//TODO: add holidays
	return c
}

// Europe

// London Stock Exchange
func LSE(years ...int) *Calendar {
	c := NewCalendar("London Stock Exchange", London, years...)
	c.addHoliday(NewYear.Copy("New Year's Day").SetObservance(nextMonday))
	c.addHoliday(GoodFriday.Copy("Good Friday"))
	c.addHoliday(EasterMonday.Copy("Easter Monday"))
	c.addHoliday(EarlyMay.Copy("Early May"))
	c.addHoliday(LateMay.Copy("Late May"))
	c.addHoliday(SummerHoliday.Copy("Summer Bank Holiday"))
	c.addHoliday(ChristmasEve.Copy("Christmas Eve")) // early-closing
	c.addHoliday(ChristmasDay.Copy("Christmas Day").SetObservance(nextMonday))
	c.addHoliday(BoxingDay.Copy("Boxing Day").SetObservance(nextMonday)) // PB!!!! if christmas on saturday
	c.addHoliday(NewYearsEve.Copy("New Year's Eve"))                     // early-closing
	return c
}

// euronext
func euronext(name string, loc *time.Location, years ...int) *Calendar {
	c := NewCalendar(name, loc, years...)
	c.addHoliday(NewYear.Copy("New Year's Day"))
	c.addHoliday(GoodFriday.Copy("Good Friday"))
	c.addHoliday(EasterMonday.Copy("Easter Monday"))
	c.addHoliday(WorkersDay.Copy("Worker's Day"))
	c.addHoliday(ChristmasDay.Copy("Christmas Day"))
	c.addHoliday(ChristmasEve.Copy("Christmas Eve")) // early-closing
	return c
}

// Euronext Amsterdam
func EuronextAmsterdam(years ...int) *Calendar {
	return euronext("Euronext Amsterdam", Amsterdam, years...)
}

// Euronext Brussels
func EuronextBrussels(years ...int) *Calendar {
	return euronext("Euronext Brussels", Brussels, years...)
}

// Euronext Lisbon
func EuronextLisbon(years ...int) *Calendar {
	return euronext("Euronext Lisbon", Lisbon, years...)
}

// Euronext Paris
func EuronextParis(years ...int) *Calendar {
	return euronext("Euronext Paris", Paris, years...)
}

// Deutsche Borse
func DB(years ...int) *Calendar {
	c := NewCalendar("Deutsche Borse", Franckfurt, years...)
	//TODO: add holidays
	return c
}

// SIX Group (SWX Swiss Exchange)
func SIX(years ...int) *Calendar {
	c := NewCalendar("SIX Group", Zurich, years...)
	//TODO: add holidays
	return c
}

// Asia

// Bombay Stock Exchange
func BSE(years ...int) *Calendar {
	c := NewCalendar("Bombay Stock Exchange", Bombay, years...)
	//TODO: add holidays
	return c
}

// Singapore Exchange (SGX)
func SGX(years ...int) *Calendar {
	c := NewCalendar("Singapore Exchange", Singapore, years...)
	//TODO: add holidays
	return c
}

// Stock Exchange of Hong Kong
func SEHK(years ...int) *Calendar {
	c := NewCalendar("Stock Exchange of Hong Kong", HongKong, years...)
	//TODO: add holidays
	return c
}

// Shenzhen Stock Exchange
func ShenzhenSE(years ...int) *Calendar {
	c := NewCalendar("Shenzhen Stock Exchange", Shenzhen, years...)
	//TODO: add holidays
	return c
}

// Shanghai Stock Exchange
func ShanghaiSE(years ...int) *Calendar {
	c := NewCalendar("Shanghai Stock Exchange", Shanghai, years...)
	//TODO: add holidays
	return c
}

// Korea Exchange
func KRX(years ...int) *Calendar {
	c := NewCalendar("Korea Exchange", Seoul, years...)
	//TODO: add holidays
	return c
}

// Japan Exchange Group
func JEG(years ...int) *Calendar {
	c := NewCalendar("Japan Exchange Group", Tokyo, years...)
	//TODO: add holidays
	return c
}

// Australian Securities Exchange
func ASE(years ...int) *Calendar {
	c := NewCalendar("Australian Securities Exchange", Sydney, years...)
	//TODO: add holidays
	return c
}
