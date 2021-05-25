package calendar

import (
	"fmt"
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

func TestEpiphany(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(2010, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2010, Paris))
	assert.Equal(time.Date(2011, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2011, Paris))
	assert.Equal(time.Date(2012, 1, 6, 0, 0, 0, 0, Paris), Epiphany.Calc(2012, Paris))
}

func TestMaundyThursday(t *testing.T) {
	assert := assert.New(t)
	fmt.Println(MaundyThursday.Calc(2012, Paris))
	assert.Equal(time.Date(2010, 3, 29, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2010, Paris))
	assert.Equal(time.Date(2011, 4, 18, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2011, Paris))
	assert.Equal(time.Date(2012, 4, 2, 0, 0, 0, 0, Paris), MaundyThursday.Calc(2012, Paris))
}
