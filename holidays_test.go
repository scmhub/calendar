package calendar

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Common Holidays

func TestNewYear(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 1, 1, 0, 0, 0, 0, Paris), NewYear.Calc(2010, Paris)) // friday
	assert.Equal(time.Time{}, NewYear.Calc(2011, Paris))                              // saturday
	assert.Equal(time.Time{}, NewYear.Calc(2012, Paris))                              // sunday
	assert.Equal(time.Date(2013, 1, 1, 0, 0, 0, 0, Paris), NewYear.Calc(2013, Paris)) // monday
	assert.Equal(time.Date(2014, 1, 1, 0, 0, 0, 0, Paris), NewYear.Calc(2014, Paris)) // tuesday
	assert.Equal(time.Date(2015, 1, 1, 0, 0, 0, 0, Paris), NewYear.Calc(2015, Paris)) // thursday
	assert.Equal(time.Date(2016, 1, 1, 0, 0, 0, 0, Paris), NewYear.Calc(2016, Paris)) // friday
	assert.Equal(time.Time{}, NewYear.Calc(2017, Paris))                              // sunday
	assert.Equal(time.Date(2018, 1, 1, 0, 0, 0, 0, Paris), NewYear.Calc(2018, Paris)) // monday
	assert.Equal(time.Date(2019, 1, 1, 0, 0, 0, 0, Paris), NewYear.Calc(2019, Paris)) // tuesday
	assert.Equal(time.Date(2020, 1, 1, 0, 0, 0, 0, Paris), NewYear.Calc(2020, Paris)) // wednesday
	assert.Equal(time.Date(2021, 1, 1, 0, 0, 0, 0, Paris), NewYear.Calc(2021, Paris)) // friday
	assert.Equal(time.Time{}, NewYear.Calc(2022, Paris))                              // saturday
	assert.Equal(time.Time{}, NewYear.Calc(2023, Paris))                              // sunday
	assert.Equal(time.Date(2024, 1, 1, 0, 0, 0, 0, Paris), NewYear.Calc(2024, Paris)) // monday
	assert.Equal(time.Date(2025, 1, 1, 0, 0, 0, 0, Paris), NewYear.Calc(2025, Paris)) // wednesday
}

func TestEpiphany(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2020, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2020, Paris))
	assert.Equal(time.Date(2021, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2023, Paris))
	assert.Equal(time.Time{}, Epiphany.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2025, Paris))
}
func TestCarnival(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2020, 2, 25, 0, 0, 0, 0, Paris), Carnival.Calc(2020, Paris))
	assert.Equal(time.Date(2021, 2, 16, 0, 0, 0, 0, Paris), Carnival.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 3, 1, 0, 0, 0, 0, Paris), Carnival.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 2, 21, 0, 0, 0, 0, Paris), Carnival.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 2, 13, 0, 0, 0, 0, Paris), Carnival.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 3, 4, 0, 0, 0, 0, Paris), Carnival.Calc(2025, Paris))
}
func TestMaundyThursday(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 4, 1, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2010, Paris))
	assert.Equal(time.Date(2011, 4, 21, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2011, Paris))
	assert.Equal(time.Date(2012, 4, 5, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2012, Paris))
	assert.Equal(time.Date(2013, 3, 28, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2013, Paris))
	assert.Equal(time.Date(2014, 4, 17, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2014, Paris))
	assert.Equal(time.Date(2015, 4, 2, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2015, Paris))
	assert.Equal(time.Date(2016, 3, 24, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2016, Paris))
	assert.Equal(time.Date(2017, 4, 13, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2017, Paris))
	assert.Equal(time.Date(2018, 3, 29, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2018, Paris))
	assert.Equal(time.Date(2019, 4, 18, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2019, Paris))
	assert.Equal(time.Date(2020, 4, 9, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2020, Paris))
	assert.Equal(time.Date(2021, 4, 1, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 4, 14, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 4, 6, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 3, 28, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 4, 17, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2025, Paris))
}

func TestGoodFriday(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 4, 2, 0, 0, 0, 0, Paris), GoodFriday.Calc(2010, Paris))
	assert.Equal(time.Date(2011, 4, 22, 0, 0, 0, 0, Paris), GoodFriday.Calc(2011, Paris))
	assert.Equal(time.Date(2012, 4, 6, 0, 0, 0, 0, Paris), GoodFriday.Calc(2012, Paris))
	assert.Equal(time.Date(2013, 3, 29, 0, 0, 0, 0, Paris), GoodFriday.Calc(2013, Paris))
	assert.Equal(time.Date(2014, 4, 18, 0, 0, 0, 0, Paris), GoodFriday.Calc(2014, Paris))
	assert.Equal(time.Date(2015, 4, 3, 0, 0, 0, 0, Paris), GoodFriday.Calc(2015, Paris))
	assert.Equal(time.Date(2016, 3, 25, 0, 0, 0, 0, Paris), GoodFriday.Calc(2016, Paris))
	assert.Equal(time.Date(2017, 4, 14, 0, 0, 0, 0, Paris), GoodFriday.Calc(2017, Paris))
	assert.Equal(time.Date(2018, 3, 30, 0, 0, 0, 0, Paris), GoodFriday.Calc(2018, Paris))
	assert.Equal(time.Date(2019, 4, 19, 0, 0, 0, 0, Paris), GoodFriday.Calc(2019, Paris))
	assert.Equal(time.Date(2020, 4, 10, 0, 0, 0, 0, Paris), GoodFriday.Calc(2020, Paris))
	assert.Equal(time.Date(2021, 4, 2, 0, 0, 0, 0, Paris), GoodFriday.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 4, 15, 0, 0, 0, 0, Paris), GoodFriday.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 4, 7, 0, 0, 0, 0, Paris), GoodFriday.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 3, 29, 0, 0, 0, 0, Paris), GoodFriday.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 4, 18, 0, 0, 0, 0, Paris), GoodFriday.Calc(2025, Paris))
}
func TestNorthwardEquinox(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Time{}, NorthwardEquinox.Calc(2010, time.UTC))
	assert.Equal(time.Time{}, NorthwardEquinox.Calc(2011, time.UTC))
	assert.Equal(time.Date(2012, 3, 20, 0, 0, 0, 0, time.UTC), NorthwardEquinox.Calc(2012, time.UTC))
	assert.Equal(time.Date(2013, 3, 20, 0, 0, 0, 0, time.UTC), NorthwardEquinox.Calc(2013, time.UTC))
	assert.Equal(time.Date(2014, 3, 20, 0, 0, 0, 0, time.UTC), NorthwardEquinox.Calc(2014, time.UTC))
	assert.Equal(time.Date(2015, 3, 20, 0, 0, 0, 0, time.UTC), NorthwardEquinox.Calc(2015, time.UTC))
	assert.Equal(time.Time{}, NorthwardEquinox.Calc(2016, time.UTC))
	assert.Equal(time.Date(2017, 3, 20, 0, 0, 0, 0, time.UTC), NorthwardEquinox.Calc(2017, time.UTC))
	assert.Equal(time.Date(2018, 3, 20, 0, 0, 0, 0, time.UTC), NorthwardEquinox.Calc(2018, time.UTC))
	assert.Equal(time.Date(2019, 3, 20, 0, 0, 0, 0, time.UTC), NorthwardEquinox.Calc(2019, time.UTC))
	assert.Equal(time.Date(2020, 3, 20, 0, 0, 0, 0, time.UTC), NorthwardEquinox.Calc(2020, time.UTC))
	assert.Equal(time.Time{}, NorthwardEquinox.Calc(2021, time.UTC))
	assert.Equal(time.Time{}, NorthwardEquinox.Calc(2022, time.UTC))
	assert.Equal(time.Date(2023, 3, 20, 0, 0, 0, 0, time.UTC), NorthwardEquinox.Calc(2023, time.UTC))
	assert.Equal(time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC), NorthwardEquinox.Calc(2024, time.UTC))
	assert.Equal(time.Date(2025, 3, 20, 0, 0, 0, 0, time.UTC), NorthwardEquinox.Calc(2025, time.UTC))
}
func TestEasterMonday(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 4, 5, 0, 0, 0, 0, Paris), EasterMonday.Calc(2010, Paris))
	assert.Equal(time.Date(2011, 4, 25, 0, 0, 0, 0, Paris), EasterMonday.Calc(2011, Paris))
	assert.Equal(time.Date(2012, 4, 9, 0, 0, 0, 0, Paris), EasterMonday.Calc(2012, Paris))
	assert.Equal(time.Date(2013, 4, 1, 0, 0, 0, 0, Paris), EasterMonday.Calc(2013, Paris))
	assert.Equal(time.Date(2014, 4, 21, 0, 0, 0, 0, Paris), EasterMonday.Calc(2014, Paris))
	assert.Equal(time.Date(2015, 4, 6, 0, 0, 0, 0, Paris), EasterMonday.Calc(2015, Paris))
	assert.Equal(time.Date(2016, 3, 28, 0, 0, 0, 0, Paris), EasterMonday.Calc(2016, Paris))
	assert.Equal(time.Date(2017, 4, 17, 0, 0, 0, 0, Paris), EasterMonday.Calc(2017, Paris))
	assert.Equal(time.Date(2018, 4, 2, 0, 0, 0, 0, Paris), EasterMonday.Calc(2018, Paris))
	assert.Equal(time.Date(2019, 4, 22, 0, 0, 0, 0, Paris), EasterMonday.Calc(2019, Paris))
	assert.Equal(time.Date(2020, 4, 13, 0, 0, 0, 0, Paris), EasterMonday.Calc(2020, Paris))
	assert.Equal(time.Date(2021, 4, 5, 0, 0, 0, 0, Paris), EasterMonday.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 4, 18, 0, 0, 0, 0, Paris), EasterMonday.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 4, 10, 0, 0, 0, 0, Paris), EasterMonday.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 4, 1, 0, 0, 0, 0, Paris), EasterMonday.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 4, 21, 0, 0, 0, 0, Paris), EasterMonday.Calc(2025, Paris))
}

