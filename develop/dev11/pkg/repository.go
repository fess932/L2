package pkg

import (
	"github.com/google/uuid"
	"sync"
	"time"
)

var _ IRepo = &CalRepo{}

type IRepo interface {
	AddEvent(e *Event) error
	UpdateEvent(e Event) (Event, error)
	DeleteEventByID(id string) error

	GetEventsForDateRange(from time.Time, to time.Time) ([]Event, error)
}

func NewCalRepo() *CalRepo {
	return &CalRepo{}
}

type CalRepo struct {
	sync.RWMutex
	Events []Event
}

func (c *CalRepo) AddEvent(e *Event) error {
	c.Lock()
	defer c.Unlock()

	e.ID = uuid.NewString()
	c.Events = append(c.Events, *e)

	return nil
}

func (c *CalRepo) UpdateEvent(e Event) (Event, error) {
	c.Lock()
	defer c.Unlock()

	for i, v := range c.Events {
		if v.ID == e.ID {
			c.Events[i] = e

			return e, nil
		}
	}

	return Event{}, ErrNotFound
}

func (c *CalRepo) DeleteEventByID(id string) error {
	c.Lock()
	defer c.Unlock()

	for i, v := range c.Events {
		if v.ID == id {
			c.Events = append(c.Events[:i], c.Events[i+1:]...)
			return nil
		}
	}

	return ErrNotFound
}

func (c *CalRepo) GetEventsForDateRange(from time.Time, to time.Time) ([]Event, error) {
	c.RLock()
	defer c.RUnlock()

	var e []Event

	// [...)
	for _, v := range c.Events {
		if (v.Date.After(from) || v.Date == from) && v.Date.Before(to) {
			e = append(e, v)
		}
	}

	return e, nil
}
