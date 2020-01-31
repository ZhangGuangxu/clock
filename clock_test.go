package clock

import (
	"sync"
	"testing"
	"time"
)

func TestClock(t *testing.T) {
	wg := &sync.WaitGroup{}
	quit := make(chan bool)
	o := NewClock()
	o.Start(wg, quit)
	time.Sleep(time.Second)

	t1 := o.GetNow()
	time.Sleep(time.Second)
	t2 := o.GetNow()
	if t1 >= t2 {
		t.Errorf("t1[%v] >= t2[%v]", t1, t2)
	}

	close(quit)
	wg.Wait()
}
