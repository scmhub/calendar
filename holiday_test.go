package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalcDayOfMonth(t *testing.T) {
	assert := assert.New(t)
	h := &Holiday{Month: time.December, Day: 25}
	paris, _ := time.LoadLocation("Europe/Paris")
	assert.Equal(time.Date(2020, 12, 25, 0, 0, 0, 0, paris), CalcDayOfMonth(h, 2020, paris))
	assert.Equal(time.Date(2021, 12, 25, 0, 0, 0, 0, paris), CalcDayOfMonth(h, 2021, paris))
}
