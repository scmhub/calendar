package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// America
func TestUsEquities(t *testing.T) {
	assert := assert.New(t)
	x := usEquities("New York Stock Exchange", NewYork, 2000, 2025)
	assert.Equal("New York Stock Exchange", x.Name)
	assert.Equal(NewYork, x.Loc)
	assert.Equal(7*time.Hour, x.Session().EarlyOpen)
	assert.Equal(9*time.Hour+30*time.Minute, x.Session().Open)
	assert.Equal(time.Duration(0), x.Session().BreakStart)
	assert.Equal(time.Duration(0), x.Session().BreakStop)
	assert.Equal(16*time.Hour, x.Session().Close)
	assert.Equal(13*time.Hour, x.Session().EarlyClose)
	assert.Equal(20*time.Hour, x.Session().LateClose)
	// 2015
	assert.True(x.IsHoliday(time.Date(2015, 1, 1, 0, 0, 0, 0, NewYork)))      //New Year's day
	assert.True(x.IsHoliday(time.Date(2015, 1, 19, 0, 0, 0, 0, NewYork)))     // Martin Luther King Jr. day
	assert.True(x.IsHoliday(time.Date(2015, 2, 16, 0, 0, 0, 0, NewYork)))     // President's day
	assert.True(x.IsHoliday(time.Date(2015, 4, 3, 0, 0, 0, 0, NewYork)))      // Good Friday
	assert.True(x.IsHoliday(time.Date(2015, 5, 25, 0, 0, 0, 0, NewYork)))     // Memorial day
	assert.True(x.IsHoliday(time.Date(2015, 7, 3, 0, 0, 0, 0, NewYork)))      // Independence day
	assert.True(x.IsHoliday(time.Date(2015, 9, 7, 0, 0, 0, 0, NewYork)))      // Labor day
	assert.True(x.IsHoliday(time.Date(2015, 11, 26, 0, 0, 0, 0, NewYork)))    // Thanksgiving
	assert.True(x.IsEarlyClose(time.Date(2015, 11, 27, 0, 0, 0, 0, NewYork))) // Black Friday
	assert.True(x.IsEarlyClose(time.Date(2015, 12, 24, 0, 0, 0, 0, NewYork))) // Christmas Eve
	assert.True(x.IsHoliday(time.Date(2015, 12, 25, 0, 0, 0, 0, NewYork)))    // Christmas
	// 2016
	assert.True(x.IsHoliday(time.Date(2016, 1, 1, 0, 0, 0, 0, NewYork)))       //New Year's day
	assert.True(x.IsHoliday(time.Date(2016, 1, 18, 0, 0, 0, 0, NewYork)))      // Martin Luther King Jr. day
	assert.True(x.IsHoliday(time.Date(2016, 2, 15, 0, 0, 0, 0, NewYork)))      // President's day
	assert.True(x.IsHoliday(time.Date(2016, 3, 25, 0, 0, 0, 0, NewYork)))      // Good Friday
	assert.True(x.IsHoliday(time.Date(2016, 5, 30, 0, 0, 0, 0, NewYork)))      // Memorial day
	assert.True(x.IsHoliday(time.Date(2016, 7, 4, 0, 0, 0, 0, NewYork)))       // Independence day
	assert.True(x.IsHoliday(time.Date(2016, 9, 5, 0, 0, 0, 0, NewYork)))       // Labor day
	assert.True(x.IsHoliday(time.Date(2016, 11, 24, 0, 0, 0, 0, NewYork)))     // Thanksgiving
	assert.True(x.IsEarlyClose(time.Date(2016, 11, 25, 0, 0, 0, 0, NewYork)))  // Black Friday
	assert.False(x.IsEarlyClose(time.Date(2016, 12, 24, 0, 0, 0, 0, NewYork))) // Christmas Eve not observed is a sturday
	assert.True(x.IsHoliday(time.Date(2016, 12, 26, 0, 0, 0, 0, NewYork)))     // Christmas
	// 2017
	assert.True(x.IsHoliday(time.Date(2017, 1, 2, 0, 0, 0, 0, NewYork)))       //New Year's day
	assert.True(x.IsHoliday(time.Date(2017, 1, 16, 0, 0, 0, 0, NewYork)))      // Martin Luther King Jr. day
	assert.True(x.IsHoliday(time.Date(2017, 2, 20, 0, 0, 0, 0, NewYork)))      // President's day
	assert.True(x.IsHoliday(time.Date(2017, 4, 14, 0, 0, 0, 0, NewYork)))      // Good Friday
	assert.True(x.IsHoliday(time.Date(2017, 5, 29, 0, 0, 0, 0, NewYork)))      // Memorial day
	assert.True(x.IsEarlyClose(time.Date(2018, 7, 3, 0, 0, 0, 0, NewYork)))    // day before Independence day
	assert.True(x.IsHoliday(time.Date(2017, 7, 4, 0, 0, 0, 0, NewYork)))       // Independence day
	assert.True(x.IsHoliday(time.Date(2017, 9, 4, 0, 0, 0, 0, NewYork)))       // Labor day
	assert.True(x.IsHoliday(time.Date(2017, 11, 23, 0, 0, 0, 0, NewYork)))     // Thanksgiving
	assert.True(x.IsEarlyClose(time.Date(2017, 11, 24, 0, 0, 0, 0, NewYork)))  // Black Friday
	assert.False(x.IsEarlyClose(time.Date(2017, 12, 24, 0, 0, 0, 0, NewYork))) // Christmas Eve not observed is a sunday
	assert.True(x.IsHoliday(time.Date(2017, 12, 25, 0, 0, 0, 0, NewYork)))     // Christmas
	// 2018
	assert.True(x.IsHoliday(time.Date(2018, 1, 1, 0, 0, 0, 0, NewYork)))      //New Year's day
	assert.True(x.IsHoliday(time.Date(2018, 1, 15, 0, 0, 0, 0, NewYork)))     // Martin Luther King Jr. day
	assert.True(x.IsHoliday(time.Date(2018, 2, 19, 0, 0, 0, 0, NewYork)))     // President's day
	assert.True(x.IsHoliday(time.Date(2018, 3, 30, 0, 0, 0, 0, NewYork)))     // Good Friday
	assert.True(x.IsHoliday(time.Date(2018, 5, 28, 0, 0, 0, 0, NewYork)))     // Memorial day
	assert.True(x.IsEarlyClose(time.Date(2018, 7, 3, 0, 0, 0, 0, NewYork)))   // day before Independence day
	assert.True(x.IsHoliday(time.Date(2018, 7, 4, 0, 0, 0, 0, NewYork)))      // Independence day
	assert.False(x.IsEarlyClose(time.Date(2018, 7, 5, 0, 0, 0, 0, NewYork)))  // day after Independence day
	assert.True(x.IsHoliday(time.Date(2018, 9, 3, 0, 0, 0, 0, NewYork)))      // Labor day
	assert.True(x.IsHoliday(time.Date(2018, 11, 22, 0, 0, 0, 0, NewYork)))    // Thanksgiving
	assert.True(x.IsEarlyClose(time.Date(2018, 11, 23, 0, 0, 0, 0, NewYork))) // Black Friday
	assert.True(x.IsEarlyClose(time.Date(2018, 12, 24, 0, 0, 0, 0, NewYork))) // Christmas Eve not observed is a sunday
	assert.True(x.IsHoliday(time.Date(2018, 12, 25, 0, 0, 0, 0, NewYork)))    // Christmas
	// 2019
	assert.True(x.IsHoliday(time.Date(2019, 1, 1, 0, 0, 0, 0, NewYork)))      //New Year's day
	assert.True(x.IsHoliday(time.Date(2019, 1, 21, 0, 0, 0, 0, NewYork)))     // Martin Luther King Jr. day
	assert.True(x.IsHoliday(time.Date(2019, 2, 18, 0, 0, 0, 0, NewYork)))     // President's day
	assert.True(x.IsHoliday(time.Date(2019, 4, 19, 0, 0, 0, 0, NewYork)))     // Good Friday
	assert.True(x.IsHoliday(time.Date(2019, 5, 27, 0, 0, 0, 0, NewYork)))     // Memorial day
	assert.True(x.IsEarlyClose(time.Date(2019, 7, 3, 0, 0, 0, 0, NewYork)))   // day before Independence day
	assert.True(x.IsHoliday(time.Date(2019, 7, 4, 0, 0, 0, 0, NewYork)))      // Independence day
	assert.False(x.IsEarlyClose(time.Date(2019, 7, 5, 0, 0, 0, 0, NewYork)))  // day after Independence day
	assert.True(x.IsHoliday(time.Date(2019, 9, 2, 0, 0, 0, 0, NewYork)))      // Labor day
	assert.True(x.IsHoliday(time.Date(2019, 11, 28, 0, 0, 0, 0, NewYork)))    // Thanksgiving
	assert.True(x.IsEarlyClose(time.Date(2019, 11, 29, 0, 0, 0, 0, NewYork))) // Black Friday
	assert.True(x.IsEarlyClose(time.Date(2019, 12, 24, 0, 0, 0, 0, NewYork))) // Christmas Eve not observed is a sunday
	assert.True(x.IsHoliday(time.Date(2019, 12, 25, 0, 0, 0, 0, NewYork)))    // Christmas
	// 2020
	assert.True(x.IsHoliday(time.Date(2020, 1, 1, 0, 0, 0, 0, NewYork)))      //New Year's day
	assert.True(x.IsHoliday(time.Date(2020, 1, 20, 0, 0, 0, 0, NewYork)))     // Martin Luther King Jr. day
	assert.True(x.IsHoliday(time.Date(2020, 2, 17, 0, 0, 0, 0, NewYork)))     // President's day
	assert.True(x.IsHoliday(time.Date(2020, 4, 10, 0, 0, 0, 0, NewYork)))     // Good Friday
	assert.True(x.IsHoliday(time.Date(2020, 5, 25, 0, 0, 0, 0, NewYork)))     // Memorial day
	assert.False(x.IsEarlyClose(time.Date(2020, 7, 3, 0, 0, 0, 0, NewYork)))  // day before Independence day
	assert.True(x.IsHoliday(time.Date(2020, 7, 3, 0, 0, 0, 0, NewYork)))      // Independence day
	assert.False(x.IsEarlyClose(time.Date(2020, 7, 5, 0, 0, 0, 0, NewYork)))  // day after Independence day
	assert.True(x.IsHoliday(time.Date(2020, 9, 7, 0, 0, 0, 0, NewYork)))      // Labor day
	assert.True(x.IsHoliday(time.Date(2020, 11, 26, 0, 0, 0, 0, NewYork)))    // Thanksgiving
	assert.True(x.IsEarlyClose(time.Date(2020, 11, 27, 0, 0, 0, 0, NewYork))) // Black Friday
	assert.True(x.IsEarlyClose(time.Date(2020, 12, 24, 0, 0, 0, 0, NewYork))) // Christmas Eve not observed is a sunday
	assert.True(x.IsHoliday(time.Date(2020, 12, 25, 0, 0, 0, 0, NewYork)))    // Christmas
	// Independence day on Thursday
	assert.False(x.IsEarlyClose(time.Date(2002, 7, 3, 0, 0, 0, 0, NewYork))) // day before Independence day
	assert.True(x.IsHoliday(time.Date(2002, 7, 4, 0, 0, 0, 0, NewYork)))     // Independence day
	assert.True(x.IsEarlyClose(time.Date(2002, 7, 5, 0, 0, 0, 0, NewYork)))  // day after Independence day
	assert.True(x.IsEarlyClose(time.Date(2013, 7, 3, 0, 0, 0, 0, NewYork)))  // day before Independence day
	assert.True(x.IsHoliday(time.Date(2013, 7, 4, 0, 0, 0, 0, NewYork)))     // Independence day
	assert.False(x.IsEarlyClose(time.Date(2013, 7, 5, 0, 0, 0, 0, NewYork))) // day after Independence day
	// National Days of Mourning
	assert.True(x.IsHoliday(time.Date(2004, 6, 11, 0, 0, 0, 0, NewYork)))
	assert.True(x.IsHoliday(time.Date(2007, 1, 2, 0, 0, 0, 0, NewYork)))
	assert.True(x.IsHoliday(time.Date(2018, 11, 30, 0, 0, 0, 0, NewYork)))
	// September 11
	assert.True(x.IsHoliday(time.Date(2001, 9, 11, 0, 0, 0, 0, NewYork)))
	assert.True(x.IsHoliday(time.Date(2001, 9, 12, 0, 0, 0, 0, NewYork)))
	assert.True(x.IsHoliday(time.Date(2001, 9, 13, 0, 0, 0, 0, NewYork)))
	assert.True(x.IsHoliday(time.Date(2001, 9, 14, 0, 0, 0, 0, NewYork)))
	// Hurricane Sandy
	assert.True(x.IsHoliday(time.Date(2012, 10, 29, 0, 0, 0, 0, NewYork)))
	assert.True(x.IsHoliday(time.Date(2012, 10, 30, 0, 0, 0, 0, NewYork)))
}

