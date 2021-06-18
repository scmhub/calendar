package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLeapMonth(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0, GetLeapMonth(2015))
	assert.Equal(0, GetLeapMonth(2016))
	assert.Equal(6, GetLeapMonth(2017))
	assert.Equal(0, GetLeapMonth(2018))
	assert.Equal(0, GetLeapMonth(2019))
	assert.Equal(4, GetLeapMonth(2020))
	assert.Equal(0, GetLeapMonth(2021))
	assert.Equal(0, GetLeapMonth(2022))
	assert.Equal(2, GetLeapMonth(2023))
	assert.Equal(0, GetLeapMonth(2024))
	assert.Equal(6, GetLeapMonth(2025))
}

func TestLunarConv(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		solar time.Time
		lunar time.Time
		leap  bool
	}{
		{
			solar: time.Date(2015, 1, 14, 12, 30, 0, 0, Paris),
			lunar: time.Date(2014, 11, 24, 12, 30, 0, 0, Paris),
			leap:  false,
		},
		{
			solar: time.Date(2015, 2, 19, 12, 30, 0, 0, Paris),
			lunar: time.Date(2015, 1, 1, 12, 30, 0, 0, Paris),
			leap:  false,
		},
		{
			solar: time.Date(2015, 4, 14, 12, 30, 0, 0, Paris),
			lunar: time.Date(2015, 2, 26, 12, 30, 0, 0, Paris),
			leap:  false,
		},
		{
			solar: time.Date(2020, 1, 14, 12, 30, 0, 0, Paris),
			lunar: time.Date(2019, 12, 20, 12, 30, 0, 0, Paris),
			leap:  false,
		},
		{
			solar: time.Date(2020, 1, 25, 12, 30, 0, 0, Paris),
			lunar: time.Date(2020, 1, 1, 12, 30, 0, 0, Paris),
			leap:  false,
		},
		{
			solar: time.Date(2020, 5, 14, 12, 30, 0, 0, Paris),
			lunar: time.Date(2020, 4, 22, 12, 30, 0, 0, Paris),
			leap:  false,
		},
		{
			solar: time.Date(2020, 6, 13, 12, 30, 0, 0, Paris),
			lunar: time.Date(2020, 4, 22, 12, 30, 0, 0, Paris),
			leap:  true,
		},

		{
			solar: time.Date(2025, 1, 14, 12, 30, 0, 0, Paris),
			lunar: time.Date(2024, 12, 15, 12, 30, 0, 0, Paris),
			leap:  false,
		},
		{
			solar: time.Date(2025, 1, 29, 12, 30, 0, 0, Paris),
			lunar: time.Date(2025, 1, 1, 12, 30, 0, 0, Paris),
			leap:  false,
		},
		{
			solar: time.Date(2025, 7, 15, 12, 30, 0, 0, Paris),
			lunar: time.Date(2025, 6, 21, 12, 30, 0, 0, Paris),
			leap:  false,
		},
		{
			solar: time.Date(2025, 8, 14, 12, 30, 0, 0, Paris),
			lunar: time.Date(2025, 6, 21, 12, 30, 0, 0, Paris),
			leap:  true,
		},
	}
	for _, test := range tests {
		toLunarTime, toLunarLeap := GregorianToLunisolar(test.solar)
		assert.Equal(test.lunar, toLunarTime)
		assert.Equal(test.leap, toLunarLeap)
		assert.Equal(test.solar, LunisolarToGregorian(test.lunar, test.leap))
	}
}

func TestGetLunarMonthDays(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(29, GetLunisolarMonthDays(2020, time.Month(1), false))
	assert.Equal(30, GetLunisolarMonthDays(2020, time.Month(2), false))
	assert.Equal(30, GetLunisolarMonthDays(2020, time.Month(3), false))
	assert.Equal(30, GetLunisolarMonthDays(2020, time.Month(4), false))
	assert.Equal(29, GetLunisolarMonthDays(2020, time.Month(4), true))
	assert.Equal(30, GetLunisolarMonthDays(2020, time.Month(5), false))
	assert.Equal(29, GetLunisolarMonthDays(2020, time.Month(6), false))
	assert.Equal(29, GetLunisolarMonthDays(2020, time.Month(7), false))
	assert.Equal(30, GetLunisolarMonthDays(2020, time.Month(8), false))
	assert.Equal(29, GetLunisolarMonthDays(2020, time.Month(9), false))
	assert.Equal(30, GetLunisolarMonthDays(2020, time.Month(10), false))
	assert.Equal(29, GetLunisolarMonthDays(2020, time.Month(11), false))
	assert.Equal(30, GetLunisolarMonthDays(2020, time.Month(12), false))
	assert.Equal(29, GetLunisolarMonthDays(2021, time.Month(1), false))
	assert.Equal(30, GetLunisolarMonthDays(2021, time.Month(2), false))
	assert.Equal(30, GetLunisolarMonthDays(2021, time.Month(3), false))
	assert.Equal(29, GetLunisolarMonthDays(2021, time.Month(4), false))
	assert.Equal(30, GetLunisolarMonthDays(2021, time.Month(5), false))
	assert.Equal(29, GetLunisolarMonthDays(2021, time.Month(6), false))
	assert.Equal(30, GetLunisolarMonthDays(2021, time.Month(7), false))
	assert.Equal(29, GetLunisolarMonthDays(2021, time.Month(8), false))
	assert.Equal(30, GetLunisolarMonthDays(2021, time.Month(9), false))
	assert.Equal(29, GetLunisolarMonthDays(2021, time.Month(10), false))
	assert.Equal(30, GetLunisolarMonthDays(2021, time.Month(11), false))
	assert.Equal(29, GetLunisolarMonthDays(2021, time.Month(12), false))
}
