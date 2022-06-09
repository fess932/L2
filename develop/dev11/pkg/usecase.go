package pkg

import "time"

func NewCalendar() *Calendar {
	return &Calendar{}
}

type Calendar struct {
	Events []Event
}

func (c *Calendar) AddEvent(e Event) error {
	c.Events = append(c.Events, e)
}

func (c *Calendar) GetEventsForDate(from time.Time, to time.Time) ([]Event, error) {

}