func TestXNYS(t *testing.T) {
	assert := assert.New(t)
	x := XNYS(2010, 2025)
	assert.Equal("New York Stock Exchange", x.Name)
	assert.Equal(NewYork, x.Loc)
}
func TestXNAS(t *testing.T) {
	assert := assert.New(t)
	x := XNAS(2010, 2025)
	assert.Equal("NASDAQ", x.Name)
	assert.Equal(NewYork, x.Loc)
}
func TestXCBO(t *testing.T) {
	assert := assert.New(t)
	x := XCBO(2010, 2025)
	assert.Equal("Chicago Board Options Exchange", x.Name)
	assert.Equal(Chicago, x.Loc)
}
func TestXCBF(t *testing.T) {
	assert := assert.New(t)
	x := XCBF(2010, 2025)
	assert.Equal("Cboe Futures Exchange", x.Name)
	assert.Equal(Chicago, x.Loc)
}
func TestXTSE(t *testing.T) {
	assert := assert.New(t)
	x := XTSE(2010, 2025)
	assert.Equal("Toronto Stock Exchange", x.Name)
	assert.Equal(Chicago, x.Loc)
}
func TestXMEX(t *testing.T) {
	assert := assert.New(t)
	x := XMEX(2010, 2025)
	assert.Equal("Mexican Stock Exchange", x.Name)
	assert.Equal(Mexico, x.Loc)
}
func TestBVMF(t *testing.T) {
	assert := assert.New(t)
	x := BVMF(2010, 2025)
	assert.Equal("Brasilian Stock Exchange", x.Name)
	assert.Equal(SaoPaulo, x.Loc)
}

