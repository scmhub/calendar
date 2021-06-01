# ![SCM logo](logo.png) Exchange Calendars

![Build Status](https://github.com/scmhub/calendar/workflows/Build%20and%20Test/badge.svg)

A Golang package of Exchange Calendars

Exchange calendars are defined by their [ISO-10383](https://www.iso20022.org/10383/iso-10383-market-identifier-codes) Market Identifier Code (MIC).

## Usage
```go
// Get an exchange calendar between a start year and and end year (both included)
nyse := XNYS(2010, 2025)

today := time.Now()

nyse.IsBusinessDay(today)
nyse.IsHoliday(today)

nyse.NextBusinessDay(today)
nyse.NextHoliday(today)


```
## Existing Calendar
| Market place   | Exchange                                                                                 | MIC      |
| -------------- | ---------------------------------------------------------------------------------------- | -------- |
| New York       | [New York Stock Exchange](https://www.nyse.com/index)                                    | XNYS     |
| New York       | [NASDAQ](https://www.nasdaq.com/)                                                        | XNAS     |
| Chicago        | [CBOE](http://markets.cboe.com)                                                          | XCBO     |
| Chicago        | [CBOE Futures Exchange](http://www.cfe.cboe.com)                                         | XCBF     |
| Toronto        | [Toronto Stock Exchange](https://www.tsx.com/)                                           | XTSE     |
| Mexico         | [Mexican Stock Exchange](https://www.bmv.com.mx)                                         | XMEX     |
| Sao Paulo      | [BMF Bovespa](http://www.b3.com.br/en_us/)                                               | BVMF     |
| London         | [London Stock Exchange](https://www.londonstockexchange.com)                             | XLON     |
| Amsterdam      | [Euronext Amsterdam](http://www.euronext.com)                                            | XAMS     |
| Brussels       | [Euronext Brussels](http://www.euronext.com)                                             | XBRU     |
| Lisbon         | [Euronext Lisbon](http://www.euronext.com)                                               | XLIS     |
| Paris          | [Euronext Paris](http://www.euronext.com)                                                | XPAR     |
| Milan          | [Euronext Milan - Borsa Italiana](http://www.borsaitaliana.it)                           | XMIL     |
| Madrid         | [Bolsa de Madrid](http://www.bolsamadrid.es)                                             | XMAD     |
| Franckfurt     | [Deutsche Boerse](http://www.deutsche-boerse.com)                                        | XFRA     |
| Franckfurt     | [Xetra](http://www.deutsche-boerse.com)                                                  | XETR     |
| Zurich         | [SIX Swiss Exchange](http://www.six-group.com/en/site/exchanges.html)                    | XSWX     |
| Mumbai         | [Bombay Stock Exchange](https://www.bseindia.com)                                        | XBOM     |
| Bangkok        | [Stock Exchange of Thailand](http://www.set.or.th/set/mainpage.do)                       | XBKK     |
| Singapore      | [Singapore Exchange](https://www.sgx.com)                                                | XSES     |
| Hong Kong      | [Hong Kong Stock Exchange](https://www.hkex.com.hk/index.html)                           | XHKG     |
| Shenzhen       | [Shenzhen Stock Exchange](http://www.szse.cn/English/index.html)                         | XSHE     |
| Shanghai       | [Shanghai Stock Exchange](http://www.sse.com.cn/sseportal/en/home/home.html)             | XSHG     |
| Seoul          | [Korea Exchange](http://eng.krx.co.kr)                                                   | XKRX     |
| Tokyo          | [Japan Exchange Group](https://www.jpx.co.jp/english/)                                   | XJPX     |
| Sidney         | [Austrialian Securities Exchange](https://www.asx.com.au/)                               | XASX     |
| South Korea    | [Korea Exchange](http://global.krx.co.kr)                                                | XKRX     |

<!---
| Chile          | [Santiago Stock Exchange](http://inter.bolsadesantiago.com/sitios/en/Paginas/home.aspx)  | XSGO     |
| Colombia       | [Colombia Securities Exchange](https://www.bvc.com.co/nueva/index_en.html)               | XBOG     |
| Peru           | [Lima Stock Exchange](https://www.bvl.com.pe)                                            | XLIM     |
| Iceland        | [Iceland Stock Exchange](http://www.nasdaqomxnordic.com/)                                | XICE     |
| Ireland        | [Irish Stock Exchange](http://www.ise.ie/)                                               | XDUB     |
| Denmark        | [Copenhagen Stock Exchange](http://www.nasdaqomxnordic.com/)                             | XCSE     |
| Finland        | [Helsinki Stock Exchange](http://www.nasdaqomxnordic.com/)                               | XHEL     |
| Sweden         | [Stockholm Stock Exchange](http://www.nasdaqomxnordic.com/)                              | XSTO     |
| Norway         | [Oslo Stock Exchange](https://www.oslobors.no/ob_eng/)                                   | XOSL     |
| Austria        | [Wiener Borse](https://www.wienerborse.at/en/)                                           | XWBO     |
| Czech Republic | [Prague Stock Exchange](https://www.pse.cz/en/)                                          | XPRA     |
| Hungary        | [Budapest Stock Exchange](https://bse.hu/)                                               | XBUD     |
| Poland         | [Poland Stock Exchange](http://www.gpw.pl)                                               | XWAR     |
| Greece         | [Athens Stock Exchange](http://www.helex.gr/)                                            | ASEX     |
| Turkey         | [Istanbul Stock Exchange](https://www.borsaistanbul.com/en/)                             | XIST     |
| Russia         | [Moscow Exchange](https://www.moex.com/en/)                                              | XMOS     |
| South Africa   | [Johannesburg Stock Exchange](https://www.jse.co.za/z)                                   | XJSE     |
| Malaysia       | [Malaysia Stock Exchange](http://www.bursamalaysia.com/market/)                          | XKLS     |
| Philippines    | [Philippine Stock Exchange](https://www.pse.com.ph/stockMarket/home.html)                | XPHS     |
| New Zealand    | [New Zealand Exchange](https://www.nzx.com/)                                             | XNZE     |
--->

## Creating Holidays & Calendars

```go
// Create Holidays
MemorialDay = &Holiday{
    Name:       "Memorial Day",
    Month:      time.May,
    Weekday:    time.Monday,
    NthWeekday: -1,
    calc:       CalcNthWeekday,
}
IndependenceDay = &Holiday{
    Name:       "Independence Day",
    Month:      time.July,
    Day:        4,
    observance: nearestWorkday,
    calc:       CalcDayOfMonth,
}
// Copy an Holiday and set observance
NewYear.Copy("New Year's Day").SetObservance(sundayToMonday)

// Create a Calendar
c := NewCalendar("New York Stock Exchange", NewYork)
// Set Sessions
c.SetEarlySession(&Session{7 * time.Hour, 9*time.Hour + 30*time.Minute})
c.SetCoreSessions(&Session{9*time.Hour + 30*time.Minute, 16 * time.Hour})
c.SetLateSession(&Session{16 * time.Hour, 20 * time.Hour})
// Add Holidays
c.AddHolidays(
    NewYear.Copy("New Year's Day").SetObservance(sundayToMonday),
    MLKDay,
    PresidentsDay,
    GoodFriday,
    MemorialDay,
    IndependenceDay,
    LaborDay,
    ThanksgivingDay,
    BlackFriday,
    ChristmasDay.Copy("Christmas Day").SetObservance(nearestWorkday),
)

```