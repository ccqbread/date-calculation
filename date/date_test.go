package date

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadLocation(t *testing.T) {
	_, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Error(err)
	}
}

func TestGetWeekDay(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Error(err)
	}

	testCases := []struct {
		name         string
		year         Year
		weekdayGet   func(y Year) Day
		weekdayCheck func(d Day) error
	}{
		{
			name: "MLK day of 2022",
			year: NewYear(2022, loc),
			weekdayGet: func(y Year) Day {
				return y.GetMonth(time.January).GetWeekDay(time.Monday)[2]
			},
			weekdayCheck: func(d Day) error {
				if d.start.Month() != time.January || d.start.Weekday() != time.Monday {
					return errors.New("Check failed")
				}
				return nil
			},
		},
		{
			name: "President day of 2022",
			year: NewYear(2022, loc),
			weekdayGet: func(y Year) Day {
				return y.GetMonth(time.February).GetWeekDay(time.Monday)[2]
			},
			weekdayCheck: func(d Day) error {
				if d.start.Month() != time.February || d.start.Weekday() != time.Monday {
					return errors.New("check failed")
				}
				return nil
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.weekdayGet(tt.year)
			err := tt.weekdayCheck(got)
			assert.NoError(t, err, "GetWeekDay() expected no error, got error: %+v", err)
		})
	}

}

func TestGetDay(t *testing.T) {
	testCases := []struct {
		name         string
		year         int
		loc          string
		weekdayGet   func(y Year) Day
		weekdayCheck func(d Day) error
	}{
		{
			name: "2022 summer saving day",
			year: 2022,
			loc:  "America/New_York",
			weekdayGet: func(y Year) Day {
				return y.GetMonth(time.March).GetDay(13)
			},
			weekdayCheck: func(d Day) error {
				if d.start.UTC() != time.Date(2022, 3, 13, 5, 0, 0, 0, time.UTC) ||
					d.end.UTC() != time.Date(2022, 3, 14, 3, 59, 59, 999999999, time.UTC) {
					return errors.New("check failed")
				}
				return nil
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			loc, err := time.LoadLocation(tt.loc)
			assert.NoError(t, err, "LoadLocation() expected no error, got error: %+v", err)

			y := NewYear(tt.year, loc)

			got := tt.weekdayGet(y)
			err = tt.weekdayCheck(got)
			assert.NoError(t, err, "GetDay() expected no error, got error: %+v", err)
		})
	}

}

func TestCAConflictAfterShift(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Error(err)
	}

	dayslice := []Day{
		Day{
			DateRange: DateRange{
				start: time.Date(2021, 12, 27, 0, 0, 0, 0, loc),
				end:   time.Date(2021, 12, 27, 23, 59, 59, 999999999, loc),
			},
			OffsetFromOriginal: 1,
		},
		Day{
			DateRange: DateRange{
				start: time.Date(2021, 12, 27, 0, 0, 0, 0, loc),
				end:   time.Date(2021, 12, 27, 23, 59, 59, 999999999, loc),
			},
			OffsetFromOriginal: 2,
		},
		Day{
			DateRange: DateRange{
				start: time.Date(2021, 12, 27, 0, 0, 0, 0, loc),
				end:   time.Date(2021, 12, 27, 23, 59, 59, 999999999, loc),
			},
			OffsetFromOriginal: 0,
		},
	}

	t.Logf("%+v\n", dayslice) // offset should from 2 -> 1 -> 0
	dayslice = CAConflictAfterShift(dayslice)
	t.Logf("%+v\n", dayslice)

}
