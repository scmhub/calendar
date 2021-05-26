package calendar

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

// Chicago Board Options Exchang
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
	//TODO: add holidays
	return c
}

// Euronext Paris
func EuronextParis(start, end int) *Calendar {
	c := NewCalendar("Euronext Paris", Paris, start, end)
	//TODO: add holidays
	return c
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
