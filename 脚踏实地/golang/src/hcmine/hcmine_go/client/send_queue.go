package client

import "encoding/json"

type SendQueue struct {
	queue         []interface{}
	highWaterMark uint32
}

func NewQueue(highWaterMark uint32) *SendQueue {
	return &SendQueue{
		queue:         make([]interface{}, 0, highWaterMark),
		highWaterMark: highWaterMark,
	}
}

func (q *SendQueue) Add(item interface{}) bool {
	q.queue = append(q.queue, item)
	return q.IsFull()
}

func (q *SendQueue) IsFull() bool {
	return uint32(len(q.queue)) >= q.highWaterMark
}

func (q *SendQueue) Marshal() ([]byte, error) {
	return json.Marshal(q.queue)
}

func (q *SendQueue) GetCount() uint32 {
	return uint32(len(q.queue))
}
