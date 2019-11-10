package loatr

import (
	"fmt"
	"sync"
	"time"
	//"sync"
)

// SafeCounter ,,,
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc ...
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

// Value ...
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

// TestMutex ...
func TestMutex() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("keyone")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("keyone"))
}