func TestWorkerDays(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2020, 5, 1, 0, 0, 0, 0, Paris), WorkersDay.Calc(2020, Paris))
	assert.Equal(time.Time{}, WorkersDay.Calc(2021, Paris))
	assert.Equal(time.Time{}, WorkersDay.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 5, 1, 0, 0, 0, 0, Paris), WorkersDay.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 5, 1, 0, 0, 0, 0, Paris), WorkersDay.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 5, 1, 0, 0, 0, 0, Paris), WorkersDay.Calc(2025, Paris))
}

func TestAscensionDay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 5, 13, 0, 0, 0, 0, Paris), AscensionDay.Calc(2010, Paris))
	assert.Equal(time.Date(2011, 6, 2, 0, 0, 0, 0, Paris), AscensionDay.Calc(2011, Paris))
	assert.Equal(time.Date(2012, 5, 17, 0, 0, 0, 0, Paris), AscensionDay.Calc(2012, Paris))
	assert.Equal(time.Date(2013, 5, 9, 0, 0, 0, 0, Paris), AscensionDay.Calc(2013, Paris))
	assert.Equal(time.Date(2014, 5, 29, 0, 0, 0, 0, Paris), AscensionDay.Calc(2014, Paris))
	assert.Equal(time.Date(2015, 5, 14, 0, 0, 0, 0, Paris), AscensionDay.Calc(2015, Paris))
	assert.Equal(time.Date(2016, 5, 5, 0, 0, 0, 0, Paris), AscensionDay.Calc(2016, Paris))
	assert.Equal(time.Date(2017, 5, 25, 0, 0, 0, 0, Paris), AscensionDay.Calc(2017, Paris))
	assert.Equal(time.Date(2018, 5, 10, 0, 0, 0, 0, Paris), AscensionDay.Calc(2018, Paris))
	assert.Equal(time.Date(2019, 5, 30, 0, 0, 0, 0, Paris), AscensionDay.Calc(2019, Paris))
	assert.Equal(time.Date(2020, 5, 21, 0, 0, 0, 0, Paris), AscensionDay.Calc(2020, Paris))
	assert.Equal(time.Date(2021, 5, 13, 0, 0, 0, 0, Paris), AscensionDay.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 5, 26, 0, 0, 0, 0, Paris), AscensionDay.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 5, 18, 0, 0, 0, 0, Paris), AscensionDay.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 5, 9, 0, 0, 0, 0, Paris), AscensionDay.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 5, 29, 0, 0, 0, 0, Paris), AscensionDay.Calc(2025, Paris))
}

func TestPentecostMonday(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 5, 24, 0, 0, 0, 0, Paris), PentecostMonday.Calc(2010, Paris))
	assert.Equal(time.Date(2011, 6, 13, 0, 0, 0, 0, Paris), PentecostMonday.Calc(2011, Paris))
	assert.Equal(time.Date(2012, 5, 28, 0, 0, 0, 0, Paris), PentecostMonday.Calc(2012, Paris))
	assert.Equal(time.Date(2013, 5, 20, 0, 0, 0, 0, Paris), PentecostMonday.Calc(2013, Paris))
	assert.Equal(time.Date(2014, 6, 9, 0, 0, 0, 0, Paris), PentecostMonday.Calc(2014, Paris))
	assert.Equal(time.Date(2015, 5, 25, 0, 0, 0, 0, Paris), PentecostMonday.Calc(2015, Paris))
	assert.Equal(time.Date(2016, 5, 16, 0, 0, 0, 0, Paris), PentecostMonday.Calc(2016, Paris))
	assert.Equal(time.Date(2017, 6, 5, 0, 0, 0, 0, Paris), PentecostMonday.Calc(2017, Paris))
	assert.Equal(time.Date(2018, 5, 21, 0, 0, 0, 0, Paris), PentecostMonday.Calc(2018, Paris))
	assert.Equal(time.Date(2019, 6, 10, 0, 0, 0, 0, Paris), PentecostMonday.Calc(2019, Paris))
	assert.Equal(time.Date(2020, 6, 1, 0, 0, 0, 0, Paris), PentecostMonday.Calc(2020, Paris))
	assert.Equal(time.Date(2021, 5, 24, 0, 0, 0, 0, Paris), PentecostMonday.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 6, 6, 0, 0, 0, 0, Paris), PentecostMonday.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 5, 29, 0, 0, 0, 0, Paris), PentecostMonday.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 5, 20, 0, 0, 0, 0, Paris), PentecostMonday.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 6, 9, 0, 0, 0, 0, Paris), PentecostMonday.Calc(2025, Paris))
}

