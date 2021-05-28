# ![SCM logo](logo.png) Calendar

Exchange Calendar

```go
// Get an exchange calendar
nyse := XNYS(2010, 2025)

// Create an Holiday
MemorialDay = &Holiday{
    Name:       "Memorial Day",
    Month:      time.May,
    Weekday:    time.Monday,
    NthWeekday: -1,
    calc:       CalcNthWeekday,
}
IndependenceDay = &Holiday{
    Name:       "Independence Day",
    Month:      time.July,
    Day:        4,
    observance: nearestWorkday,
    calc:       CalcDayOfMonth,
}
// Copy an Holiday and set observance
NewYear.Copy("New Year's Day").SetObservance(sundayToMonday)

// Create a Calendar
c := NewCalendar("New York Stock Exchange", NewYork, years...)
	c.SetEarlySession(&Session{7 * time.Hour, 9*time.Hour + 30*time.Minute})
	c.SetCoreSessions(&Session{9*time.Hour + 30*time.Minute, 16 * time.Hour})
	c.SetLateSession(&Session{16 * time.Hour, 20 * time.Hour})
	c.AddHolidays(
		NewYear.Copy("New Year's Day").SetObservance(sundayToMonday),
		MLKDay,
		PresidentsDay,
		GoodFriday,
		MemorialDay,
		IndependenceDay.Copy("Independence Day").SetObservance(nearestWorkday), // !! half monday if tuesday
		LaborDay,
		ThanksgivingDay,
		BlackFriday, // !! early-closing 1 pm
		ChristmasDay.Copy("Christmas Day").SetObservance(nearestWorkday),
	)
```