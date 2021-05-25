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

func TestCalcDayOfMonth(t *testing.T) {
	assert := assert.New(t)
	h := &Holiday{Month: time.December, Day: 25}
	paris, _ := time.LoadLocation("Europe/Paris")
	assert.Equal(time.Date(2020, 12, 25, 0, 0, 0, 0, paris), CalcDayOfMonth(h, 2020, paris))
	assert.Equal(time.Date(2021, 12, 25, 0, 0, 0, 0, paris), CalcDayOfMonth(h, 2021, paris))
}