func TestCorpsChristi(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 6, 3, 0, 0, 0, 0, Paris), CorpusChristi.Calc(2010, Paris))
	assert.Equal(time.Date(2011, 6, 23, 0, 0, 0, 0, Paris), CorpusChristi.Calc(2011, Paris))
	assert.Equal(time.Date(2012, 6, 7, 0, 0, 0, 0, Paris), CorpusChristi.Calc(2012, Paris))
	assert.Equal(time.Date(2013, 5, 30, 0, 0, 0, 0, Paris), CorpusChristi.Calc(2013, Paris))
	assert.Equal(time.Date(2014, 6, 19, 0, 0, 0, 0, Paris), CorpusChristi.Calc(2014, Paris))
	assert.Equal(time.Date(2015, 6, 4, 0, 0, 0, 0, Paris), CorpusChristi.Calc(2015, Paris))
	assert.Equal(time.Date(2016, 5, 26, 0, 0, 0, 0, Paris), CorpusChristi.Calc(2016, Paris))
	assert.Equal(time.Date(2017, 6, 15, 0, 0, 0, 0, Paris), CorpusChristi.Calc(2017, Paris))
	assert.Equal(time.Date(2018, 5, 31, 0, 0, 0, 0, Paris), CorpusChristi.Calc(2018, Paris))
	assert.Equal(time.Date(2019, 6, 20, 0, 0, 0, 0, Paris), CorpusChristi.Calc(2019, Paris))
	assert.Equal(time.Date(2020, 6, 11, 0, 0, 0, 0, Paris), CorpusChristi.Calc(2020, Paris))
	assert.Equal(time.Date(2021, 6, 3, 0, 0, 0, 0, Paris), CorpusChristi.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 6, 16, 0, 0, 0, 0, Paris), CorpusChristi.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 6, 8, 0, 0, 0, 0, Paris), CorpusChristi.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 5, 30, 0, 0, 0, 0, Paris), CorpusChristi.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 6, 19, 0, 0, 0, 0, Paris), CorpusChristi.Calc(2025, Paris))
}
func TestNorthernSolstice(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 6, 21, 0, 0, 0, 0, time.UTC), NorthernSolstice.Calc(2010, time.UTC))
	assert.Equal(time.Date(2011, 6, 21, 0, 0, 0, 0, time.UTC), NorthernSolstice.Calc(2011, time.UTC))
	assert.Equal(time.Date(2012, 6, 20, 0, 0, 0, 0, time.UTC), NorthernSolstice.Calc(2012, time.UTC))
	assert.Equal(time.Date(2013, 6, 21, 0, 0, 0, 0, time.UTC), NorthernSolstice.Calc(2013, time.UTC))
	assert.Equal(time.Time{}, NorthernSolstice.Calc(2014, time.UTC))
	assert.Equal(time.Time{}, NorthernSolstice.Calc(2015, time.UTC))
	assert.Equal(time.Date(2016, 6, 20, 0, 0, 0, 0, time.UTC), NorthernSolstice.Calc(2016, time.UTC))
	assert.Equal(time.Date(2017, 6, 21, 0, 0, 0, 0, time.UTC), NorthernSolstice.Calc(2017, time.UTC))
	assert.Equal(time.Date(2018, 6, 21, 0, 0, 0, 0, time.UTC), NorthernSolstice.Calc(2018, time.UTC))
	assert.Equal(time.Date(2019, 6, 21, 0, 0, 0, 0, time.UTC), NorthernSolstice.Calc(2019, time.UTC))
	assert.Equal(time.Time{}, NorthernSolstice.Calc(2020, time.UTC))
	assert.Equal(time.Date(2021, 6, 21, 0, 0, 0, 0, time.UTC), NorthernSolstice.Calc(2021, time.UTC))
	assert.Equal(time.Date(2022, 6, 21, 0, 0, 0, 0, time.UTC), NorthernSolstice.Calc(2022, time.UTC))
	assert.Equal(time.Date(2023, 6, 21, 0, 0, 0, 0, time.UTC), NorthernSolstice.Calc(2023, time.UTC))
	assert.Equal(time.Date(2024, 6, 20, 0, 0, 0, 0, time.UTC), NorthernSolstice.Calc(2024, time.UTC))
	assert.Equal(time.Time{}, NorthernSolstice.Calc(2025, time.UTC))
}
func TestAssumptionOfMary(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Time{}, AssumptionOfMary.Calc(2010, Paris))
	assert.Equal(time.Date(2011, 8, 15, 0, 0, 0, 0, Paris), AssumptionOfMary.Calc(2011, Paris))
	assert.Equal(time.Date(2012, 8, 15, 0, 0, 0, 0, Paris), AssumptionOfMary.Calc(2012, Paris))
	assert.Equal(time.Date(2013, 8, 15, 0, 0, 0, 0, Paris), AssumptionOfMary.Calc(2013, Paris))
	assert.Equal(time.Date(2014, 8, 15, 0, 0, 0, 0, Paris), AssumptionOfMary.Calc(2014, Paris))
	assert.Equal(time.Time{}, AssumptionOfMary.Calc(2015, Paris))
	assert.Equal(time.Date(2016, 8, 15, 0, 0, 0, 0, Paris), AssumptionOfMary.Calc(2016, Paris))
	assert.Equal(time.Date(2017, 8, 15, 0, 0, 0, 0, Paris), AssumptionOfMary.Calc(2017, Paris))
	assert.Equal(time.Date(2018, 8, 15, 0, 0, 0, 0, Paris), AssumptionOfMary.Calc(2018, Paris))
	assert.Equal(time.Date(2019, 8, 15, 0, 0, 0, 0, Paris), AssumptionOfMary.Calc(2019, Paris))
	assert.Equal(time.Time{}, AssumptionOfMary.Calc(2020, Paris))
	assert.Equal(time.Time{}, AssumptionOfMary.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 8, 15, 0, 0, 0, 0, Paris), AssumptionOfMary.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 8, 15, 0, 0, 0, 0, Paris), AssumptionOfMary.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 8, 15, 0, 0, 0, 0, Paris), AssumptionOfMary.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 8, 15, 0, 0, 0, 0, Paris), AssumptionOfMary.Calc(2025, Paris))
}
func TestSouthwardEquinox(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 9, 23, 0, 0, 0, 0, time.UTC), SouthwardEquinox.Calc(2010, time.UTC))
	assert.Equal(time.Date(2011, 9, 23, 0, 0, 0, 0, time.UTC), SouthwardEquinox.Calc(2011, time.UTC))
	assert.Equal(time.Time{}, SouthwardEquinox.Calc(2012, time.UTC))
	assert.Equal(time.Time{}, SouthwardEquinox.Calc(2013, time.UTC))
	assert.Equal(time.Date(2014, 9, 23, 0, 0, 0, 0, time.UTC), SouthwardEquinox.Calc(2014, time.UTC))
	assert.Equal(time.Date(2015, 9, 23, 0, 0, 0, 0, time.UTC), SouthwardEquinox.Calc(2015, time.UTC))
	assert.Equal(time.Date(2016, 9, 22, 0, 0, 0, 0, time.UTC), SouthwardEquinox.Calc(2016, time.UTC))
	assert.Equal(time.Date(2017, 9, 22, 0, 0, 0, 0, time.UTC), SouthwardEquinox.Calc(2017, time.UTC))
	assert.Equal(time.Time{}, SouthwardEquinox.Calc(2018, time.UTC))
	assert.Equal(time.Date(2019, 9, 23, 0, 0, 0, 0, time.UTC), SouthwardEquinox.Calc(2019, time.UTC))
	assert.Equal(time.Date(2020, 9, 22, 0, 0, 0, 0, time.UTC), SouthwardEquinox.Calc(2020, time.UTC))
	assert.Equal(time.Date(2021, 9, 22, 0, 0, 0, 0, time.UTC), SouthwardEquinox.Calc(2021, time.UTC))
	assert.Equal(time.Date(2022, 9, 23, 0, 0, 0, 0, time.UTC), SouthwardEquinox.Calc(2022, time.UTC))
	assert.Equal(time.Time{}, SouthwardEquinox.Calc(2023, time.UTC))
	assert.Equal(time.Time{}, SouthwardEquinox.Calc(2024, time.UTC))
	assert.Equal(time.Date(2025, 9, 22, 0, 0, 0, 0, time.UTC), SouthwardEquinox.Calc(2025, time.UTC))
}
func TestReformationDay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Time{}, ReformationDay.Calc(2020, Paris))
	assert.Equal(time.Time{}, ReformationDay.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 10, 31, 0, 0, 0, 0, Paris), ReformationDay.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 10, 31, 0, 0, 0, 0, Paris), ReformationDay.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 10, 31, 0, 0, 0, 0, Paris), ReformationDay.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 10, 31, 0, 0, 0, 0, Paris), ReformationDay.Calc(2025, Paris))
}

