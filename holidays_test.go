package calendar

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
	assert.Equal(time.Date(2010, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2010, Paris))
	assert.Equal(time.Date(2011, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2011, Paris))
	assert.Equal(time.Date(2012, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2012, Paris))
	assert.Equal(time.Time{}, Epiphany.Calc(2013, Paris))
	assert.Equal(time.Date(2014, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2014, Paris))
	assert.Equal(time.Date(2015, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2015, Paris))
	assert.Equal(time.Date(2016, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2016, Paris))
	assert.Equal(time.Date(2017, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2017, Paris))
	assert.Equal(time.Time{}, Epiphany.Calc(2018, Paris))
	assert.Equal(time.Time{}, Epiphany.Calc(2019, Paris))
	assert.Equal(time.Date(2020, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2020, Paris))
	assert.Equal(time.Date(2021, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2023, Paris))
	assert.Equal(time.Time{}, Epiphany.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2025, Paris))
}

func TestMaundyThursday(t *testing.T) {
	assert := assert.New(t)
	fmt.Println(MaundyThursday.Calc(2012, Paris))
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
	fmt.Println(MaundyThursday.Calc(2012, Paris))
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

func TestEasterMonday(t *testing.T) {
	assert := assert.New(t)
	fmt.Println(MaundyThursday.Calc(2012, Paris))
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
	assert.Equal(time.Time{}, WorkersDay.Calc(2010, Paris))
	assert.Equal(time.Time{}, WorkersDay.Calc(2011, Paris))
	assert.Equal(time.Date(2012, 5, 1, 0, 0, 0, 0, Paris), WorkersDay.Calc(2012, Paris))
	assert.Equal(time.Date(2013, 5, 1, 0, 0, 0, 0, Paris), WorkersDay.Calc(2013, Paris))
	assert.Equal(time.Date(2014, 5, 1, 0, 0, 0, 0, Paris), WorkersDay.Calc(2014, Paris))
	assert.Equal(time.Date(2015, 5, 1, 0, 0, 0, 0, Paris), WorkersDay.Calc(2015, Paris))
	assert.Equal(time.Time{}, WorkersDay.Calc(2016, Paris))
	assert.Equal(time.Date(2017, 5, 1, 0, 0, 0, 0, Paris), WorkersDay.Calc(2017, Paris))
	assert.Equal(time.Date(2018, 5, 1, 0, 0, 0, 0, Paris), WorkersDay.Calc(2018, Paris))
	assert.Equal(time.Date(2019, 5, 1, 0, 0, 0, 0, Paris), WorkersDay.Calc(2019, Paris))
	assert.Equal(time.Date(2020, 5, 1, 0, 0, 0, 0, Paris), WorkersDay.Calc(2020, Paris))
	assert.Equal(time.Time{}, WorkersDay.Calc(2021, Paris))
	assert.Equal(time.Time{}, WorkersDay.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 5, 1, 0, 0, 0, 0, Paris), WorkersDay.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 5, 1, 0, 0, 0, 0, Paris), WorkersDay.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 5, 1, 0, 0, 0, 0, Paris), WorkersDay.Calc(2025, Paris))
}

func TestAscensionDay(t *testing.T) {
	assert := assert.New(t)
	fmt.Println(MaundyThursday.Calc(2012, Paris))
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
	fmt.Println(MaundyThursday.Calc(2012, Paris))
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

func TestAllSaintsDay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 11, 1, 0, 0, 0, 0, Paris), AllSaintsDay.Calc(2010, Paris))
	assert.Equal(time.Date(2011, 11, 1, 0, 0, 0, 0, Paris), AllSaintsDay.Calc(2011, Paris))
	assert.Equal(time.Date(2012, 11, 1, 0, 0, 0, 0, Paris), AllSaintsDay.Calc(2012, Paris))
	assert.Equal(time.Date(2013, 11, 1, 0, 0, 0, 0, Paris), AllSaintsDay.Calc(2013, Paris))
	assert.Equal(time.Time{}, AllSaintsDay.Calc(2014, Paris))
	assert.Equal(time.Time{}, AllSaintsDay.Calc(2015, Paris))
	assert.Equal(time.Date(2016, 11, 1, 0, 0, 0, 0, Paris), AllSaintsDay.Calc(2016, Paris))
	assert.Equal(time.Date(2017, 11, 1, 0, 0, 0, 0, Paris), AllSaintsDay.Calc(2017, Paris))
	assert.Equal(time.Date(2018, 11, 1, 0, 0, 0, 0, Paris), AllSaintsDay.Calc(2018, Paris))
	assert.Equal(time.Date(2019, 11, 1, 0, 0, 0, 0, Paris), AllSaintsDay.Calc(2019, Paris))
	assert.Equal(time.Time{}, AllSaintsDay.Calc(2020, Paris))
	assert.Equal(time.Date(2021, 11, 1, 0, 0, 0, 0, Paris), AllSaintsDay.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 11, 1, 0, 0, 0, 0, Paris), AllSaintsDay.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 11, 1, 0, 0, 0, 0, Paris), AllSaintsDay.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 11, 1, 0, 0, 0, 0, Paris), AllSaintsDay.Calc(2024, Paris))
	assert.Equal(time.Time{}, AllSaintsDay.Calc(2025, Paris))
}

