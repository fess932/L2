package pkg

import (
	"net/http"
)

func NewCalendarHttpDelivery() *CalendarHttpDelivery {
	return &CalendarHttpDelivery{}
}

type CalendarHttpDelivery struct {
}

func (c *CalendarHttpDelivery) GetEventForDay(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get events_for_day"))
}