func TestAllSaintsDay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Time{}, AllSaintsDay.Calc(2020, Paris))
	assert.Equal(time.Date(2021, 11, 1, 0, 0, 0, 0, Paris), AllSaintsDay.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 11, 1, 0, 0, 0, 0, Paris), AllSaintsDay.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 11, 1, 0, 0, 0, 0, Paris), AllSaintsDay.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 11, 1, 0, 0, 0, 0, Paris), AllSaintsDay.Calc(2024, Paris))
	assert.Equal(time.Time{}, AllSaintsDay.Calc(2025, Paris))
}

func TestArmisticeDay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2020, 11, 11, 0, 0, 0, 0, Paris), ArmisticeDay.Calc(2020, Paris))
	assert.Equal(time.Date(2021, 11, 11, 0, 0, 0, 0, Paris), ArmisticeDay.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 11, 11, 0, 0, 0, 0, Paris), ArmisticeDay.Calc(2022, Paris))
	assert.Equal(time.Time{}, ArmisticeDay.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 11, 11, 0, 0, 0, 0, Paris), ArmisticeDay.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 11, 11, 0, 0, 0, 0, Paris), ArmisticeDay.Calc(2025, Paris))
}

func TestImmaculateConception(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2020, 12, 8, 0, 0, 0, 0, Paris), ImmaculateConception.Calc(2020, Paris))
	assert.Equal(time.Date(2021, 12, 8, 0, 0, 0, 0, Paris), ImmaculateConception.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 12, 8, 0, 0, 0, 0, Paris), ImmaculateConception.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 12, 8, 0, 0, 0, 0, Paris), ImmaculateConception.Calc(2023, Paris))
	assert.Equal(time.Time{}, ImmaculateConception.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 12, 8, 0, 0, 0, 0, Paris), ImmaculateConception.Calc(2025, Paris))
}

func TestSouthernSolstice(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 12, 21, 0, 0, 0, 0, time.UTC), SouthernSolstice.Calc(2010, time.UTC))
	assert.Equal(time.Date(2011, 12, 22, 0, 0, 0, 0, time.UTC), SouthernSolstice.Calc(2011, time.UTC))
	assert.Equal(time.Date(2012, 12, 21, 0, 0, 0, 0, time.UTC), SouthernSolstice.Calc(2012, time.UTC))
	assert.Equal(time.Time{}, SouthernSolstice.Calc(2013, time.UTC))
	assert.Equal(time.Time{}, SouthernSolstice.Calc(2014, time.UTC))
	assert.Equal(time.Date(2015, 12, 22, 0, 0, 0, 0, time.UTC), SouthernSolstice.Calc(2015, time.UTC))
	assert.Equal(time.Date(2016, 12, 21, 0, 0, 0, 0, time.UTC), SouthernSolstice.Calc(2016, time.UTC))
	assert.Equal(time.Date(2017, 12, 21, 0, 0, 0, 0, time.UTC), SouthernSolstice.Calc(2017, time.UTC))
	assert.Equal(time.Date(2018, 12, 21, 0, 0, 0, 0, time.UTC), SouthernSolstice.Calc(2018, time.UTC))
	assert.Equal(time.Time{}, SouthernSolstice.Calc(2019, time.UTC))
	assert.Equal(time.Date(2020, 12, 21, 0, 0, 0, 0, time.UTC), SouthernSolstice.Calc(2020, time.UTC))
	assert.Equal(time.Date(2021, 12, 21, 0, 0, 0, 0, time.UTC), SouthernSolstice.Calc(2021, time.UTC))
	assert.Equal(time.Date(2022, 12, 21, 0, 0, 0, 0, time.UTC), SouthernSolstice.Calc(2022, time.UTC))
	assert.Equal(time.Date(2023, 12, 22, 0, 0, 0, 0, time.UTC), SouthernSolstice.Calc(2023, time.UTC))
	assert.Equal(time.Time{}, SouthernSolstice.Calc(2024, time.UTC))
	assert.Equal(time.Time{}, SouthernSolstice.Calc(2025, time.UTC))
}

func TestChristmasEve(t *testing.T) {
	assert := assert.New(t)
	fmt.Println(MaundyThursday.Calc(2012, Paris))
	assert.Equal(time.Date(2020, 12, 24, 0, 0, 0, 0, Paris), ChristmasEve.Calc(2020, Paris))
	assert.Equal(time.Date(2021, 12, 24, 0, 0, 0, 0, Paris), ChristmasEve.Calc(2021, Paris))
	assert.Equal(time.Time{}, ChristmasEve.Calc(2022, Paris))
	assert.Equal(time.Time{}, ChristmasEve.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 12, 24, 0, 0, 0, 0, Paris), ChristmasEve.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 12, 24, 0, 0, 0, 0, Paris), ChristmasEve.Calc(2025, Paris))
}

func TestChristmasDay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2020, 12, 25, 0, 0, 0, 0, Paris), ChristmasDay.Calc(2020, Paris))
	assert.Equal(time.Time{}, ChristmasDay.Calc(2021, Paris))
	assert.Equal(time.Time{}, ChristmasDay.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 12, 25, 0, 0, 0, 0, Paris), ChristmasDay.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 12, 25, 0, 0, 0, 0, Paris), ChristmasDay.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 12, 25, 0, 0, 0, 0, Paris), ChristmasDay.Calc(2025, Paris))
}

