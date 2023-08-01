package concurrency

import (
	"sync/atomic"
	"time"
)

type Queue struct {
	counter uint64
	next    uint64
}

// New creates a wait group that can be incremented/decremented
func New(available uint8) Queue {
	return Queue{next: uint64(available)}
}

// Wait blocks until a slot is ready
func (q *Queue) Wait() {
	var myTurn = atomic.AddUint64(&q.counter, 1)

	for myTurn > q.next {
		time.Sleep(500 * time.Millisecond)
	}
}

func (q *Queue) Done() {
	q.next++
}

func (q *Queue) GetCounter() uint64 {
	return atomic.LoadUint64(&q.counter)
}