// Europe

func TestXLON(t *testing.T) {
	assert := assert.New(t)
	x := XLON(2010, 2025)
	assert.Equal("London Stock Exchange", x.Name)
	assert.Equal(London, x.Loc)
}
func TestEuronext(t *testing.T) {
	assert := assert.New(t)
	x := euronext("Euronext Paris", Paris, 2015, 2025)
	assert.Equal("Euronext Paris", x.Name)
	assert.Equal(Paris, x.Loc)
	assert.Equal(time.Duration(0), x.Session().EarlyOpen)
	assert.Equal(9*time.Hour, x.Session().Open)
	assert.Equal(time.Duration(0), x.Session().BreakStart)
	assert.Equal(time.Duration(0), x.Session().BreakStop)
	assert.Equal(17*time.Hour+30*time.Minute, x.Session().Close)
	assert.Equal(14*time.Hour+5*time.Minute, x.Session().EarlyClose)
	assert.Equal(time.Duration(0), x.Session().LateClose)
	// 2015
	assert.True(x.IsHoliday(time.Date(2015, 1, 1, 0, 0, 0, 0, Paris)))      //New Year's day
	assert.True(x.IsHoliday(time.Date(2015, 4, 3, 0, 0, 0, 0, Paris)))      // Good Friday
	assert.True(x.IsHoliday(time.Date(2015, 4, 6, 0, 0, 0, 0, Paris)))      // Easter Monday
	assert.True(x.IsHoliday(time.Date(2015, 5, 1, 0, 0, 0, 0, Paris)))      // Worker's day
	assert.True(x.IsEarlyClose(time.Date(2015, 12, 24, 0, 0, 0, 0, Paris))) // Christmas Eve
	assert.True(x.IsHoliday(time.Date(2015, 12, 25, 0, 0, 0, 0, Paris)))    // Christmas
	assert.False(x.IsHoliday(time.Date(2015, 12, 26, 0, 0, 0, 0, Paris)))   // Boxing day
	assert.True(x.IsEarlyClose(time.Date(2015, 12, 31, 0, 0, 0, 0, Paris))) // New year's eve
	// 2016
	assert.True(x.IsHoliday(time.Date(2016, 1, 1, 0, 0, 0, 0, Paris)))       //New Year's day
	assert.True(x.IsHoliday(time.Date(2016, 3, 25, 0, 0, 0, 0, Paris)))      // Good Friday
	assert.True(x.IsHoliday(time.Date(2016, 3, 28, 0, 0, 0, 0, Paris)))      // Easter Monday
	assert.False(x.IsHoliday(time.Date(2016, 5, 1, 0, 0, 0, 0, Paris)))      // Worker's day
	assert.False(x.IsEarlyClose(time.Date(2016, 12, 24, 0, 0, 0, 0, Paris))) // Christmas Eve
	assert.True(x.IsHoliday(time.Date(2016, 12, 26, 0, 0, 0, 0, Paris)))     // Christmas
	assert.True(x.IsHoliday(time.Date(2016, 12, 26, 0, 0, 0, 0, Paris)))     // Boxing day
	assert.False(x.IsEarlyClose(time.Date(2016, 12, 31, 0, 0, 0, 0, Paris))) // New year's eve
	// 2017
	assert.False(x.IsHoliday(time.Date(2017, 1, 1, 0, 0, 0, 0, Paris)))      //New Year's day
	assert.True(x.IsHoliday(time.Date(2017, 4, 14, 0, 0, 0, 0, Paris)))      // Good Friday
	assert.True(x.IsHoliday(time.Date(2017, 4, 17, 0, 0, 0, 0, Paris)))      // Easter Monday
	assert.True(x.IsHoliday(time.Date(2017, 5, 1, 0, 0, 0, 0, Paris)))       // Worker's day
	assert.False(x.IsEarlyClose(time.Date(2017, 12, 24, 0, 0, 0, 0, Paris))) // Christmas Eve
	assert.True(x.IsHoliday(time.Date(2017, 12, 25, 0, 0, 0, 0, Paris)))     // Christmas
	assert.True(x.IsHoliday(time.Date(2017, 12, 26, 0, 0, 0, 0, Paris)))     // Boxing day
	assert.False(x.IsEarlyClose(time.Date(2017, 12, 31, 0, 0, 0, 0, Paris))) // New year's eve
	// 2018
	assert.True(x.IsHoliday(time.Date(2018, 1, 1, 0, 0, 0, 0, Paris)))      //New Year's day
	assert.True(x.IsHoliday(time.Date(2018, 3, 30, 0, 0, 0, 0, Paris)))     // Good Friday
	assert.True(x.IsHoliday(time.Date(2018, 4, 2, 0, 0, 0, 0, Paris)))      // Easter Monday
	assert.True(x.IsHoliday(time.Date(2018, 5, 1, 0, 0, 0, 0, Paris)))      // Worker's day
	assert.True(x.IsEarlyClose(time.Date(2018, 12, 24, 0, 0, 0, 0, Paris))) // Christmas Eve
	assert.True(x.IsHoliday(time.Date(2018, 12, 25, 0, 0, 0, 0, Paris)))    // Christmas
	assert.True(x.IsHoliday(time.Date(2018, 12, 26, 0, 0, 0, 0, Paris)))    // Boxing day
	assert.True(x.IsEarlyClose(time.Date(2018, 12, 31, 0, 0, 0, 0, Paris))) // New year's eve
	// 2019
	assert.True(x.IsHoliday(time.Date(2019, 1, 1, 0, 0, 0, 0, Paris)))      //New Year's day
	assert.True(x.IsHoliday(time.Date(2019, 4, 19, 0, 0, 0, 0, Paris)))     // Good Friday
	assert.True(x.IsHoliday(time.Date(2019, 4, 22, 0, 0, 0, 0, Paris)))     // Easter Monday
	assert.True(x.IsHoliday(time.Date(2019, 5, 1, 0, 0, 0, 0, Paris)))      // Worker's day
	assert.True(x.IsEarlyClose(time.Date(2019, 12, 24, 0, 0, 0, 0, Paris))) // Christmas Eve
	assert.True(x.IsHoliday(time.Date(2019, 12, 25, 0, 0, 0, 0, Paris)))    // Christmas
	assert.True(x.IsHoliday(time.Date(2019, 12, 26, 0, 0, 0, 0, Paris)))    // Boxing day
	assert.True(x.IsEarlyClose(time.Date(2019, 12, 31, 0, 0, 0, 0, Paris))) // New year's eve
	// 2020
	assert.True(x.IsHoliday(time.Date(2020, 1, 1, 0, 0, 0, 0, Paris)))      //New Year's day
	assert.True(x.IsHoliday(time.Date(2020, 4, 10, 0, 0, 0, 0, Paris)))     // Good Friday
	assert.True(x.IsHoliday(time.Date(2020, 4, 13, 0, 0, 0, 0, Paris)))     // Easter Monday
	assert.True(x.IsHoliday(time.Date(2020, 5, 1, 0, 0, 0, 0, Paris)))      // Worker's day
	assert.True(x.IsEarlyClose(time.Date(2020, 12, 24, 0, 0, 0, 0, Paris))) // Christmas Eve
	assert.True(x.IsHoliday(time.Date(2020, 12, 25, 0, 0, 0, 0, Paris)))    // Christmas
	assert.False(x.IsHoliday(time.Date(2020, 12, 26, 0, 0, 0, 0, Paris)))   // Boxing day
	assert.True(x.IsEarlyClose(time.Date(2020, 12, 31, 0, 0, 0, 0, Paris))) // New year's eve

}