func TestBoxingDay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Time{}, BoxingDay.Calc(2020, Paris))
	assert.Equal(time.Time{}, BoxingDay.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 12, 26, 0, 0, 0, 0, Paris), BoxingDay.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 12, 26, 0, 0, 0, 0, Paris), BoxingDay.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 12, 26, 0, 0, 0, 0, Paris), BoxingDay.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 12, 26, 0, 0, 0, 0, Paris), BoxingDay.Calc(2025, Paris))
}

func TestNewYearsEve(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2020, 12, 31, 0, 0, 0, 0, Paris), NewYearsEve.Calc(2020, Paris))
	assert.Equal(time.Date(2021, 12, 31, 0, 0, 0, 0, Paris), NewYearsEve.Calc(2021, Paris))
	assert.Equal(time.Time{}, NewYearsEve.Calc(2022, Paris))
	assert.Equal(time.Time{}, NewYearsEve.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 12, 31, 0, 0, 0, 0, Paris), NewYearsEve.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 12, 31, 0, 0, 0, 0, Paris), NewYearsEve.Calc(2025, Paris))
}

// US Holidays

func TestMLK(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 1, 18, 0, 0, 0, 0, NewYork), MLKDay.Calc(2010, NewYork))
	assert.Equal(time.Date(2011, 1, 17, 0, 0, 0, 0, NewYork), MLKDay.Calc(2011, NewYork))
	assert.Equal(time.Date(2012, 1, 16, 0, 0, 0, 0, NewYork), MLKDay.Calc(2012, NewYork))
	assert.Equal(time.Date(2013, 1, 21, 0, 0, 0, 0, NewYork), MLKDay.Calc(2013, NewYork))
	assert.Equal(time.Date(2014, 1, 20, 0, 0, 0, 0, NewYork), MLKDay.Calc(2014, NewYork))
	assert.Equal(time.Date(2015, 1, 19, 0, 0, 0, 0, NewYork), MLKDay.Calc(2015, NewYork))
	assert.Equal(time.Date(2016, 1, 18, 0, 0, 0, 0, NewYork), MLKDay.Calc(2016, NewYork))
	assert.Equal(time.Date(2017, 1, 16, 0, 0, 0, 0, NewYork), MLKDay.Calc(2017, NewYork))
	assert.Equal(time.Date(2018, 1, 15, 0, 0, 0, 0, NewYork), MLKDay.Calc(2018, NewYork))
	assert.Equal(time.Date(2019, 1, 21, 0, 0, 0, 0, NewYork), MLKDay.Calc(2019, NewYork))
	assert.Equal(time.Date(2020, 1, 20, 0, 0, 0, 0, NewYork), MLKDay.Calc(2020, NewYork))
	assert.Equal(time.Date(2021, 1, 18, 0, 0, 0, 0, NewYork), MLKDay.Calc(2021, NewYork))
	assert.Equal(time.Date(2022, 1, 17, 0, 0, 0, 0, NewYork), MLKDay.Calc(2022, NewYork))
	assert.Equal(time.Date(2023, 1, 16, 0, 0, 0, 0, NewYork), MLKDay.Calc(2023, NewYork))
	assert.Equal(time.Date(2024, 1, 15, 0, 0, 0, 0, NewYork), MLKDay.Calc(2024, NewYork))
	assert.Equal(time.Date(2025, 1, 20, 0, 0, 0, 0, NewYork), MLKDay.Calc(2025, NewYork))
}

func TestPresidentsDay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 2, 15, 0, 0, 0, 0, NewYork), PresidentsDay.Calc(2010, NewYork))
	assert.Equal(time.Date(2011, 2, 21, 0, 0, 0, 0, NewYork), PresidentsDay.Calc(2011, NewYork))
	assert.Equal(time.Date(2012, 2, 20, 0, 0, 0, 0, NewYork), PresidentsDay.Calc(2012, NewYork))
	assert.Equal(time.Date(2013, 2, 18, 0, 0, 0, 0, NewYork), PresidentsDay.Calc(2013, NewYork))
	assert.Equal(time.Date(2014, 2, 17, 0, 0, 0, 0, NewYork), PresidentsDay.Calc(2014, NewYork))
	assert.Equal(time.Date(2015, 2, 16, 0, 0, 0, 0, NewYork), PresidentsDay.Calc(2015, NewYork))
	assert.Equal(time.Date(2016, 2, 15, 0, 0, 0, 0, NewYork), PresidentsDay.Calc(2016, NewYork))
	assert.Equal(time.Date(2017, 2, 20, 0, 0, 0, 0, NewYork), PresidentsDay.Calc(2017, NewYork))
	assert.Equal(time.Date(2018, 2, 19, 0, 0, 0, 0, NewYork), PresidentsDay.Calc(2018, NewYork))
	assert.Equal(time.Date(2019, 2, 18, 0, 0, 0, 0, NewYork), PresidentsDay.Calc(2019, NewYork))
	assert.Equal(time.Date(2020, 2, 17, 0, 0, 0, 0, NewYork), PresidentsDay.Calc(2020, NewYork))
	assert.Equal(time.Date(2021, 2, 15, 0, 0, 0, 0, NewYork), PresidentsDay.Calc(2021, NewYork))
	assert.Equal(time.Date(2022, 2, 21, 0, 0, 0, 0, NewYork), PresidentsDay.Calc(2022, NewYork))
	assert.Equal(time.Date(2023, 2, 20, 0, 0, 0, 0, NewYork), PresidentsDay.Calc(2023, NewYork))
	assert.Equal(time.Date(2024, 2, 19, 0, 0, 0, 0, NewYork), PresidentsDay.Calc(2024, NewYork))
	assert.Equal(time.Date(2025, 2, 17, 0, 0, 0, 0, NewYork), PresidentsDay.Calc(2025, NewYork))
}

func TestMemorialDay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 5, 31, 0, 0, 0, 0, NewYork), MemorialDay.Calc(2010, NewYork))
	assert.Equal(time.Date(2011, 5, 30, 0, 0, 0, 0, NewYork), MemorialDay.Calc(2011, NewYork))
	assert.Equal(time.Date(2012, 5, 28, 0, 0, 0, 0, NewYork), MemorialDay.Calc(2012, NewYork))
	assert.Equal(time.Date(2013, 5, 27, 0, 0, 0, 0, NewYork), MemorialDay.Calc(2013, NewYork))
	assert.Equal(time.Date(2014, 5, 26, 0, 0, 0, 0, NewYork), MemorialDay.Calc(2014, NewYork))
	assert.Equal(time.Date(2015, 5, 25, 0, 0, 0, 0, NewYork), MemorialDay.Calc(2015, NewYork))
	assert.Equal(time.Date(2016, 5, 30, 0, 0, 0, 0, NewYork), MemorialDay.Calc(2016, NewYork))
	assert.Equal(time.Date(2017, 5, 29, 0, 0, 0, 0, NewYork), MemorialDay.Calc(2017, NewYork))
	assert.Equal(time.Date(2018, 5, 28, 0, 0, 0, 0, NewYork), MemorialDay.Calc(2018, NewYork))
	assert.Equal(time.Date(2019, 5, 27, 0, 0, 0, 0, NewYork), MemorialDay.Calc(2019, NewYork))
	assert.Equal(time.Date(2020, 5, 25, 0, 0, 0, 0, NewYork), MemorialDay.Calc(2020, NewYork))
	assert.Equal(time.Date(2021, 5, 31, 0, 0, 0, 0, NewYork), MemorialDay.Calc(2021, NewYork))
	assert.Equal(time.Date(2022, 5, 30, 0, 0, 0, 0, NewYork), MemorialDay.Calc(2022, NewYork))
	assert.Equal(time.Date(2023, 5, 29, 0, 0, 0, 0, NewYork), MemorialDay.Calc(2023, NewYork))
	assert.Equal(time.Date(2024, 5, 27, 0, 0, 0, 0, NewYork), MemorialDay.Calc(2024, NewYork))
	assert.Equal(time.Date(2025, 5, 26, 0, 0, 0, 0, NewYork), MemorialDay.Calc(2025, NewYork))
}

