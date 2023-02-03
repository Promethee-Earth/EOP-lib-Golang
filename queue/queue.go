package queue

import (
	"sync/atomic"
	"time"
)

type Queue struct {
	counter uint64
	next    uint64
}

// New creates a blocking wait group that can be incremented
func New(available uint8) Queue {
	return Queue{next: uint64(available)}
}

func (q *Queue) Wait() {
	myTurn := atomic.AddUint64(&q.counter, 1)

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
