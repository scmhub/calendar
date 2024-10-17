# ![SCM logo](logo.png) Exchange Calendars

![Build Status](https://github.com/scmhub/calendar/workflows/Build%20and%20Test/badge.svg)
![Coverage](https://img.shields.io/badge/Coverage-97.5%25-brightgreen)
[![GoReportCard](https://goreportcard.com/badge/github.com/scmhub/calendar)](https://goreportcard.com/report/github.com/scmhub/calendar)
[![Go Reference](https://pkg.go.dev/badge/github.com/scmhub/calendar.svg)](https://pkg.go.dev/github.com/scmhub/calendar)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Overview

`calendar` is a Golang package that provides exchange-specific calendar functionality. It enables you to manage market holidays, business days, early closes, and sessions defined by Market Identifier Codes (MIC) for various financial exchanges globally.

## Features

- **Business Day Calculations**: Check if a date is a business day, holiday, or early close.
- **Custom Holidays**: Define and manage custom holidays (recurring or one-time events).
- **Exchange Sessions**: Handle exchange sessions with customizable open/close times.
- **Multi-Year Support**: Supports predefined calendars for financial exchanges over several years.
-  **Market Identifier Code**: Exchange calendars are defined by their [ISO-10383](https://www.iso20022.org/10383/iso-10383-market-identifier-codes) Market Identifier Code (MIC).

## Installation

Install the package by running:

```bash
go get github.com/scmhub/calendar
```

## Usage
#### Basic example
```go
import (
    "fmt"
    "time"
    "github.com/scmhub/calendar"
)

func main() {
    // Create a calendar with default years (current year ±5 years)
	nyse := calendar.XNYS()

    // Or create a calendar with a custom start and end year
	customCal := calendar.XNYS(2010, 2035) // From 2010 to 2035

    // Check if today is a business day
     today := time.Now()
    if nyse.IsBusinessDay(today) {
        fmt.Println("Today is a business day.")
    } else {
        fmt.Println("Today is not a business day.")
    }

    // Get the next business day
    nextBusinessDay := nyse.NextBusinessDay(today)
    fmt.Printf("Next business day: %v\n", nextBusinessDay)

	// Check if Black Friday 2030 is an early close
	blackFriday := time.Date(2030, time.November, 29, 0, 0, 0, 0, calendar.NewYork)
	fmt.Println(customCal.IsEarlyClose(blackFriday)) // Outputs: true
}
```

#### Print specific calendar
```go
import (
    "fmt"
    "github.com/scmhub/calendar"
)

func main() {
    // Load the New York Stock Exchange (XNYS) calendar
    nyse := calendar.XNYS()

    nyse.SetYears(2021,2022)

    fmt.Print(nyse)
}
```
Output
```
Calendar New York Stock Exchange:
        2021-Jan-01 Fri    New Year's Day
        2021-Jan-18 Mon    Martin Luther King Jr. Day
        2021-Feb-15 Mon    Presidents' Day
        2021-Apr-02 Fri    Good Friday
        2021-May-31 Mon    Memorial Day
        2021-Jul-05 Mon    Independence Day
        2021-Sep-06 Mon    Labor Day
        2021-Nov-25 Thu    Thanksgiving Day
        2021-Nov-26 Fri    Black Friday
        2021-Dec-24 Fri    Christmas Day
        2022-Jan-17 Mon    Martin Luther King Jr. Day
        2022-Feb-21 Mon    Presidents' Day
        2022-Apr-15 Fri    Good Friday
        2022-May-30 Mon    Memorial Day
        2022-Jun-20 Mon    Juneteenth National Independence Day
        2022-Jul-04 Mon    Independence Day
        2022-Sep-05 Mon    Labor Day
        2022-Nov-24 Thu    Thanksgiving Day
        2022-Nov-25 Fri    Black Friday
        2022-Dec-26 Mon    Christmas Day

```

#### Creating Holidays & Calendars

```go
// Create Recurring Holidays
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

// Create Non Recurring Holidays
// September 11 - september 11, 2001
SeptemberEleven = &Holiday{
    Name:   "Sepember 11",
    Month:  time.September,
    Day:    11,
    OnYear: 2001,
    calc:   CalcDayOfMonth,
}

// September 11 -14 range
SeptemberElevenDays = []*Holiday{
    SeptemberEleven,
    SeptemberEleven.Copy("Sepember 11 day 2").SetOffset(1),
    SeptemberEleven.Copy("Sepember 11 day 3").SetOffset(2),
    SeptemberEleven.Copy("Sepember 11 day 4").SetOffset(3),
}

// Copy an Holiday and set observance
NewYear.Copy("New Year's Day").SetObservance(sundayToMonday)

// Create a Calendar
c := NewCalendar("New York Stock Exchange", NewYork, 2010, 2025)
// Set Session
c.SetSession(&Session{
    EarlyOpen:  7 * time.Hour,
    Open:       9*time.Hour + 30*time.Minute,
    Close:      16 * time.Hour,
    EarlyClose: 13 * time.Hour,
    LateClose:  20 * time.Hour,
})
// Add Recurring Holidays
c.AddHolidays(
    NewYear.Copy().SetObservance(sundayToMonday),
    MLKDay,
    PresidentsDay,
    GoodFriday,
    MemorialDay,
    JuneteenthDay,
    IndependenceDay,
    LaborDay,
    ThanksgivingDay,
    ChristmasDay.Copy().SetObservance(nearestWorkday),
)
// Add Non Recurring Holidays
c.AddHolidays(USNationalDaysOfMourning...)
c.AddHolidays(SeptemberElevenDays...)
c.AddHolidays(HurricaneSandyDays...)
// Early Closing
c.AddEarlyClosingDays(
    BeforeIndependenceDay.Copy().SetObservance(onlyOnWeekdays(time.Monday, time.Tuesday, time.Thursday)),
    AfterIndependenceDay.Copy().SetBeforeYear(2013).SetObservance(onlyOnWeekdays(time.Friday)),
    BeforeIndependenceDay.Copy().SetAfterYear(2013).SetObservance(onlyOnWeekdays(time.Wednesday)),
    BlackFriday,
    ChristmasEve.Copy().SetObservance(exeptOnWeekdays(time.Friday)), // Overlap Christmas day observance if friday
)

```
## Existing Calendar

| Market place | Exchange                                                                     | MIC  |     |
| ------------ | ---------------------------------------------------------------------------- | ---- | --- |
| New York     | [New York Stock Exchange](https://www.nyse.com/index)                        | XNYS | ✅  |
| New York     | [NASDAQ](https://www.nasdaq.com/)                                            | XNAS | ✅  |
| Chicago      | [CBOE](http://markets.cboe.com)                                              | XCBO | ✅  |
| Chicago      | [CBOE Futures Exchange](http://www.cfe.cboe.com)                             | XCBF | ✅  |
| Toronto      | [Toronto Stock Exchange](https://www.tsx.com/)                               | XTSE |     |
| Mexico       | [Mexican Stock Exchange](https://www.bmv.com.mx)                             | XMEX |     |
| Sao Paulo    | [BMF Bovespa](http://www.b3.com.br/en_us/)                                   | BVMF |     |
| London       | [London Stock Exchange](https://www.londonstockexchange.com)                 | XLON | ✅  |
| Amsterdam    | [Euronext Amsterdam](http://www.euronext.com)                                | XAMS | ✅  |
| Brussels     | [Euronext Brussels](http://www.euronext.com)                                 | XBRU | ✅  |
| Lisbon       | [Euronext Lisbon](http://www.euronext.com)                                   | XLIS | ✅  |
| Paris        | [Euronext Paris](http://www.euronext.com)                                    | XPAR | ✅  |
| Milan        | [Euronext Milan - Borsa Italiana](http://www.borsaitaliana.it)               | XMIL | ✅  |
| Madrid       | [Bolsa de Madrid](http://www.bolsamadrid.es)                                 | XMAD | ✅  |
| Franckfurt   | [Deutsche Boerse](http://www.deutsche-boerse.com)                            | XFRA | ✅  |
| Franckfurt   | [Xetra](http://www.deutsche-boerse.com)                                      | XETR | ✅  |
| Zurich       | [SIX Swiss Exchange](http://www.six-group.com/en/site/exchanges.html)        | XSWX | ✅  |
| Mumbai       | [Bombay Stock Exchange](https://www.bseindia.com)                            | XBOM |     |
| Bangkok      | [Stock Exchange of Thailand](http://www.set.or.th/set/mainpage.do)           | XBKK |     |
| Singapore    | [Singapore Exchange](https://www.sgx.com)                                    | XSES | ✅  |
| Hong Kong    | [Hong Kong Stock Exchange](https://www.hkex.com.hk/index.html)               | XHKG | ✅  |
| Shenzhen     | [Shenzhen Stock Exchange](http://www.szse.cn/English/index.html)             | XSHE | ✅  |
| Shanghai     | [Shanghai Stock Exchange](http://www.sse.com.cn/sseportal/en/home/home.html) | XSHG | ✅  |
| Seoul        | [Korea Exchange](http://eng.krx.co.kr)                                       | XKRX |     |
| Tokyo        | [Japan Exchange Group](https://www.jpx.co.jp/english/)                       | XJPX | ✅  |
| Sidney       | [Austrialian Securities Exchange](https://www.asx.com.au/)                   | XASX |     |

<!---
| Chile          | [Santiago Stock Exchange](http://inter.bolsadesantiago.com/sitios/en/Paginas/home.aspx)  | XSGO |
| Colombia       | [Colombia Securities Exchange](https://www.bvc.com.co/nueva/index_en.html)               | XBOG |
| Peru           | [Lima Stock Exchange](https://www.bvl.com.pe)                                            | XLIM |
| Iceland        | [Iceland Stock Exchange](http://www.nasdaqomxnordic.com/)                                | XICE |
| Ireland        | [Irish Stock Exchange](http://www.ise.ie/)                                               | XDUB |
| Denmark        | [Copenhagen Stock Exchange](http://www.nasdaqomxnordic.com/)                             | XCSE |
| Finland        | [Helsinki Stock Exchange](http://www.nasdaqomxnordic.com/)                               | XHEL |
| Sweden         | [Stockholm Stock Exchange](http://www.nasdaqomxnordic.com/)                              | XSTO |
| Norway         | [Oslo Stock Exchange](https://www.oslobors.no/ob_eng/)                                   | XOSL |
| Austria        | [Wiener Borse](https://www.wienerborse.at/en/)                                           | XWBO |
| Czech Republic | [Prague Stock Exchange](https://www.pse.cz/en/)                                          | XPRA |
| Hungary        | [Budapest Stock Exchange](https://bse.hu/)                                               | XBUD |
| Poland         | [Poland Stock Exchange](http://www.gpw.pl)                                               | XWAR |
| Greece         | [Athens Stock Exchange](http://www.helex.gr/)                                            | ASEX |
| Turkey         | [Istanbul Stock Exchange](https://www.borsaistanbul.com/en/)                             | XIST |
| Russia         | [Moscow Exchange](https://www.moex.com/en/)                                              | XMOS |
| South Africa   | [Johannesburg Stock Exchange](https://www.jse.co.za/z)                                   | XJSE |
| Malaysia       | [Malaysia Stock Exchange](http://www.bursamalaysia.com/market/)                          | XKLS |
| Philippines    | [Philippine Stock Exchange](https://www.pse.com.ph/stockMarket/home.html)                | XPHS |
| New Zealand    | [New Zealand Exchange](https://www.nzx.com/)                                             | XNZE |
--->

## API References
Some key functions:
- **IsBusinessDay(t time.Time)**: Returns whether the date is a business day.
- **NextBusinessDay(t time.Time)**: Gets the next business day after the given date.
- **AddHolidays(h ...Holiday)**: Adds holidays to the calendar.
- **SetSession(session *Session)**: Configures open/close times for the exchange.

## Contributing
Contributions are welcome! Please submit issues or pull requests to improve the package. For significant changes, open an issue to discuss your ideas first.