func TestXAMS(t *testing.T) {
	assert := assert.New(t)
	x := XAMS(2010, 2025)
	assert.Equal("Euronext Amsterdam", x.Name)
	assert.Equal(Amsterdam, x.Loc)
}
func TestXBRU(t *testing.T) {
	assert := assert.New(t)
	x := XBRU(2010, 2025)
	assert.Equal("Euronext Brussels", x.Name)
	assert.Equal(Brussels, x.Loc)
}
func TestXLIS(t *testing.T) {
	assert := assert.New(t)
	x := XLIS(2010, 2025)
	assert.Equal("Euronext Lisbon", x.Name)
	assert.Equal(Lisbon, x.Loc)
}
func TestXPAR(t *testing.T) {
	assert := assert.New(t)
	x := XPAR(2010, 2025)
	assert.Equal("Euronext Paris", x.Name)
	assert.Equal(Paris, x.Loc)
}
func TestXMIL(t *testing.T) {
	assert := assert.New(t)
	x := XMIL(2010, 2025)
	assert.Equal("Euronext Milan", x.Name)
	assert.Equal(Milan, x.Loc)
}
func TestXMAD(t *testing.T) {
	assert := assert.New(t)
	x := XMAD(2010, 2025)
	assert.Equal("Madrid Stock Exchange", x.Name)
	assert.Equal(Madrid, x.Loc)
}
func TestXFRA(t *testing.T) {
	assert := assert.New(t)
	x := XFRA(2010, 2025)
	assert.Equal("Deutsche Boerse", x.Name)
	assert.Equal(Franckfurt, x.Loc)
}
func TestXETR(t *testing.T) {
	assert := assert.New(t)
	x := XETR(2010, 2025)
	assert.Equal("Deutsche Borse Xetra", x.Name)
	assert.Equal(Franckfurt, x.Loc)
}
func TestXSWX(t *testing.T) {
	assert := assert.New(t)
	x := XSWX(2010, 2025)
	assert.Equal("SIX Group", x.Name)
	assert.Equal(Zurich, x.Loc)
}

