package calendar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistry(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(GetCalendar("xabc"))
	assert.NotNil(GetCalendar("xnys"))
	assert.Equal(XNYS().Name, GetCalendar("xnys").Name)
	assert.Equal(XNYS().Loc, GetCalendar("xnys").Loc)
	assert.Equal(XNYS().startYear, GetCalendar("xnys").startYear)
	assert.Equal(XNYS().endYear, GetCalendar("xnys").endYear)
	assert.Equal(XNYS().session, GetCalendar("xnys").session)
}
