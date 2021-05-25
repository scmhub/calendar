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
	assert.Empty(h1.Observance())
	assert.NotEmpty(h2.observance)
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
	assert.Equal(time.Date(2020, 3, 15, 0, 0, 0, 0, Paris), h.Calc(2020, Paris))
	h.SetObservance(nearestWorkday)
	assert.Equal(time.Date(2020, 3, 16, 0, 0, 0, 0, Paris), h.Calc(2020, Paris))
	h.Year = 2020
	assert.Equal(time.Date(2020, 3, 16, 0, 0, 0, 0, Paris), h.Calc(2020, Paris))
	h.Year = 2021
	assert.Equal(time.Time{}, h.Calc(2020, Paris))
}

func TestCalcDayOfMonth(t *testing.T) {
	assert := assert.New(t)
	h := &Holiday{Month: time.December, Day: 25}
	assert.Equal(time.Date(2020, 12, 25, 0, 0, 0, 0, Paris), CalcDayOfMonth(h, 2020, Paris))
	assert.Equal(time.Date(2021, 12, 25, 0, 0, 0, 0, Paris), CalcDayOfMonth(h, 2021, Paris))
}
