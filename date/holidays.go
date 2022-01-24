package date

import (
	"sort"
	"time"
)

type Holiday func(year int, loc *time.Location) Day

// US holiday
var (
	USMLKDay Holiday = func(y int, loc *time.Location) Day {
		return USWeekendShift(NewYear(y, loc).GetMonth(time.January).GetWeekDay(time.Monday)[2])
	}

	USPresidentsDay Holiday = func(y int, loc *time.Location) Day {
		return USWeekendShift(NewYear(y, loc).GetMonth(time.February).GetWeekDay(time.Monday)[2])
	}

	USMemorialDay Holiday = func(y int, loc *time.Location) Day {
		allMon := NewYear(y, loc).GetMonth(time.May).GetWeekDay(time.Monday)
		return USWeekendShift(allMon[len(allMon)-1])
	}

	USJuneteenth Holiday = func(y int, loc *time.Location) Day {
		return USWeekendShift(NewYear(y, loc).GetMonth(time.June).GetDay(19))
	}

	USIndependenceDay Holiday = func(y int, loc *time.Location) Day {
		return USWeekendShift(NewYear(y, loc).GetMonth(time.July).GetDay(4))
	}

	USLaborDay Holiday = func(y int, loc *time.Location) Day {
		return USWeekendShift(NewYear(y, loc).GetMonth(time.September).GetWeekDay(time.Monday)[0])
	}

	USIndigenousPeoplesDay Holiday = func(y int, loc *time.Location) Day {
		return USWeekendShift(NewYear(y, loc).GetMonth(time.October).GetWeekDay(time.Monday)[1])
	}

	USVeteransDay Holiday = func(y int, loc *time.Location) Day {
		return USWeekendShift(NewYear(y, loc).GetMonth(time.November).GetDay(11))
	}

	USThanksgiving Holiday = func(y int, loc *time.Location) Day {
		return USWeekendShift(NewYear(y, loc).GetMonth(time.November).GetWeekDay(time.Thursday)[3])
	}

	USChristmas Holiday = func(y int, loc *time.Location) Day {
		return USWeekendShift(NewYear(y, loc).GetMonth(time.December).GetDay(25))
	}

	USNewYears Holiday = func(y int, loc *time.Location) Day {
		return USWeekendShift(NewYear(y, loc).GetDay(1))
	}
)

func ifSunday(d Day) Day {
	if d.start.Weekday() == time.Sunday {
		return d.AddDay(1)
	}
	return d
}

func ifSaturday(d Day) Day {
	if d.start.Weekday() == time.Saturday {
		return d.AddDay(-1)
	}
	return d
}

func USWeekendShift(d Day) Day {
	return ifSunday(ifSaturday(d))
}

// canada holiday
var (
	CANewYears Holiday = func(y int, loc *time.Location) Day {
		return CAWeekendShift(NewYear(y, loc).GetDay(1))
	}

	//goodfriday
	//eastermonday

	CAVictoriaDay Holiday = func(y int, loc *time.Location) Day {
		var result Day
		for _, dd := range NewYear(y, loc).GetMonth(time.May).GetWeekDay(time.Monday) {
			if dd.start.Day() < 25 {
				result = dd
			} else {
				break
			}
		}
		return result
	}

	CACanadaDay Holiday = func(y int, loc *time.Location) Day {
		return CAWeekendShift(NewYear(y, loc).GetMonth(time.July).GetDay(1))
	}

	CACivicHoliday Holiday = func(y int, loc *time.Location) Day {
		return NewYear(y, loc).GetMonth(time.August).GetWeekDay(time.Monday)[0]
	}

	CALabourDay Holiday = func(y int, loc *time.Location) Day {
		return NewYear(y, loc).GetMonth(time.September).GetWeekDay(time.Monday)[0]
	}

	CANationalDayforTruthandReconciliation Holiday = func(y int, loc *time.Location) Day {
		return CAWeekendShift(NewYear(y, loc).GetMonth(time.September).GetDay(30))
	}

	CAThanksgiving Holiday = func(y int, loc *time.Location) Day {
		return NewYear(y, loc).GetMonth(time.October).GetWeekDay(time.Monday)[1]
	}

	CARemembranceDay Holiday = func(y int, loc *time.Location) Day {
		return CAWeekendShift(NewYear(y, loc).GetMonth(time.November).GetDay(11))
	}

	CAChristmasDay Holiday = func(y int, loc *time.Location) Day {
		return CAWeekendShift(NewYear(y, loc).GetMonth(time.December).GetDay(25))
	}

	CABoxingDay Holiday = func(y int, loc *time.Location) Day {
		return CAWeekendShift(NewYear(y, loc).GetMonth(time.December).GetDay(26))
	}
)

func CAWeekendShift(d Day) Day {
	if d.start.Weekday() == time.Saturday {
		return d.AddDay(2)
	}
	if d.start.Weekday() == time.Sunday {
		return d.AddDay(1)
	}
	return d
}

// make sure the conflict happen
func CAConflictAfterShift(d []Day) []Day {
	sort.Slice(d, func(i, j int) bool {
		return d[i].OffsetFromOriginal > d[j].OffsetFromOriginal
	})

	result := []Day{d[0]}
	for offset, dd := range d[1:] {
		result = append(result, dd.AddDay(offset+1))
	}

	return result
}

var CAEasterMondays = map[int]Day{
	2022: Day{
		DateRange: DateRange{
			start: time.Date(2022, 4, 18, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2022, 4, 18, 23, 59, 59, 999999999, time.UTC),
		},
		OffsetFromOriginal: 0,
	},

	2023: Day{
		DateRange: DateRange{
			start: time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2023, 4, 10, 23, 59, 59, 999999999, time.UTC),
		},
		OffsetFromOriginal: 0,
	},

	2024: Day{
		DateRange: DateRange{
			start: time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2024, 4, 1, 23, 59, 59, 999999999, time.UTC),
		},
		OffsetFromOriginal: 0,
	},

	2025: Day{
		DateRange: DateRange{
			start: time.Date(2025, 4, 21, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2025, 4, 21, 23, 59, 59, 999999999, time.UTC),
		},
		OffsetFromOriginal: 0,
	},

	2026: Day{
		DateRange: DateRange{
			start: time.Date(2026, 4, 6, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2026, 4, 6, 23, 59, 59, 999999999, time.UTC),
		},
		OffsetFromOriginal: 0,
	},
}
