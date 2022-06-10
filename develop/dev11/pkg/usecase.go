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

func (c *Calendar) AddEvent(e Event) error {
	return c.repo.AddEvent(e)
}

func (c *Calendar) GetEventsForDateRange(from time.Time, to time.Time) ([]Event, error) {
	return c.repo.GetEventsForDateRange(from, to)
}