func TestIndependenceDay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 7, 5, 0, 0, 0, 0, NewYork), IndependenceDay.Calc(2010, NewYork))
	assert.Equal(time.Date(2011, 7, 4, 0, 0, 0, 0, NewYork), IndependenceDay.Calc(2011, NewYork))
	assert.Equal(time.Date(2012, 7, 4, 0, 0, 0, 0, NewYork), IndependenceDay.Calc(2012, NewYork))
	assert.Equal(time.Date(2013, 7, 4, 0, 0, 0, 0, NewYork), IndependenceDay.Calc(2013, NewYork))
	assert.Equal(time.Date(2014, 7, 4, 0, 0, 0, 0, NewYork), IndependenceDay.Calc(2014, NewYork))
	assert.Equal(time.Date(2015, 7, 3, 0, 0, 0, 0, NewYork), IndependenceDay.Calc(2015, NewYork))
	assert.Equal(time.Date(2016, 7, 4, 0, 0, 0, 0, NewYork), IndependenceDay.Calc(2016, NewYork))
	assert.Equal(time.Date(2017, 7, 4, 0, 0, 0, 0, NewYork), IndependenceDay.Calc(2017, NewYork))
	assert.Equal(time.Date(2018, 7, 4, 0, 0, 0, 0, NewYork), IndependenceDay.Calc(2018, NewYork))
	assert.Equal(time.Date(2019, 7, 4, 0, 0, 0, 0, NewYork), IndependenceDay.Calc(2019, NewYork))
	assert.Equal(time.Date(2020, 7, 3, 0, 0, 0, 0, NewYork), IndependenceDay.Calc(2020, NewYork))
	assert.Equal(time.Date(2021, 7, 5, 0, 0, 0, 0, NewYork), IndependenceDay.Calc(2021, NewYork))
	assert.Equal(time.Date(2022, 7, 4, 0, 0, 0, 0, NewYork), IndependenceDay.Calc(2022, NewYork))
	assert.Equal(time.Date(2023, 7, 4, 0, 0, 0, 0, NewYork), IndependenceDay.Calc(2023, NewYork))
	assert.Equal(time.Date(2024, 7, 4, 0, 0, 0, 0, NewYork), IndependenceDay.Calc(2024, NewYork))
	assert.Equal(time.Date(2025, 7, 4, 0, 0, 0, 0, NewYork), IndependenceDay.Calc(2025, NewYork))
}

func TestLaborDay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 9, 6, 0, 0, 0, 0, NewYork), LaborDay.Calc(2010, NewYork))
	assert.Equal(time.Date(2011, 9, 5, 0, 0, 0, 0, NewYork), LaborDay.Calc(2011, NewYork))
	assert.Equal(time.Date(2012, 9, 3, 0, 0, 0, 0, NewYork), LaborDay.Calc(2012, NewYork))
	assert.Equal(time.Date(2013, 9, 2, 0, 0, 0, 0, NewYork), LaborDay.Calc(2013, NewYork))
	assert.Equal(time.Date(2014, 9, 1, 0, 0, 0, 0, NewYork), LaborDay.Calc(2014, NewYork))
	assert.Equal(time.Date(2015, 9, 7, 0, 0, 0, 0, NewYork), LaborDay.Calc(2015, NewYork))
	assert.Equal(time.Date(2016, 9, 5, 0, 0, 0, 0, NewYork), LaborDay.Calc(2016, NewYork))
	assert.Equal(time.Date(2017, 9, 4, 0, 0, 0, 0, NewYork), LaborDay.Calc(2017, NewYork))
	assert.Equal(time.Date(2018, 9, 3, 0, 0, 0, 0, NewYork), LaborDay.Calc(2018, NewYork))
	assert.Equal(time.Date(2019, 9, 2, 0, 0, 0, 0, NewYork), LaborDay.Calc(2019, NewYork))
	assert.Equal(time.Date(2020, 9, 7, 0, 0, 0, 0, NewYork), LaborDay.Calc(2020, NewYork))
	assert.Equal(time.Date(2021, 9, 6, 0, 0, 0, 0, NewYork), LaborDay.Calc(2021, NewYork))
	assert.Equal(time.Date(2022, 9, 5, 0, 0, 0, 0, NewYork), LaborDay.Calc(2022, NewYork))
	assert.Equal(time.Date(2023, 9, 4, 0, 0, 0, 0, NewYork), LaborDay.Calc(2023, NewYork))
	assert.Equal(time.Date(2024, 9, 2, 0, 0, 0, 0, NewYork), LaborDay.Calc(2024, NewYork))
	assert.Equal(time.Date(2025, 9, 1, 0, 0, 0, 0, NewYork), LaborDay.Calc(2025, NewYork))
}

func TestColombusDay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 10, 11, 0, 0, 0, 0, NewYork), ColumbusDay.Calc(2010, NewYork))
	assert.Equal(time.Date(2011, 10, 10, 0, 0, 0, 0, NewYork), ColumbusDay.Calc(2011, NewYork))
	assert.Equal(time.Date(2012, 10, 8, 0, 0, 0, 0, NewYork), ColumbusDay.Calc(2012, NewYork))
	assert.Equal(time.Date(2013, 10, 14, 0, 0, 0, 0, NewYork), ColumbusDay.Calc(2013, NewYork))
	assert.Equal(time.Date(2014, 10, 13, 0, 0, 0, 0, NewYork), ColumbusDay.Calc(2014, NewYork))
	assert.Equal(time.Date(2015, 10, 12, 0, 0, 0, 0, NewYork), ColumbusDay.Calc(2015, NewYork))
	assert.Equal(time.Date(2016, 10, 10, 0, 0, 0, 0, NewYork), ColumbusDay.Calc(2016, NewYork))
	assert.Equal(time.Date(2017, 10, 9, 0, 0, 0, 0, NewYork), ColumbusDay.Calc(2017, NewYork))
	assert.Equal(time.Date(2018, 10, 8, 0, 0, 0, 0, NewYork), ColumbusDay.Calc(2018, NewYork))
	assert.Equal(time.Date(2019, 10, 14, 0, 0, 0, 0, NewYork), ColumbusDay.Calc(2019, NewYork))
	assert.Equal(time.Date(2020, 10, 12, 0, 0, 0, 0, NewYork), ColumbusDay.Calc(2020, NewYork))
	assert.Equal(time.Date(2021, 10, 11, 0, 0, 0, 0, NewYork), ColumbusDay.Calc(2021, NewYork))
	assert.Equal(time.Date(2022, 10, 10, 0, 0, 0, 0, NewYork), ColumbusDay.Calc(2022, NewYork))
	assert.Equal(time.Date(2023, 10, 9, 0, 0, 0, 0, NewYork), ColumbusDay.Calc(2023, NewYork))
	assert.Equal(time.Date(2024, 10, 14, 0, 0, 0, 0, NewYork), ColumbusDay.Calc(2024, NewYork))
	assert.Equal(time.Date(2025, 10, 13, 0, 0, 0, 0, NewYork), ColumbusDay.Calc(2025, NewYork))
}

