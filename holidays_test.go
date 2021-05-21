package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewYear(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 1, 1, 0, 0, 0, 0, Paris), NewYear.Calc(2010, Paris)) // saturday
	assert.Equal(time.Date(2011, 1, 1, 0, 0, 0, 0, Paris), NewYear.Calc(2011, Paris)) // sunday
	assert.Equal(time.Date(2012, 1, 1, 0, 0, 0, 0, Paris), NewYear.Calc(2012, Paris)) // monday
}
