package date

import (
	"fmt"
	"time"
)

type DateRange struct {
	start, end time.Time
}

func (dr DateRange) String() string {
	return fmt.Sprintf("start from %s to %s", dr.start, dr.end)
}

type Day struct {
	DateRange
}

type GetDay interface {
	GetDay(nth int) Day
}

type Week struct {
	DateRange
}

type GetWeek interface {
	GetWeek(nth int) Week
}

type Month struct {
	DateRange
}

func (m Month) String() string {
	return m.DateRange.String()
}

type GetMonth interface {
	GetMonth(nth int) Month
}

type Year struct {
	DateRange
}

func (y Year) String() string {
	return y.DateRange.String()
}

func (y *Year) GetDay(nth int) {

}

// wait, what?
func (y *Year) GetWeek(nth int) Week {
	m := y.start.AddDate(0, 0, 7)
	return Week{
		DateRange: DateRange{
			start: m,
			end:   m.AddDate(0, 0, 7).Add(time.Nanosecond * -1),
		},
	}
}

func (y *Year) GetMonth(nth int) Month {
	m := y.start.AddDate(0, nth-1, 0)
	return Month{
		DateRange: DateRange{
			start: m,
			end:   m.AddDate(0, 1, 0).Add(time.Nanosecond * -1),
		},
	}
}

func (y *Month) GetDay(nth int) {

}

//:= TODO: timezone support
func NewYear(y int) Year {
	return Year{
		DateRange: DateRange{
			start: time.Date(y, 1, 1, 0, 0, 0, 0, time.UTC),
			end:   time.Date(y, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
	}
}
