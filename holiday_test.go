package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	assert := assert.New(t)
	h1 := &Holiday{
		Name:  "Christmas",
		Month: time.December,
		Day:   25,
	}
	h2 := h1.Copy("USChristmas")
	h2.SetObservance(nearestWorkday)
	assert.Equal("Christmas", h1.Name)
	assert.Equal("USChristmas", h2.Name)
	assert.Empty(h1.Observance)
	assert.NotEmpty(h2.Observance)
	h3 := h1.Copy()
	assert.Equal("Christmas", h3.Name)
}

func TestOffset(t *testing.T) {
	assert := assert.New(t)
	h := &Holiday{
		Name:  "Christmas",
		Month: time.December,
		Day:   25,
	}
	assert.Equal(0, h.Offset)
	h.SetOffset(1)
	assert.Equal(1, h.Offset)
}
func TestObservance(t *testing.T) {
	assert := assert.New(t)
	h := &Holiday{
		Name:  "Christmas",
		Month: time.December,
		Day:   25,
	}
	assert.Nil(h.Observance)
	h.SetObservance(nearestWorkday)
	assert.NotNil(h.Observance)
}

func TestCalc(t *testing.T) {
	assert := assert.New(t)
	h := &Holiday{
		Name:  "A holiday",
		Month: time.March,
		Day:   17,
		calc:  CalcDayOfMonth,
	}
	assert.Equal(time.Date(2020, 3, 17, 0, 0, 0, 0, Paris), h.Calc(2020, Paris))
	h.Offset = 1
	assert.Equal(time.Date(2020, 3, 18, 0, 0, 0, 0, Paris), h.Calc(2020, Paris))
	h.Offset = -1
	assert.Equal(time.Date(2020, 3, 16, 0, 0, 0, 0, Paris), h.Calc(2020, Paris))
	h.Offset = -2
	assert.Equal(time.Time{}, h.Calc(2020, Paris))
	h.SetObservance(nearestWorkday)
	assert.Equal(time.Date(2020, 3, 16, 0, 0, 0, 0, Paris), h.Calc(2020, Paris))
	h.OnYear = 2020
	assert.Equal(time.Date(2020, 3, 16, 0, 0, 0, 0, Paris), h.Calc(2020, Paris))
	h.OnYear = 2021
	assert.Equal(time.Time{}, h.Calc(2020, Paris))
}

func TestCalcDayOfMonth(t *testing.T) {
	assert := assert.New(t)
	h := &Holiday{Month: time.December, Day: 25}
	assert.Equal(time.Date(2020, 12, 25, 0, 0, 0, 0, Paris), CalcDayOfMonth(h, 2020, Paris))
	assert.Equal(time.Date(2021, 12, 25, 0, 0, 0, 0, Paris), CalcDayOfMonth(h, 2021, Paris))
}

func TestCalcLunarNewYear(t *testing.T) {
	assert := assert.New(t)
	h := &Holiday{Month: time.January, Day: 1}
	assert.Equal(time.Date(2020, 1, 25, 0, 0, 0, 0, HongKong), CalcLunarDayOfMonth(h, 2020, HongKong))
	assert.Equal(time.Date(2021, 2, 12, 0, 0, 0, 0, HongKong), CalcLunarDayOfMonth(h, 2021, HongKong))
}

func TestCalcNthWeekday(t *testing.T) {
	assert := assert.New(t)
	h := &Holiday{Month: time.December, Weekday: time.Thursday, NthWeekday: 3}
	assert.Equal(time.Date(2020, 12, 17, 0, 0, 0, 0, Paris), CalcNthWeekday(h, 2020, Paris))
}

func TestCalcEasterOffset(t *testing.T) {
	assert := assert.New(t)
	h := &Holiday{}
	assert.Equal(time.Date(2020, 4, 12, 0, 0, 0, 0, Paris), CalcEasterOffset(h, 2020, Paris))
	assert.Equal(time.Date(2021, 4, 4, 0, 0, 0, 0, Paris), CalcEasterOffset(h, 2021, Paris))
}

func TestCalcNorthwardEquinox(t *testing.T) {
	assert := assert.New(t)
	h := &Holiday{}
	assert.Equal(time.Date(2020, 3, 20, 0, 0, 0, 0, Paris), CalcNorthwardEquinox(h, 2020, Paris))
	assert.Equal(time.Date(2020, 3, 20, 0, 0, 0, 0, Tokyo), CalcNorthwardEquinox(h, 2020, Tokyo))
}

func TestCalcNorthernSolstice(t *testing.T) {
	assert := assert.New(t)
	h := &Holiday{}
	assert.Equal(time.Date(2020, 6, 20, 0, 0, 0, 0, Paris), CalcNorthernSolstice(h, 2020, Paris))
	assert.Equal(time.Date(2020, 6, 21, 0, 0, 0, 0, Tokyo), CalcNorthernSolstice(h, 2020, Tokyo))
}

func TestCalcSouthwardEquinox(t *testing.T) {
	assert := assert.New(t)
	h := &Holiday{}
	assert.Equal(time.Date(2020, 9, 22, 0, 0, 0, 0, Paris), CalcSouthwardEquinox(h, 2020, Paris))
	assert.Equal(time.Date(2020, 9, 22, 0, 0, 0, 0, Tokyo), CalcSouthwardEquinox(h, 2020, Tokyo))
}

func TestCalcSouthernSolstice(t *testing.T) {
	assert := assert.New(t)
	h := &Holiday{}
	assert.Equal(time.Date(2020, 12, 21, 0, 0, 0, 0, Paris), CalcSouthernSolstice(h, 2020, Paris))
	assert.Equal(time.Date(2020, 12, 21, 0, 0, 0, 0, Tokyo), CalcSouthernSolstice(h, 2020, Tokyo))
}
