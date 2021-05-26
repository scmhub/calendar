package calendar

import "time"

// America

// New York Stock Exchange
func NYSE(start, end int) *Calendar {
	c := NewCalendar("New York Stock Exchange", NewYork, start, end)
	//TODO: add holidays
	return c
}

// NASDAQ
func NASDAQ(start, end int) *Calendar {
	c := NewCalendar("NASDAQ", NewYork, start, end)
	//TODO: add holidays
	return c
}

// Chicago Board Options Exchange
func CBOE(start, end int) *Calendar {
	c := NewCalendar("Chicago Board Options Exchange", Chicago, start, end)
	//TODO: add holidays
	return c
}

// Cboe Futures Exchange
func CFE(start, end int) *Calendar {
	c := NewCalendar("Cboe Futures Exchange", Chicago, start, end)
	//TODO: add holidays
	return c
}

// Europe

// London Stock Exchange
func LSE(start, end int) *Calendar {
	c := NewCalendar("London Stock Exchange", London, start, end)
	c.addHoliday(NewYear.Copy("New Year's Day").SetObservance(nextMonday))
	c.addHoliday(GoodFriday.Copy("Good Friday"))
	c.addHoliday(EasterMonday.Copy("Easter Monday"))
	c.addHoliday(EarlyMay.Copy("Early May"))
	c.addHoliday(LateMay.Copy("Late May"))
	c.addHoliday(SummerHoliday.Copy("Summer Bank Holiday"))
	c.addHoliday(ChristmasDay.Copy("Christmas Day").SetObservance(nextMonday))
	c.addHoliday(BoxingDay.Copy("Boxing Day").SetObservance(nextMonday)) // PB!!!! if christmas on saturday
	c.addHoliday(ChristmasEve.Copy("Christmas Eve"))                     // early-closing
	return c
}

// euronext
func euronext(name string, loc *time.Location, start, end int) *Calendar {
	c := NewCalendar(name, loc, start, end)
	c.addHoliday(NewYear.Copy("New Year's Day"))
	c.addHoliday(GoodFriday.Copy("Good Friday"))
	c.addHoliday(EasterMonday.Copy("Easter Monday"))
	c.addHoliday(WorkersDay.Copy("Worker's Day"))
	c.addHoliday(ChristmasDay.Copy("Christmas Day"))
	c.addHoliday(ChristmasEve.Copy("Christmas Eve")) // early-closing
	return c
}

// Euronext Amsterdam
func EuronextAmsterdam(start, end int) *Calendar {
	return euronext("Euronext Amsterdam", Amsterdam, start, end)
}

// Euronext Brussels
func EuronextBrussels(start, end int) *Calendar {
	return euronext("Euronext Brussels", Brussels, start, end)
}

// Euronext Lisbon
func EuronextLisbon(start, end int) *Calendar {
	return euronext("Euronext Lisbon", Lisbon, start, end)
}

// Euronext Paris
func EuronextParis(start, end int) *Calendar {
	return euronext("Euronext Paris", Paris, start, end)
}

// Asia

// Hong Kong Exchanges
func HKE(start, end int) *Calendar {
	c := NewCalendar("Hong Kong Exchanges", HongKong, start, end)
	//TODO: add holidays
	return c
}

// Shenzhen Stock Exchange
func ShenzhenSE(start, end int) *Calendar {
	c := NewCalendar("Shenzhen Stock Exchange", Shenzhen, start, end)
	//TODO: add holidays
	return c
}

// Shanghai Stock Exchange
func ShanghaiSE(start, end int) *Calendar {
	c := NewCalendar("Shanghai Stock Exchange", Shanghai, start, end)
	//TODO: add holidays
	return c
}

// Korea Exchange
func KRX(start, end int) *Calendar {
	c := NewCalendar("Korea Exchange", Seoul, start, end)
	//TODO: add holidays
	return c
}

// Japan Exchange Group
func JEG(start, end int) *Calendar {
	c := NewCalendar("Japan Exchange Group", Tokyo, start, end)
	//TODO: add holidays
	return c
}

// Australian Securities Exchange
func ASE(start, end int) *Calendar {
	c := NewCalendar("Australian Securities Exchange", Sydney, start, end)
	//TODO: add holidays
	return c
}
