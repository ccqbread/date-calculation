package date

import (
	"testing"
	"time"
)

func TestGetWeekDay(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Error(err)
	}

	y0 := NewYear(2022, loc)
	t.Log(y0.GetMonth(time.January).GetWeekDay(time.Monday)[2])

	t.Log(y0.GetMonth(time.February).GetWeekDay(time.Monday)[2])

}

func TestGetDay(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Error(err)
	}

	y0 := NewYear(2022, loc)
	t.Log(y0.GetMonth(time.March).GetDay(13))

	t.Log(y0.GetMonth(time.July).GetDay(4))
}