func TestVeteransDay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2020, 11, 11, 0, 0, 0, 0, NewYork), VeteransDay.Calc(2020, NewYork))
	assert.Equal(time.Date(2021, 11, 11, 0, 0, 0, 0, NewYork), VeteransDay.Calc(2021, NewYork))
	assert.Equal(time.Date(2022, 11, 11, 0, 0, 0, 0, NewYork), VeteransDay.Calc(2022, NewYork))
	assert.Equal(time.Time{}, VeteransDay.Calc(2023, NewYork))
	assert.Equal(time.Date(2024, 11, 11, 0, 0, 0, 0, NewYork), VeteransDay.Calc(2024, NewYork))
	assert.Equal(time.Date(2025, 11, 11, 0, 0, 0, 0, NewYork), VeteransDay.Calc(2025, NewYork))
}

func TestThanksgivingDay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 11, 25, 0, 0, 0, 0, NewYork), ThanksgivingDay.Calc(2010, NewYork))
	assert.Equal(time.Date(2011, 11, 24, 0, 0, 0, 0, NewYork), ThanksgivingDay.Calc(2011, NewYork))
	assert.Equal(time.Date(2012, 11, 22, 0, 0, 0, 0, NewYork), ThanksgivingDay.Calc(2012, NewYork))
	assert.Equal(time.Date(2013, 11, 28, 0, 0, 0, 0, NewYork), ThanksgivingDay.Calc(2013, NewYork))
	assert.Equal(time.Date(2014, 11, 27, 0, 0, 0, 0, NewYork), ThanksgivingDay.Calc(2014, NewYork))
	assert.Equal(time.Date(2015, 11, 26, 0, 0, 0, 0, NewYork), ThanksgivingDay.Calc(2015, NewYork))
	assert.Equal(time.Date(2016, 11, 24, 0, 0, 0, 0, NewYork), ThanksgivingDay.Calc(2016, NewYork))
	assert.Equal(time.Date(2017, 11, 23, 0, 0, 0, 0, NewYork), ThanksgivingDay.Calc(2017, NewYork))
	assert.Equal(time.Date(2018, 11, 22, 0, 0, 0, 0, NewYork), ThanksgivingDay.Calc(2018, NewYork))
	assert.Equal(time.Date(2019, 11, 28, 0, 0, 0, 0, NewYork), ThanksgivingDay.Calc(2019, NewYork))
	assert.Equal(time.Date(2020, 11, 26, 0, 0, 0, 0, NewYork), ThanksgivingDay.Calc(2020, NewYork))
	assert.Equal(time.Date(2021, 11, 25, 0, 0, 0, 0, NewYork), ThanksgivingDay.Calc(2021, NewYork))
	assert.Equal(time.Date(2022, 11, 24, 0, 0, 0, 0, NewYork), ThanksgivingDay.Calc(2022, NewYork))
	assert.Equal(time.Date(2023, 11, 23, 0, 0, 0, 0, NewYork), ThanksgivingDay.Calc(2023, NewYork))
	assert.Equal(time.Date(2024, 11, 28, 0, 0, 0, 0, NewYork), ThanksgivingDay.Calc(2024, NewYork))
	assert.Equal(time.Date(2025, 11, 27, 0, 0, 0, 0, NewYork), ThanksgivingDay.Calc(2025, NewYork))
}

func TestBlackFriday(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 11, 25+1, 0, 0, 0, 0, NewYork), BlackFriday.Calc(2010, NewYork))
	assert.Equal(time.Date(2011, 11, 24+1, 0, 0, 0, 0, NewYork), BlackFriday.Calc(2011, NewYork))
	assert.Equal(time.Date(2012, 11, 22+1, 0, 0, 0, 0, NewYork), BlackFriday.Calc(2012, NewYork))
	assert.Equal(time.Date(2013, 11, 28+1, 0, 0, 0, 0, NewYork), BlackFriday.Calc(2013, NewYork))
	assert.Equal(time.Date(2014, 11, 27+1, 0, 0, 0, 0, NewYork), BlackFriday.Calc(2014, NewYork))
	assert.Equal(time.Date(2015, 11, 26+1, 0, 0, 0, 0, NewYork), BlackFriday.Calc(2015, NewYork))
	assert.Equal(time.Date(2016, 11, 24+1, 0, 0, 0, 0, NewYork), BlackFriday.Calc(2016, NewYork))
	assert.Equal(time.Date(2017, 11, 23+1, 0, 0, 0, 0, NewYork), BlackFriday.Calc(2017, NewYork))
	assert.Equal(time.Date(2018, 11, 22+1, 0, 0, 0, 0, NewYork), BlackFriday.Calc(2018, NewYork))
	assert.Equal(time.Date(2019, 11, 28+1, 0, 0, 0, 0, NewYork), BlackFriday.Calc(2019, NewYork))
	assert.Equal(time.Date(2020, 11, 26+1, 0, 0, 0, 0, NewYork), BlackFriday.Calc(2020, NewYork))
	assert.Equal(time.Date(2021, 11, 25+1, 0, 0, 0, 0, NewYork), BlackFriday.Calc(2021, NewYork))
	assert.Equal(time.Date(2022, 11, 24+1, 0, 0, 0, 0, NewYork), BlackFriday.Calc(2022, NewYork))
	assert.Equal(time.Date(2023, 11, 23+1, 0, 0, 0, 0, NewYork), BlackFriday.Calc(2023, NewYork))
	assert.Equal(time.Date(2024, 11, 28+1, 0, 0, 0, 0, NewYork), BlackFriday.Calc(2024, NewYork))
	assert.Equal(time.Date(2025, 11, 27+1, 0, 0, 0, 0, NewYork), BlackFriday.Calc(2025, NewYork))
}

// UK