// Asia/Pacific

func TestXBOM(t *testing.T) {
	assert := assert.New(t)
	x := XBOM(2010, 2025)
	assert.Equal("Bombay Stock Exchange", x.Name)
	assert.Equal(Mumbai, x.Loc)
}
func TestXBKK(t *testing.T) {
	assert := assert.New(t)
	x := XBKK(2010, 2025)
	assert.Equal("Stock Exchange of Thailand", x.Name)
	assert.Equal(Bangkok, x.Loc)
}
func TestXSES(t *testing.T) {
	assert := assert.New(t)
	x := XSES(2010, 2025)
	assert.Equal("Singapore Exchange", x.Name)
	assert.Equal(Singapore, x.Loc)
}
func TestXHKG(t *testing.T) {
	assert := assert.New(t)
	x := XHKG(2017, 2021)
	assert.Equal("Stock Exchange of Hong Kong", x.Name)
	assert.Equal(HongKong, x.Loc)
	assert.Equal(9*time.Hour, x.Session().EarlyOpen)
	assert.Equal(9*time.Hour+30*time.Minute, x.Session().Open)
	assert.Equal(12*time.Hour, x.Session().BreakStart)
	assert.Equal(13*time.Hour, x.Session().BreakStop)
	assert.Equal(16*time.Hour, x.Session().Close)
	assert.Equal(12*time.Hour, x.Session().EarlyClose)
	assert.Equal(time.Duration(0), x.Session().LateClose)

}
func TestXSHE(t *testing.T) {
	assert := assert.New(t)
	x := XSHE(2010, 2025)
	assert.Equal("Shenzhen Stock Exchange", x.Name)
	assert.Equal(Shenzhen, x.Loc)
}
func TestXSHG(t *testing.T) {
	assert := assert.New(t)
	x := XSHG(2010, 2025)
	assert.Equal("Shanghai Stock Exchange", x.Name)
	assert.Equal(Shanghai, x.Loc)
}
func TestXKRX(t *testing.T) {
	assert := assert.New(t)
	x := XKRX(2010, 2025)
	assert.Equal("Korea Exchange", x.Name)
	assert.Equal(Seoul, x.Loc)
}
func TestXJPX(t *testing.T) {
	assert := assert.New(t)
	x := XJPX(2010, 2025)
	assert.Equal("Japan Exchange Group", x.Name)
	assert.Equal(Tokyo, x.Loc)
}
func TestXASX(t *testing.T) {
	assert := assert.New(t)
	x := XASX(2010, 2025)
	assert.Equal("Australian Securities Exchange", x.Name)
	assert.Equal(Sydney, x.Loc)
}
