package pkg

import (
	"time"
)

func NewCalendar(repo IRepo) *Calendar {
	return &Calendar{repo}
}

type Calendar struct {
	repo IRepo
}

func (c *Calendar) UpdateEvent(event Event) (Event, error) {
	return c.repo.UpdateEvent(event)
}

func (c *Calendar) DeleteEvent(id string) error {
	return c.repo.DeleteEventByID(id)
}

func (c *Calendar) AddEvent(e *Event) error {
	return c.repo.AddEvent(e)
}

func (c *Calendar) GetEventsForDateRange(from time.Time, to time.Time) ([]Event, error) {
	return c.repo.GetEventsForDateRange(from, to)
}
