package calendar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// America

func TestNYSE(t *testing.T) {
	assert := assert.New(t)
	x := NYSE(2010, 2025)
	assert.Equal("New York Stock Exchange", x.Name)
	assert.Equal(NewYork, x.Loc)
}

func TestNASDAQ(t *testing.T) {
	assert := assert.New(t)
	x := NASDAQ(2010, 2025)
	assert.Equal("NASDAQ", x.Name)
	assert.Equal(NewYork, x.Loc)
}

func TestCBOE(t *testing.T) {
	assert := assert.New(t)
	x := CBOE(2010, 2025)
	assert.Equal("Chicago Board Options Exchange", x.Name)
	assert.Equal(Chicago, x.Loc)
}

func TestCFE(t *testing.T) {
	assert := assert.New(t)
	x := CFE(2010, 2025)
	assert.Equal("Cboe Futures Exchange", x.Name)
	assert.Equal(Chicago, x.Loc)
}

// Europe

func TestLSE(t *testing.T) {
	assert := assert.New(t)
	x := LSE(2010, 2025)
	assert.Equal("London Stock Exchange", x.Name)
	assert.Equal(London, x.Loc)
}

func TestEuronextAmsterdam(t *testing.T) {
	assert := assert.New(t)
	x := EuronextAmsterdam(2010, 2025)
	assert.Equal("Euronext Amsterdam", x.Name)
	assert.Equal(Amsterdam, x.Loc)
}

func TestEuronextBrussels(t *testing.T) {
	assert := assert.New(t)
	x := EuronextBrussels(2010, 2025)
	assert.Equal("Euronext Brussels", x.Name)
	assert.Equal(Brussels, x.Loc)
}
func TestEuronextLisbon(t *testing.T) {
	assert := assert.New(t)
	x := EuronextLisbon(2010, 2025)
	assert.Equal("Euronext Lisbon", x.Name)
	assert.Equal(Lisbon, x.Loc)
}
func TestEuronextParis(t *testing.T) {
	assert := assert.New(t)
	x := EuronextParis(2010, 2025)
	assert.Equal("Euronext Paris", x.Name)
	assert.Equal(Paris, x.Loc)
}

// Asia

func TestHKE(t *testing.T) {
	assert := assert.New(t)
	x := HKE(2010, 2025)
	assert.Equal("Hong Kong Exchanges", x.Name)
	assert.Equal(HongKong, x.Loc)
}
func TestShenzhenSE(t *testing.T) {
	assert := assert.New(t)
	x := ShenzhenSE(2010, 2025)
	assert.Equal("Shenzhen Stock Exchange", x.Name)
	assert.Equal(Shenzhen, x.Loc)
}
func TestShanghaiSE(t *testing.T) {
	assert := assert.New(t)
	x := ShanghaiSE(2010, 2025)
	assert.Equal("Shanghai Stock Exchange", x.Name)
	assert.Equal(Shanghai, x.Loc)
}
func TestKRX(t *testing.T) {
	assert := assert.New(t)
	x := KRX(2010, 2025)
	assert.Equal("Korea Exchange", x.Name)
	assert.Equal(Seoul, x.Loc)
}
func TestJEG(t *testing.T) {
	assert := assert.New(t)
	x := JEG(2010, 2025)
	assert.Equal("Japan Exchange Group", x.Name)
	assert.Equal(Tokyo, x.Loc)
}
func TestASE(t *testing.T) {
	assert := assert.New(t)
	x := ASE(2010, 2025)
	assert.Equal("Australian Securities Exchange", x.Name)
	assert.Equal(Sydney, x.Loc)
}
