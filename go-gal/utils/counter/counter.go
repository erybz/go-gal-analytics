package counter

import (
	"sync"
)

// Counter is go routine safe counter used to count events
type Counter struct {
	sync.RWMutex
	counter map[string]uint64
}

// NewCounter returns new Counter
func NewCounter() *Counter {
	return &Counter{
		counter: make(map[string]uint64),
	}
}

// Incr increments counter for specified key
func (c *Counter) Incr(k string) {
	c.Lock()
	c.counter[k]++
	c.Unlock()
}

// Val returns current value for specified key
func (c *Counter) Val(k string) uint64 {
	var v uint64
	c.RLock()
	v = c.counter[k]
	c.RUnlock()
	return v
}

// Items returns all the counter items
func (c *Counter) Items() map[string]uint64 {
	c.RLock()
	items := make(map[string]uint64, len(c.counter))
	for k, v := range c.counter {
		items[k] = v
	}
	c.RUnlock()
	return items
}
