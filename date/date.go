package date

import (
	"fmt"
	"time"
)

type GetDay interface {
	GetDay(nth int) Day
}

type GetWeekDay interface {
	GetWeekDay(nth time.Weekday) []Day
}

type GetMonth interface {
	GetMonth(nth time.Month) Month
}

// type GetWeek interface {
// 	GetWeek(nth int) Week
// }

///

type DateRange struct {
	start, end time.Time
}

func (dr DateRange) String() string {
	return fmt.Sprintf("start from %s to %s", dr.start.UTC(), dr.end.UTC())
}

type Day struct {
	DateRange
}

// type Week struct {
// 	DateRange
// }

type Year struct {
	DateRange
}

func (y Year) String() string {
	return y.DateRange.String()
}

func (y Year) GetDay(nth int) Day {
	d := y.start.AddDate(0, 0, nth-1)
	return Day{
		DateRange: DateRange{
			start: d,
			end:   d.AddDate(0, 0, 1).Add(time.Nanosecond * -1),
		},
	}
}

func (y Year) GetMonth(nth time.Month) Month {
	m := y.start.AddDate(0, int(nth)-1, 0)
	return Month{
		DateRange: DateRange{
			start: m,
			end:   m.AddDate(0, 1, 0).Add(time.Nanosecond * -1),
		},
	}
}

// wait, what?
// I dont think it is good idea for year have get week because
// the first day of this year is in the 0 week? or first week
// func (y *Year) GetWeek(nth int) Week {
// 	m := y.start.AddDate(0, 0, 7)
// 	return Week{
// 		DateRange: DateRange{
// 			start: m,
// 			end:   m.AddDate(0, 0, 7).Add(time.Nanosecond * -1),
// 		},
// 	}
// }

// same reason of getweek of year
// func (m *Month) GetWeek(nth int) {

// }

type Month struct {
	DateRange
}

func (m Month) String() string {
	return m.DateRange.String()
}

func (m Month) GetDay(nth int) Day {
	d := m.start.AddDate(0, 0, nth-1)
	return Day{
		DateRange: DateRange{
			start: d,
			end:   d.AddDate(0, 0, 1).Add(time.Nanosecond * -1),
		},
	}
}

func (m Month) GetWeekDay(nth time.Weekday) []Day {
	result := []Day{}
	var nextWeekdayStart time.Time
	if offset := m.start.Weekday() - nth; offset < 0 {
		// month start weekday before the weekday we want
		nextWeekdayStart = m.start.AddDate(0, 0, -int(offset))
	} else {
		nextWeekdayStart = m.start.AddDate(0, 0, 7-int(offset))
	}

	for {
		if nextWeekdayStart.Month() != m.start.Month() {
			return result
		}
		result = append(result, Day{
			DateRange: DateRange{
				start: nextWeekdayStart,
				end:   nextWeekdayStart.AddDate(0, 0, 1).Add(time.Nanosecond * -1),
			},
		})
		nextWeekdayStart = nextWeekdayStart.AddDate(0, 0, 7)
	}
}

func NewYear(y int, loc *time.Location) Year {
	return Year{
		DateRange: DateRange{
			start: time.Date(y, 1, 1, 0, 0, 0, 0, loc),
			end:   time.Date(y, 12, 31, 23, 59, 59, 999999999, loc),
		},
	}
}
