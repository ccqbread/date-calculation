package main

import (
	"date-calculation/date"
	"fmt"
	"os"
	"time"
)

func main() {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}

	// make next five years US holidays
	f, err := os.Create("./US-holidays")
	if err != nil {
		panic(err)
	}

	for year := 2022; year <= 2026; year++ {
		f.WriteString(fmt.Sprintf("New Years: %s\n", date.USNewYears(year, loc).Date()))

		f.WriteString(fmt.Sprintf("MLK Day: %s\n", date.USMLKDay(year, loc).Date()))

		f.WriteString(fmt.Sprintf("President’s Day: %s\n", date.USPresidentsDay(year, loc).Date()))

		f.WriteString(fmt.Sprintf("Memorial Day: %s\n", date.USMemorialDay(year, loc).Date()))

		f.WriteString(fmt.Sprintf("Juneteenth: %s\n", date.USJuneteenth(year, loc).Date()))

		f.WriteString(fmt.Sprintf("Independence Day: %s\n", date.USIndependenceDay(year, loc).Date()))

		f.WriteString(fmt.Sprintf("Labor Day: %s\n", date.USLaborDay(year, loc).Date()))

		f.WriteString(fmt.Sprintf("Indigenous People’s Day: %s\n", date.USIndigenousPeoplesDay(year, loc).Date()))

		f.WriteString(fmt.Sprintf("Veteran's Day: %s\n", date.USVeteransDay(year, loc).Date()))

		f.WriteString(fmt.Sprintf("Thanksgiving: %s\n", date.USThanksgiving(year, loc).Date()))

		f.WriteString(fmt.Sprintf("Christmas: %s\n", date.USChristmas(year, loc).Date()))

	}

	f.Close()

	// for canada
	f, err = os.Create("./CA-holidays")
	if err != nil {
		panic(err)
	}

	for year := 2022; year <= 2026; year++ {
		f.WriteString(fmt.Sprintf("New Year’s Day: %s\n", date.CANewYears(year, loc).Date()))

		EMon := date.CAEasterMondays[year]
		f.WriteString(fmt.Sprintf("Good Friday: %s\n", EMon.AddDay(-3).Date()))
		f.WriteString(fmt.Sprintf("Easter Monday: %s\n", EMon.Date()))

		f.WriteString(fmt.Sprintf("Victoria Day: %s\n", date.CAVictoriaDay(year, loc).Date()))

		f.WriteString(fmt.Sprintf("Canada Day: %s\n", date.CACanadaDay(year, loc).Date()))

		f.WriteString(fmt.Sprintf("Civic Holiday: %s\n", date.CACivicHoliday(year, loc).Date()))

		f.WriteString(fmt.Sprintf("Labour Day: %s\n", date.CALabourDay(year, loc).Date()))

		f.WriteString(fmt.Sprintf("National Day for Truth and Reconciliation: %s\n", date.CANationalDayforTruthandReconciliation(year, loc).Date()))

		f.WriteString(fmt.Sprintf("Thanksgiving: %s\n", date.CAThanksgiving(year, loc).Date()))

		f.WriteString(fmt.Sprintf("Remembrance Day: %s\n", date.CARemembranceDay(year, loc).Date()))

		dd := []date.Day{
			date.CAChristmasDay(year, loc),
			date.CABoxingDay(year, loc),
		}

		if dd[0].Date() == dd[1].Date() {
			dd = date.CAConflictAfterShift(dd)
		}

		f.WriteString(fmt.Sprintf("Christmas Day: %s\n", dd[0].Date()))
		f.WriteString(fmt.Sprintf("Boxing Day: %s\n", dd[1].Date()))

	}

	f.Close()
}