func TestArmisticeDay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 11, 11, 0, 0, 0, 0, Paris), ArmisticeDay.Calc(2010, Paris))
	assert.Equal(time.Date(2011, 11, 11, 0, 0, 0, 0, Paris), ArmisticeDay.Calc(2011, Paris))
	assert.Equal(time.Time{}, ArmisticeDay.Calc(2012, Paris))
	assert.Equal(time.Date(2013, 11, 11, 0, 0, 0, 0, Paris), ArmisticeDay.Calc(2013, Paris))
	assert.Equal(time.Date(2014, 11, 11, 0, 0, 0, 0, Paris), ArmisticeDay.Calc(2014, Paris))
	assert.Equal(time.Date(2015, 11, 11, 0, 0, 0, 0, Paris), ArmisticeDay.Calc(2015, Paris))
	assert.Equal(time.Date(2016, 11, 11, 0, 0, 0, 0, Paris), ArmisticeDay.Calc(2016, Paris))
	assert.Equal(time.Time{}, ArmisticeDay.Calc(2017, Paris))
	assert.Equal(time.Time{}, ArmisticeDay.Calc(2018, Paris))
	assert.Equal(time.Date(2019, 11, 11, 0, 0, 0, 0, Paris), ArmisticeDay.Calc(2019, Paris))
	assert.Equal(time.Date(2020, 11, 11, 0, 0, 0, 0, Paris), ArmisticeDay.Calc(2020, Paris))
	assert.Equal(time.Date(2021, 11, 11, 0, 0, 0, 0, Paris), ArmisticeDay.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 11, 11, 0, 0, 0, 0, Paris), ArmisticeDay.Calc(2022, Paris))
	assert.Equal(time.Time{}, ArmisticeDay.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 11, 11, 0, 0, 0, 0, Paris), ArmisticeDay.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 11, 11, 0, 0, 0, 0, Paris), ArmisticeDay.Calc(2025, Paris))
}

