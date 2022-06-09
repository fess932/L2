package pkg

import (
	"net/http"
	"time"
)

type ICalendar interface {
	AddEvent(Event) error
	GetEventsForDate(from time.Time, to time.Time) ([]Event, error)
}

func NewCalendarHTTPDelivery(calendar ICalendar) *CalendarHTTPDelivery {
	return &CalendarHTTPDelivery{calendar}
}

type CalendarHTTPDelivery struct {
	calendar ICalendar
}

func (c *CalendarHTTPDelivery) GetEventForDay(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get events_for_day"))
}

func (c *CalendarHTTPDelivery) CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get events_for_day"))
}
