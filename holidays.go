package calendar

import "time"

var (
	NewYear = &Holiday{
		Name:  "New Year's Day",
		Month: time.January,
		Day:   1,
		calc:  CalcDayOfMonth,
	}
)