func TestImmaculateConception(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 12, 8, 0, 0, 0, 0, Paris), ImmaculateConception.Calc(2010, Paris))
	assert.Equal(time.Date(2011, 12, 8, 0, 0, 0, 0, Paris), ImmaculateConception.Calc(2011, Paris))
	assert.Equal(time.Time{}, ImmaculateConception.Calc(2012, Paris))
	assert.Equal(time.Time{}, ImmaculateConception.Calc(2013, Paris))
	assert.Equal(time.Date(2014, 12, 8, 0, 0, 0, 0, Paris), ImmaculateConception.Calc(2014, Paris))
	assert.Equal(time.Date(2015, 12, 8, 0, 0, 0, 0, Paris), ImmaculateConception.Calc(2015, Paris))
	assert.Equal(time.Date(2016, 12, 8, 0, 0, 0, 0, Paris), ImmaculateConception.Calc(2016, Paris))
	assert.Equal(time.Date(2017, 12, 8, 0, 0, 0, 0, Paris), ImmaculateConception.Calc(2017, Paris))
	assert.Equal(time.Time{}, ImmaculateConception.Calc(2018, Paris))
	assert.Equal(time.Time{}, ImmaculateConception.Calc(2019, Paris))
	assert.Equal(time.Date(2020, 12, 8, 0, 0, 0, 0, Paris), ImmaculateConception.Calc(2020, Paris))
	assert.Equal(time.Date(2021, 12, 8, 0, 0, 0, 0, Paris), ImmaculateConception.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 12, 8, 0, 0, 0, 0, Paris), ImmaculateConception.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 12, 8, 0, 0, 0, 0, Paris), ImmaculateConception.Calc(2023, Paris))
	assert.Equal(time.Time{}, ImmaculateConception.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 12, 8, 0, 0, 0, 0, Paris), ImmaculateConception.Calc(2025, Paris))
}
func TestChristmasDay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Time{}, ChristmasDay.Calc(2010, Paris))
	assert.Equal(time.Time{}, ChristmasDay.Calc(2011, Paris))
	assert.Equal(time.Date(2012, 12, 25, 0, 0, 0, 0, Paris), ChristmasDay.Calc(2012, Paris))
	assert.Equal(time.Date(2013, 12, 25, 0, 0, 0, 0, Paris), ChristmasDay.Calc(2013, Paris))
	assert.Equal(time.Date(2014, 12, 25, 0, 0, 0, 0, Paris), ChristmasDay.Calc(2014, Paris))
	assert.Equal(time.Date(2015, 12, 25, 0, 0, 0, 0, Paris), ChristmasDay.Calc(2015, Paris))
	assert.Equal(time.Time{}, ChristmasDay.Calc(2016, Paris))
	assert.Equal(time.Date(2017, 12, 25, 0, 0, 0, 0, Paris), ChristmasDay.Calc(2017, Paris))
	assert.Equal(time.Date(2018, 12, 25, 0, 0, 0, 0, Paris), ChristmasDay.Calc(2018, Paris))
	assert.Equal(time.Date(2019, 12, 25, 0, 0, 0, 0, Paris), ChristmasDay.Calc(2019, Paris))
	assert.Equal(time.Date(2020, 12, 25, 0, 0, 0, 0, Paris), ChristmasDay.Calc(2020, Paris))
	assert.Equal(time.Time{}, ChristmasDay.Calc(2021, Paris))
	assert.Equal(time.Time{}, ChristmasDay.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 12, 25, 0, 0, 0, 0, Paris), ChristmasDay.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 12, 25, 0, 0, 0, 0, Paris), ChristmasDay.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 12, 25, 0, 0, 0, 0, Paris), ChristmasDay.Calc(2025, Paris))
}
func TestBoxingDay(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Time{}, BoxingDay.Calc(2010, Paris))
	assert.Equal(time.Date(2011, 12, 26, 0, 0, 0, 0, Paris), BoxingDay.Calc(2011, Paris))
	assert.Equal(time.Date(2012, 12, 26, 0, 0, 0, 0, Paris), BoxingDay.Calc(2012, Paris))
	assert.Equal(time.Date(2013, 12, 26, 0, 0, 0, 0, Paris), BoxingDay.Calc(2013, Paris))
	assert.Equal(time.Date(2014, 12, 26, 0, 0, 0, 0, Paris), BoxingDay.Calc(2014, Paris))
	assert.Equal(time.Time{}, BoxingDay.Calc(2015, Paris))
	assert.Equal(time.Date(2016, 12, 26, 0, 0, 0, 0, Paris), BoxingDay.Calc(2016, Paris))
	assert.Equal(time.Date(2017, 12, 26, 0, 0, 0, 0, Paris), BoxingDay.Calc(2017, Paris))
	assert.Equal(time.Date(2018, 12, 26, 0, 0, 0, 0, Paris), BoxingDay.Calc(2018, Paris))
	assert.Equal(time.Date(2019, 12, 26, 0, 0, 0, 0, Paris), BoxingDay.Calc(2019, Paris))
	assert.Equal(time.Time{}, BoxingDay.Calc(2020, Paris))
	assert.Equal(time.Time{}, BoxingDay.Calc(2021, Paris))
	assert.Equal(time.Date(2022, 12, 26, 0, 0, 0, 0, Paris), BoxingDay.Calc(2022, Paris))
	assert.Equal(time.Date(2023, 12, 26, 0, 0, 0, 0, Paris), BoxingDay.Calc(2023, Paris))
	assert.Equal(time.Date(2024, 12, 26, 0, 0, 0, 0, Paris), BoxingDay.Calc(2024, Paris))
	assert.Equal(time.Date(2025, 12, 26, 0, 0, 0, 0, Paris), BoxingDay.Calc(2025, Paris))
}
