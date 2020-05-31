package counter

import (
	"sync"
)

type Counter struct {
	sync.RWMutex
	counter map[string]uint64
}

func NewCounter() *Counter {
	return &Counter{
		counter: make(map[string]uint64),
	}
}

func (c *Counter) Incr(k string) {
	c.Lock()
	c.counter[k]++
	c.Unlock()
}

func (c *Counter) Val(k string) uint64 {
	var v uint64
	c.RLock()
	v = c.counter[k]
	c.RUnlock()
	return v
}

func (c *Counter) Items() map[string]uint64 {
	c.RLock()
	items := make(map[string]uint64, len(c.counter))
	for k, v := range c.counter {
		items[k] = v
	}
	c.RUnlock()
	return items
}
