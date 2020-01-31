package clock

import (
	"sync"
	"sync/atomic"
	"time"
)

// Clock is a clock accurate to seconds
type Clock struct {
	now int64
}

func NewClock() *Clock {
	o := &Clock{}
	o.updateNow()
	return o
}

// GetNow returns the number of seconds
func (o *Clock) GetNow() int64 {
	return atomic.LoadInt64(&o.now)
}

func (o *Clock) Start(wg *sync.WaitGroup, quit chan bool) {
	wg.Add(1)
	go o.run(wg, quit)
}

func (o *Clock) run(wg *sync.WaitGroup, quit chan bool) {
	defer func() {
		wg.Done()
	}()
	const d = 100 * time.Millisecond
	for {
		select {
		case <-quit:
			return
		default:
		}
		time.Sleep(d)
		o.updateNow()
	}
}

func (o *Clock) updateNow() {
	atomic.StoreInt64(&o.now, time.Now().Unix())
}
