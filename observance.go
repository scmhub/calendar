package calendar

import "time"

type observance func(time.Time) time.Time

// Saturday or Sunday => next Monday
func nextMonday(t time.Time) time.Time {
	wd := t.Weekday()
	if wd == time.Saturday {
		return t.AddDate(0, 0, 2)
	}
	if wd == time.Sunday {
		return t.AddDate(0, 0, 1)
	}
	return t
}

// Saturday or Sunday => previous Friday
func previousFriday(t time.Time) time.Time {
	wd := t.Weekday()
	if wd == time.Saturday {
		return t.AddDate(0, 0, -1)
	}
	if wd == time.Sunday {
		return t.AddDate(0, 0, -2)
	}
	return t
}

// Saturday => previous Friday, Sunday => next Monday
func nearestWorkday(t time.Time) time.Time {
	wd := t.Weekday()
	if wd == time.Saturday {
		return t.AddDate(0, 0, -1)
	}
	if wd == time.Sunday {
		return t.AddDate(0, 0, +1)
	}
	return t
}

// Sunday => next Monday
func sundayToMonday(t time.Time) time.Time {
	wd := t.Weekday()
	if wd == time.Sunday {
		return t.AddDate(0, 0, 1)
	}
	return t
}

// Closure of observance.
// time.Time{} if not a passed Weekday.
func onlyOnWeekdays(wd ...time.Weekday) observance {
	return func(t time.Time) time.Time {
		for _, weekday := range wd {
			if t.Weekday() == weekday {
				return t
			}
		}
		return time.Time{}
	}
}
