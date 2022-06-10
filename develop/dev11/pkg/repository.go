package pkg

import (
	"sync"
	"time"
)

var _ IRepo = &CalRepo{}

type IRepo interface {
	AddEvent(e Event) error

	GetEventByID(id string) (Event, error)
	GetEventsForDateRange(from time.Time, to time.Time) ([]Event, error)
}

func NewCalRepo() *CalRepo {
	return &CalRepo{}
}

type CalRepo struct {
	sync.RWMutex
	Events []Event
}

func (c *CalRepo) AddEvent(e Event) error {
	c.Lock()
	defer c.Unlock()

	c.Events = append(c.Events, e)

	return nil
}

func (c *CalRepo) GetEventByID(id string) (Event, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CalRepo) GetEventsForDateRange(from time.Time, to time.Time) ([]Event, error) {
	c.RLock()
	defer c.RUnlock()

	var e []Event

	for _, v := range c.Events {
		if v.Date.After(from) && v.Date.Before(to) {
			e = append(e, v)
		}
	}

	return e, nil
}