func TestEarlyMay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 5, 3, 0, 0, 0, 0, NewYork), EarlyMay.Calc(2010, NewYork))
	assert.Equal(time.Date(2011, 5, 2, 0, 0, 0, 0, NewYork), EarlyMay.Calc(2011, NewYork))
	assert.Equal(time.Date(2012, 5, 7, 0, 0, 0, 0, NewYork), EarlyMay.Calc(2012, NewYork))
	assert.Equal(time.Date(2013, 5, 6, 0, 0, 0, 0, NewYork), EarlyMay.Calc(2013, NewYork))
	assert.Equal(time.Date(2014, 5, 5, 0, 0, 0, 0, NewYork), EarlyMay.Calc(2014, NewYork))
	assert.Equal(time.Date(2015, 5, 4, 0, 0, 0, 0, NewYork), EarlyMay.Calc(2015, NewYork))
	assert.Equal(time.Date(2016, 5, 2, 0, 0, 0, 0, NewYork), EarlyMay.Calc(2016, NewYork))
	assert.Equal(time.Date(2017, 5, 1, 0, 0, 0, 0, NewYork), EarlyMay.Calc(2017, NewYork))
	assert.Equal(time.Date(2018, 5, 7, 0, 0, 0, 0, NewYork), EarlyMay.Calc(2018, NewYork))
	assert.Equal(time.Date(2019, 5, 6, 0, 0, 0, 0, NewYork), EarlyMay.Calc(2019, NewYork))
	assert.Equal(time.Date(2020, 5, 4, 0, 0, 0, 0, NewYork), EarlyMay.Calc(2020, NewYork))
	assert.Equal(time.Date(2021, 5, 3, 0, 0, 0, 0, NewYork), EarlyMay.Calc(2021, NewYork))
	assert.Equal(time.Date(2022, 5, 2, 0, 0, 0, 0, NewYork), EarlyMay.Calc(2022, NewYork))
	assert.Equal(time.Date(2023, 5, 1, 0, 0, 0, 0, NewYork), EarlyMay.Calc(2023, NewYork))
	assert.Equal(time.Date(2024, 5, 6, 0, 0, 0, 0, NewYork), EarlyMay.Calc(2024, NewYork))
	assert.Equal(time.Date(2025, 5, 5, 0, 0, 0, 0, NewYork), EarlyMay.Calc(2025, NewYork))
}

func TestLateMay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 5, 31, 0, 0, 0, 0, NewYork), LateMay.Calc(2010, NewYork))
	assert.Equal(time.Date(2011, 5, 30, 0, 0, 0, 0, NewYork), LateMay.Calc(2011, NewYork))
	assert.Equal(time.Date(2012, 5, 28, 0, 0, 0, 0, NewYork), LateMay.Calc(2012, NewYork))
	assert.Equal(time.Date(2013, 5, 27, 0, 0, 0, 0, NewYork), LateMay.Calc(2013, NewYork))
	assert.Equal(time.Date(2014, 5, 26, 0, 0, 0, 0, NewYork), LateMay.Calc(2014, NewYork))
	assert.Equal(time.Date(2015, 5, 25, 0, 0, 0, 0, NewYork), LateMay.Calc(2015, NewYork))
	assert.Equal(time.Date(2016, 5, 30, 0, 0, 0, 0, NewYork), LateMay.Calc(2016, NewYork))
	assert.Equal(time.Date(2017, 5, 29, 0, 0, 0, 0, NewYork), LateMay.Calc(2017, NewYork))
	assert.Equal(time.Date(2018, 5, 28, 0, 0, 0, 0, NewYork), LateMay.Calc(2018, NewYork))
	assert.Equal(time.Date(2019, 5, 27, 0, 0, 0, 0, NewYork), LateMay.Calc(2019, NewYork))
	assert.Equal(time.Date(2020, 5, 25, 0, 0, 0, 0, NewYork), LateMay.Calc(2020, NewYork))
	assert.Equal(time.Date(2021, 5, 31, 0, 0, 0, 0, NewYork), LateMay.Calc(2021, NewYork))
	assert.Equal(time.Date(2022, 5, 30, 0, 0, 0, 0, NewYork), LateMay.Calc(2022, NewYork))
	assert.Equal(time.Date(2023, 5, 29, 0, 0, 0, 0, NewYork), LateMay.Calc(2023, NewYork))
	assert.Equal(time.Date(2024, 5, 27, 0, 0, 0, 0, NewYork), LateMay.Calc(2024, NewYork))
	assert.Equal(time.Date(2025, 5, 26, 0, 0, 0, 0, NewYork), LateMay.Calc(2025, NewYork))
}

func TestSummerHoliday(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 8, 30, 0, 0, 0, 0, NewYork), SummerHoliday.Calc(2010, NewYork))
	assert.Equal(time.Date(2011, 8, 29, 0, 0, 0, 0, NewYork), SummerHoliday.Calc(2011, NewYork))
	assert.Equal(time.Date(2012, 8, 27, 0, 0, 0, 0, NewYork), SummerHoliday.Calc(2012, NewYork))
	assert.Equal(time.Date(2013, 8, 26, 0, 0, 0, 0, NewYork), SummerHoliday.Calc(2013, NewYork))
	assert.Equal(time.Date(2014, 8, 25, 0, 0, 0, 0, NewYork), SummerHoliday.Calc(2014, NewYork))
	assert.Equal(time.Date(2015, 8, 31, 0, 0, 0, 0, NewYork), SummerHoliday.Calc(2015, NewYork))
	assert.Equal(time.Date(2016, 8, 29, 0, 0, 0, 0, NewYork), SummerHoliday.Calc(2016, NewYork))
	assert.Equal(time.Date(2017, 8, 28, 0, 0, 0, 0, NewYork), SummerHoliday.Calc(2017, NewYork))
	assert.Equal(time.Date(2018, 8, 27, 0, 0, 0, 0, NewYork), SummerHoliday.Calc(2018, NewYork))
	assert.Equal(time.Date(2019, 8, 26, 0, 0, 0, 0, NewYork), SummerHoliday.Calc(2019, NewYork))
	assert.Equal(time.Date(2020, 8, 31, 0, 0, 0, 0, NewYork), SummerHoliday.Calc(2020, NewYork))
	assert.Equal(time.Date(2021, 8, 30, 0, 0, 0, 0, NewYork), SummerHoliday.Calc(2021, NewYork))
	assert.Equal(time.Date(2022, 8, 29, 0, 0, 0, 0, NewYork), SummerHoliday.Calc(2022, NewYork))
	assert.Equal(time.Date(2023, 8, 28, 0, 0, 0, 0, NewYork), SummerHoliday.Calc(2023, NewYork))
	assert.Equal(time.Date(2024, 8, 26, 0, 0, 0, 0, NewYork), SummerHoliday.Calc(2024, NewYork))
	assert.Equal(time.Date(2025, 8, 25, 0, 0, 0, 0, NewYork), SummerHoliday.Calc(2025, NewYork))
}

func TestQueensday(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2020, 4, 30, 0, 0, 0, 0, NewYork), QueensDay.Calc(2020, NewYork))
	assert.Equal(time.Date(2021, 4, 30, 0, 0, 0, 0, NewYork), QueensDay.Calc(2021, NewYork))
	assert.Equal(time.Time{}, QueensDay.Calc(2022, NewYork))
	assert.Equal(time.Time{}, QueensDay.Calc(2023, NewYork))
	assert.Equal(time.Date(2024, 4, 30, 0, 0, 0, 0, NewYork), QueensDay.Calc(2024, NewYork))
	assert.Equal(time.Date(2025, 4, 30, 0, 0, 0, 0, NewYork), QueensDay.Calc(2025, NewYork))
}
