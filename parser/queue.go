package parser

import (
	"errors"
	"fmt"
)

type Queue struct {
	queue []*Token
}

func (q *Queue) Enqueue(t *Token) {
	q.queue = append(q.queue, t)
}

func (q *Queue) Dequeue() (*Token, error) {
	if len(q.queue) == 0 {
		return nil, errors.New("Empty queue")
	}

	t := q.queue[0]
	q.queue = q.queue[1:]
	return t, nil
}

func (q *Queue) String() string {
	return fmt.Sprintf("%v", q.queue)
}
