package calendar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// America

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
	assert.Equal("Frankfurt Stock Exchange", x.Name)
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
	assert.Equal(Bombay, x.Loc)
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
	x := XHKG(2010, 2025)
	assert.Equal("Stock Exchange of Hong Kong", x.Name)
	assert.Equal(HongKong, x.Loc)
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
