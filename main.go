package main

import (
	"date-calculation/date"
	"fmt"
	"os"
	"time"
)

func panicIfWeekends(d date.Day) date.Day {
	if d.DateRange.HasWeekends() {
		panic(d)
	}
	return d
}

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
		f.WriteString(fmt.Sprintf("New Years: %s\n", panicIfWeekends(date.USNewYears(year, loc)).Date()))

		f.WriteString(fmt.Sprintf("MLK Day: %s\n", panicIfWeekends(date.USMLKDay(year, loc)).Date()))

		f.WriteString(fmt.Sprintf("President’s Day: %s\n", panicIfWeekends(date.USPresidentsDay(year, loc)).Date()))

		f.WriteString(fmt.Sprintf("Memorial Day: %s\n", panicIfWeekends(date.USMemorialDay(year, loc)).Date()))

		f.WriteString(fmt.Sprintf("Juneteenth: %s\n", panicIfWeekends(date.USJuneteenth(year, loc)).Date()))

		f.WriteString(fmt.Sprintf("Independence Day: %s\n", panicIfWeekends(date.USIndependenceDay(year, loc)).Date()))

		f.WriteString(fmt.Sprintf("Labor Day: %s\n", panicIfWeekends(date.USLaborDay(year, loc)).Date()))

		f.WriteString(fmt.Sprintf("Indigenous People’s Day: %s\n", panicIfWeekends(date.USIndigenousPeoplesDay(year, loc)).Date()))

		f.WriteString(fmt.Sprintf("Veteran's Day: %s\n", panicIfWeekends(date.USVeteransDay(year, loc)).Date()))

		f.WriteString(fmt.Sprintf("Thanksgiving: %s\n", panicIfWeekends(date.USThanksgiving(year, loc)).Date()))

		f.WriteString(fmt.Sprintf("Christmas: %s\n", panicIfWeekends(date.USChristmas(year, loc)).Date()))

	}

	f.Close()

	// for canada
	f, err = os.Create("./CA-holidays")
	if err != nil {
		panic(err)
	}

	for year := 2022; year <= 2026; year++ {
		f.WriteString(fmt.Sprintf("New Year’s Day: %s\n", panicIfWeekends(date.CANewYears(year, loc)).Date()))

		EMon := date.CAEasterMondays[year]
		f.WriteString(fmt.Sprintf("Good Friday: %s\n", panicIfWeekends(EMon.AddDay(-3)).Date()))
		f.WriteString(fmt.Sprintf("Easter Monday: %s\n", panicIfWeekends(EMon).Date()))

		f.WriteString(fmt.Sprintf("Victoria Day: %s\n", panicIfWeekends(date.CAVictoriaDay(year, loc)).Date()))

		f.WriteString(fmt.Sprintf("Canada Day: %s\n", panicIfWeekends(date.CACanadaDay(year, loc)).Date()))

		f.WriteString(fmt.Sprintf("Civic Holiday: %s\n", panicIfWeekends(date.CACivicHoliday(year, loc)).Date()))

		f.WriteString(fmt.Sprintf("Labour Day: %s\n", panicIfWeekends(date.CALabourDay(year, loc)).Date()))

		f.WriteString(fmt.Sprintf("National Day for Truth and Reconciliation: %s\n", panicIfWeekends(date.CANationalDayforTruthandReconciliation(year, loc)).Date()))

		f.WriteString(fmt.Sprintf("Thanksgiving: %s\n", panicIfWeekends(date.CAThanksgiving(year, loc)).Date()))

		f.WriteString(fmt.Sprintf("Remembrance Day: %s\n", panicIfWeekends(date.CARemembranceDay(year, loc)).Date()))

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
