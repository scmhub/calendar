package calendar

var registry map[string]func(years ...int) *Calendar

func register(name string, f func(years ...int) *Calendar) {
	registry[name] = f
}

func GetCalendar(name string, years ...int) *Calendar {
	f, ok := registry[name]
	if !ok {
		return nil
	}
	return f(years...)
}

func init() {
	registry = make(map[string]func(years ...int) *Calendar)
	// America
	register("xnys", XNYS) // New York Stock Exchange
	register("xnas", XNAS) // NASDAQ
	register("xcbo", XCBO) // Chicago Board Options Exchange
	register("xcbf", XCBF) // Cboe Futures Exchange
	register("xtse", XTSE) // Toronto Stock Exchange
	register("xmex", XMEX) // Mexican Stock Exchange
	register("bvmf", BVMF) // Brasilian Stock Exchange - Bolsa de Valores, Mercados e Futuros
	// Europe
	register("xlon", XLON) // London Stock Exchange
	register("xams", XAMS) // Euronext Amsterdam
	register("xbru", XBRU) // Euronext Brussels
	register("xlis", XLIS) // Euronext Lisbon
	register("xpar", XPAR) // Euronext Paris
	register("xmil", XMIL) // Euronext Milan - Borsa Italiana S.P.A
	register("xmad", XMAD) // Madrid Stock Exchange
	register("xfra", XFRA) // Frankfurt Stock Exchange
	register("xetr", XETR) // Deutsche Borse Xetra
	register("xswx", XSWX) // SIX Group (SWX Swiss Exchange)
	// Asia/Pacific
	register("xbom", XBOM) // Bombay Stock Exchange
	register("xbkk", XBKK) // Stock Exchange of Thailand
	register("xses", XSES) // Singapore Exchange (SGX)
	register("xhkg", XHKG) // Stock Exchange of Hong Kong
	register("xshe", XSHE) // Shenzhen Stock Exchange
	register("xshg", XSHG) // Shanghai Stock Exchange
	register("xkrx", XKRX) // Korea Exchange
	register("xjpx", XJPX) // Japan Exchange Group
	register("xasx", XASX) // Australian Securities Exchange
}
