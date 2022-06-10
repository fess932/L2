package pkg

import (
	"fmt"
	"net/http"
	"time"
)

type ICalendar interface {
	AddEvent(*Event) error
	GetEventsForDateRange(from time.Time, to time.Time) ([]Event, error)
}

func NewCalendarHTTPDelivery(calendar ICalendar) *CalendarHTTPDelivery {
	return &CalendarHTTPDelivery{calendar}
}

type CalendarHTTPDelivery struct {
	calendar ICalendar
}

// CreateEvent for date
func (c *CalendarHTTPDelivery) CreateEvent(w http.ResponseWriter, r *http.Request) {
	e := Event{
		Title: r.FormValue("title"),
		Date:  time.Time{},
	}

	date, err := parseDate(r.FormValue("date"))
	if err != nil {
		JSONError(w, http.StatusBadRequest, err)

		return
	}

	e.Date = date

	if err = c.calendar.AddEvent(&e); err != nil {
		JSONError(w, http.StatusInternalServerError, err)

		return
	}

	JSONResponse(w, http.StatusCreated, e)
}

func (c *CalendarHTTPDelivery) GetEventForDay(w http.ResponseWriter, r *http.Request) {
	c.getEvensForRange(w, r, Day)
}
func (c *CalendarHTTPDelivery) GetEventForWeek(w http.ResponseWriter, r *http.Request) {
	c.getEvensForRange(w, r, Week)
}
func (c *CalendarHTTPDelivery) GetEventForMonth(w http.ResponseWriter, r *http.Request) {
	c.getEvensForRange(w, r, Month)
}

func (c *CalendarHTTPDelivery) getEvensForRange(w http.ResponseWriter, r *http.Request, rang int) {
	from, to, err := parseRange(r.FormValue("date"), rang)
	if err != nil {
		JSONError(w, http.StatusBadRequest, err)

		return
	}

	events, err := c.calendar.GetEventsForDateRange(from, to)
	if err != nil {
		JSONError(w, http.StatusServiceUnavailable, err)

		return
	}

	JSONResponse(w, http.StatusOK, events)
}

func parseDate(input string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", input)
	if err != nil {
		return time.Time{}, fmt.Errorf("cant parse date: %w", err)
	}

	return date, nil
}

const (
	Day = iota + 1
	Week
	Month
)

func parseRange(input string, rang int) (from, to time.Time, err error) {
	date, err := parseDate(input)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("cant parse date: %w", err)
	}

	switch rang {
	case Day:
		from = date
		to = date.AddDate(0, 0, 1)
	case Week:
		from = date.AddDate(0, 0, -int(date.Weekday()))
		to = date.AddDate(0, 0, 7-int(date.Weekday()))
	case Month:
		from = date.AddDate(0, 0, -date.Day())
		to = date.AddDate(0, 1, -date.Day())
	default:
		return time.Time{}, time.Time{}, fmt.Errorf("unknown range: %w", err)
	}

	return
}
