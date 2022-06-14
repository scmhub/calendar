package main

import (
	"fmt"
	"time"

	"github.com/scmhub/calendar"
)

func main() {

	nyse := calendar.XLON(2010, 2025)

	now := time.Now()

	nyse.IsBusinessDay(now)
	nyse.IsHoliday(now)
	nyse.IsEarlyClose(now)

	nyse.IsOpen(now)

	nyse.SetYears(2021, 2023)

	fmt.Print(nyse)
}
